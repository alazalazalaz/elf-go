package redis

import (
	"elf-go/components/config"
	"fmt"
	"github.com/go-redis/redis"
)

type Redis struct {
	Conf config.ConfRedis
	*redis.Client
}

func New(c *config.Config) *Redis {
	return &Redis{
		Conf: c.GetRedisConfig(),
	}
}

func (r *Redis) Init() error {
	r.Client = redis.NewClient(r.setOptions())

	if _, err := r.Ping().Result(); err != nil {
		return err
	}

	return nil
}

func (r *Redis) setOptions() *redis.Options {
	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Conf.Ip, r.Conf.Port),
		Password: "", // no password set
		DB:       0,  // use default DB
	}
}
