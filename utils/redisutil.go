package utils

import (
		"github.com/garyburd/redigo/redis"
	"recitewords/callbacks"
)

type RedisConfig struct {
	network string
	address string
	// options redis.DialOption // TODO 看看这玩意儿有没有更好的作用
}

type RedisUtil struct {
	config RedisConfig
}


func GetRedisConfig(network, address string) *RedisConfig{
	return &RedisConfig{
		network:network,
		address:address,
	}
}

func NewRedisUtil(config RedisConfig) *RedisUtil {
	return &RedisUtil{config}
}

/**
 * 所有的 Run 方法，都需要自己去维护 conn.Close
 */


func (this *RedisUtil) Run(command string, doCommand callbacks.DoCommand, params...interface{}) (result interface{}, err error){
	conn, err := redis.Dial(this.config.network, this.config.address)
	if err != nil {
		return
	}
	defer conn.Close()
	return doCommand(conn, command, params...)
}
