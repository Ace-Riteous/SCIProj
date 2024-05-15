package utils

import (
	"fmt"
	"strings"
)

func GeneralRedisKey(originKey string, replacePart string, replaceValue string, prefix interface{}) string {

	//类型断言
	if _, ok := prefix.(string); !ok {
		prefix = ""

	}
	return fmt.Sprintf("%s%s", prefix, strings.Replace(originKey, replacePart, replaceValue, -1))

}
