package redisctr

import (
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
)

/*
RedisController redis.Conn
*/
type RedisController struct {
	Proto string
	Addr  string
	Port  int
	Conn  redis.Conn
}

/*
Connect func(Proto string, Addr string, Port int) (redis.Conn, error)
*/
func (rc *RedisController) Connect() error {
	conn, err := redis.Dial(rc.Proto, fmt.Sprintf("%v:%v", rc.Addr, rc.Port))
	rc.Conn = conn
	return err
}

/*
Set func(key string, value interface{})
*/
func (rc *RedisController) Set(key string, value interface{}) error {
	reply, err := rc.Conn.Do("SET", key, value)
	log.Printf("Redis SET Reply : %v", reply)
	return err
}

/*
Get func(key string)
*/
func (rc *RedisController) Get(key string) (interface{}, error) {
	reply, err := redis.String(rc.Conn.Do("Get", key))
	log.Printf("Redis GET Reply : %v", reply)
	return reply, err
}

/*
Lpush func(key string, value interface{})
*/
func (rc *RedisController) Lpush(key string, value interface{}) error {
	reply, err := rc.Conn.Do("LPUSH", key, value)
	log.Printf("Redis LPUSH Reply : %v", reply)
	return err
}

/*
Rpop func(key string)
*/
func (rc *RedisController) Rpop(key string) (interface{}, error) {
	reply, err := redis.String(rc.Conn.Do("RPOP", key))
	log.Printf("Redis RPOP Reply : %v", reply)
	return reply, err
}

/*
Close func()
*/
func (rc *RedisController) Close() error {
	err := rc.Conn.Close()
	return err
}
