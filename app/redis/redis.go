package redis

import (
	"app/models"
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
)

/*
RedisClient redis.Conn
*/
type RedisClient struct {
	Client redis.Conn
}

/*
Connect func(Proto string, Addr string, Port int) (redis.Conn, error)
*/
func Connect(rc models.RedisConnector) (redis.Conn, error) {
	conn, err := redis.Dial(rc.Proto, fmt.Sprintf("%v:%v", rc.Addr, rc.Port))
	return conn, err
}

/*
Set func(key string, value interface{})
*/
func (rc *RedisClient) Set(key string, value interface{}) error {
	reply, err := rc.Client.Do("SET", key, value)
	log.Printf("Redis SET Reply : %v", reply)
	return err
}

/*
Get func(key string)
*/
func (rc *RedisClient) Get(key string) (interface{}, error) {
	reply, err := redis.String(rc.Client.Do("Get", key))
	log.Printf("Redis GET Reply : %v", reply)
	return reply, err
}

/*
Lpush func(key string, value interface{})
*/
func (rc *RedisClient) Lpush(key string, value interface{}) error {
	reply, err := rc.Client.Do("LPUSH", key, value)
	log.Printf("Redis LPUSH Reply : %v", reply)
	return err
}

/*
Rpop func(key string)
*/
func (rc *RedisClient) Rpop(key string) (interface{}, error) {
	reply, err := redis.String(rc.Client.Do("RPOP", key))
	log.Printf("Redis RPOP Reply : %v", reply)
	return reply, err
}

/*
Close func()
*/
func (rc *RedisClient) Close() error {
	err := rc.Client.Close()
	return err
}
