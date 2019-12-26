package callbacks

import (
	"github.com/garyburd/redigo/redis"
	)

type DoCommand func(conn redis.Conn, commandname string, params ...interface{}) (result interface{}, err error)

func DoCommandImpl (conn redis.Conn, commandname string, params ...interface{}) (result interface{}, err error) {
	return conn.Do(commandname, params...)
}