package spider

import (
	"github.com/go-redis/redis"
	"sync"
	"strconv"
	"github.com/lwl1989/go-spider/config"
)

type redisConnection struct {
	con *redis.Client
	conf *config.Connections
}
var redisCon *redisConnection
var redisSynOnce sync.Once

func GetRedis(conf *config.Connections)  *redisConnection{
	redisSynOnce.Do(func() {
		redisCon = &redisConnection{
			conf:conf,
		}
		db,err := strconv.Atoi(redisCon.conf.Redis.Db)
		if err != nil {
			db = 0
		}
		redisCon.con = redis.NewClient(&redis.Options{
			Addr:    redisCon.conf.Redis.Host,
			Password: redisCon.conf.Redis.Pw, // no password set
			DB:      db,  // use default DB
		})
	})

	_, err := redisCon.con.Ping().Result()
	if err != nil {
		redisCon.reConnection()
	}
	return redisCon
}

func (conn *redisConnection) reConnection() {
	db,err := strconv.Atoi(conn.conf.Redis.Db)
	if err != nil {
		db = 0
	}
	conn.con = nil
	conn.con = redis.NewClient(&redis.Options{
		Addr:    conn.conf.Redis.Host,
		Password: conn.conf.Redis.Pw, // no password set
		DB:      db,  // use default DB
	})
}

func (conn *redisConnection) dConnection() {
	_, err := redisCon.con.Ping().Result()
	if err != nil {
		redisCon.reConnection()
	}
}
//Keys(pattern string) []string
//SaveKey(key string, value interface{})
//GetKey(key string) interface{}
//HSet(key string, field string, value interface{})
//HGet(key string, field string) string
//HMGet(key string) map[string]string
func (conn *redisConnection) SaveKey(key string, value interface{}) {
	conn.con.LPush(key, value)
}

func (conn *redisConnection) Keys(pattern string) []string {
	conn.dConnection()

	return conn.con.Keys(pattern).Val()
}

func (conn *redisConnection) GetKey(key string) []string {
	conn.dConnection()
	return conn.con.LRange(key, 0, -1).Val()
	//return conn.con.Get(key).Val()
}

func (conn *redisConnection) HSet(key string, field string, value interface{}) {
	conn.dConnection()

	conn.con.HSet(key, field, value)
}

func (conn *redisConnection) HMSet(key string,fields map[string]interface{}) {
	conn.dConnection()

	conn.con.HMSet(key, fields)
}

func (conn *redisConnection) HGet(key string, field string) string {
	conn.dConnection()

	return conn.con.HGet(key, field).Val()
}

func (conn *redisConnection) HMGet(key string, fields string) []interface{} {
	conn.dConnection()

	return conn.con.HMGet(key, fields).Val()
}

func (conn *redisConnection) HGetAll(key string) map[string]string {
	conn.dConnection()

	return conn.con.HGetAll(key).Val()
}

func (conn *redisConnection) LRange(key string) []string{
	conn.dConnection()

	return conn.con.LRange(key,0 , -1).Val()
}

func (conn *redisConnection) LPush(key string, values ...string) {
	conn.dConnection()
	conn.con.LPush(key, values)
}
