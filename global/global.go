package global

import (
	"SCIProj/config"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	VP     *viper.Viper
	CONFIG *config.Server
	LOG    *zap.Logger
	REDIS  *redis.Client
	JWTKey []byte
)
