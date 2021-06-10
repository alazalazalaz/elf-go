package redis

import (
	"elf-go/components/config"
	"elf-go/components/logs"
	"fmt"
	"github.com/go-redis/redis"
)

func NewRedis(conf config.ConfRedis) (*redis.Client, error){
	ip := conf.Ip
	port := conf.Port
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", ip, port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if _, err := client.Ping().Result(); err != nil{
		return nil, err
	}

	//err = client.Set("key", "value", time.Minute * 10).Err()
	//if err != nil {
	//	panic(err)
	//}
	logs.Info("redis init ok")
	return client, nil
}