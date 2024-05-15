package initialize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

var rdClient *redis.Client
var nDuration = 30 * 24 * 60 * 60 * time.Second

type RedisClient struct {
}

func Redis() (*RedisClient, error) {

	rdClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("db.redis"),
		Password: "",
		DB:       0,
	})

	_, err := rdClient.Ping(context.Background()).Result()

	if err != nil {
		return nil, err
	}

	return &RedisClient{}, nil
}

func (m *RedisClient) Set(key string, value any, rest ...any) error {
	d := nDuration
	if len(rest) > 0 {
		if v, ok := rest[0].(time.Duration); ok {
			d = v
		}
	}
	return rdClient.Set(context.Background(), key, value, d).Err()
}

func (m *RedisClient) SetUintList(key string, list []uint, rest ...any) error {
	d := nDuration
	if len(rest) > 0 {
		if v, ok := rest[0].(time.Duration); ok {
			d = v
		}
	}
	for _, item := range list {
		err := rdClient.RPush(context.Background(), key, item).Err()
		if err != nil {
			return err
		}
	}
	rdClient.Expire(context.Background(), key, d)
	return nil
}

func (m *RedisClient) DeleteItemFromUintList(key string, item uint, rest ...any) error {
	d := nDuration
	if len(rest) > 0 {
		if v, ok := rest[0].(time.Duration); ok {
			d = v
		}
	}

	rdClient.Expire(context.Background(), key, d)
	return rdClient.LRem(context.Background(), key, 0, item).Err()
}

func (m *RedisClient) AddItemToUintList(key string, item uint, rest ...any) error {
	d := nDuration
	if len(rest) > 0 {
		if v, ok := rest[0].(time.Duration); ok {
			d = v
		}
	}
	rdClient.Expire(context.Background(), key, d)
	return rdClient.RPush(context.Background(), key, item).Err()
}

func (m *RedisClient) GetUintList(key string) ([]uint, error) {
	vals, err := rdClient.LRange(context.Background(), key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	var list []uint
	for _, val := range vals {
		item, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			return nil, err
		}
		list = append(list, uint(item))
	}
	return list, nil
}

func (m *RedisClient) GetKeysAndUintListValues(pattern string) (map[string][]uint, error) {
	keys, err := rdClient.Keys(context.Background(), pattern).Result()
	if err != nil {
		return nil, err
	}

	result := make(map[string][]uint)
	for _, key := range keys {
		list, err := m.GetUintList(key)
		if err != nil {
			return nil, err
		}
		result[key] = list
	}
	return result, nil
}

func (m *RedisClient) Get(key string) (any, error) {
	return rdClient.Get(context.Background(), key).Result()
}

func (m *RedisClient) Delete(key ...string) error {
	//支持批量删除
	return rdClient.Del(context.Background(), key...).Err()
}

func (m *RedisClient) GetDuration(key string) (time.Duration, error) {
	return rdClient.TTL(context.Background(), key).Result()
}

func (m *RedisClient) GetInt64(key string) (int64, error) {

	return rdClient.Get(context.Background(), key).Int64()
}

func (m *RedisClient) GetUint(key string) (uint, error) {
	value, err := rdClient.Get(context.Background(), key).Result()
	if err != nil {
		return 0, err
	}

	u, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(u), nil
}

func (m *RedisClient) GetList(key string) ([]string, error) {
	return rdClient.LRange(context.Background(), key, 0, -1).Result()
}

func (m *RedisClient) GetBool(key string) (bool, error) {
	val, err := rdClient.Get(context.Background(), key).Result()
	if err != nil {
		//fmt.Printf("getbool-1-err:%s\n", err)
		return false, err
	}

	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		//fmt.Printf("getbool-2-err:%s\n", err)
		return false, err
	}

	return boolVal, nil
}

func (m *RedisClient) GetKeysAndValues(pattern string) (map[string]string, error) {
	keys, err := rdClient.Keys(context.Background(), pattern).Result()
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for _, key := range keys {
		value, err := rdClient.Get(context.Background(), key).Result()
		if err != nil {
			return nil, err
		}
		result[key] = value
	}

	return result, nil
}
