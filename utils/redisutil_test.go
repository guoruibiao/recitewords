package utils

import (
	"testing"
		"recitewords/callbacks"
	)

var config RedisConfig

func init() {
	config = RedisConfig{
		"tcp",
		"ip:port",
	}
}

func TestRedisUtil_Run(t *testing.T) {
	instance := NewRedisUtil(config)
	params := []interface{}{"name", "abcdefg"}
	result, err := instance.Run("set", callbacks.DoCommandImpl, params...)
	if err != nil {
		t.Error(err)
	}else{
		t.Log(result)
	}
}
