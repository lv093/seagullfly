package initialize

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
	"time"
	"seagullfly/utils"
)

func RegisterRedis() {
	logs.Info("Lib:Cache Init, Start!")
	client := initRedis("redis")
	utils.SetRedisClient(client)
	logs.Info("Lib:Cache Init, Finish!")
}

func initRedis(name string) *redis.Pool {
	logs.Info("Lib:Cache Redis Init, Start!", name)
	RedisAddress := beego.AppConfig.String(name + ".address")
	RedisDb, _ := beego.AppConfig.Int(name + ".db")
	// TODO 建立连接池, 连接池报错
	client := &redis.Pool{
		MaxIdle:     beego.AppConfig.DefaultInt(name+".max_idle", 100),
		MaxActive:   beego.AppConfig.DefaultInt(name+".max_active", 2000),
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", RedisAddress,
				redis.DialConnectTimeout(time.Duration(time.Millisecond*500)),
				redis.DialReadTimeout(time.Duration(time.Millisecond*500)),
				redis.DialWriteTimeout(time.Duration(time.Millisecond*500)),
			)
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", RedisDb)
			return c, nil
		},
	}
	logs.Info("Lib:Cache Redis Init, Finish!", name)

	return client
}
