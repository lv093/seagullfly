package utils

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
	"strconv"
)

var (
	redisClient *redis.Pool
)

func SetRedisClient(client *redis.Pool) {
	redisClient = client
}

func GetRedisClient() *redis.Pool {
	return redisClient
}

func GetRedisConn() redis.Conn {
	if redisClient == nil {
		logs.Error("GetRedisConn redisClient nil")
		return nil
	}

	conn := redisClient.Get()
	if conn.Err() != nil {
		logs.Error("GetRedisConn get nil", conn.Err())
		conn = redisClient.Get()
		if conn.Err() != nil {
			logs.Error("second GetRedisConn get nil", conn.Err())
		}
	}
	return conn
}

func LockByRedis(key string, ttl int64) bool {
	client := redisClient.Get()
	defer client.Close()

	num, _ := client.Do("incr", key)
	res := fmt.Sprint(num) == "1"
	if res {
		client.Do("expire", key, ttl)
	}
	return res
}

func ZCardByRedis(key string) int {
	client := redisClient.Get()
	defer client.Close()
	num, _ := client.Do("ZCARD", key)
	n, _ := strconv.Atoi(fmt.Sprint(num))
	return n
}

func ZCountByRedis(key string, start, end int64) int {
	client := redisClient.Get()
	defer client.Close()
	num, _ := client.Do("ZCOUNT", key, start, end)
	n, _ := strconv.Atoi(fmt.Sprint(num))
	return n
}
