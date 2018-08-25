package helper

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

type RedisExecuter interface {
	Publish(channel string, message interface{}) error
}

type RedisHelper struct {
	Redis *redis.Client
}

func (r *RedisHelper) Publish(channel string, message interface{}) error {
	err := r.Redis.Publish(channel, message).Err()
	if err != nil {
		log.Println("CreateNewConfigVersion: ", err.Error())
		return err
	}
	return nil
}
