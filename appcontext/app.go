package appcontext

import (
	"log"
	"os"

	"github.com/go-redis/redis"

	"github.com/go-squads/reuni-server/helper"

	"github.com/go-squads/reuni-server/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type appContext struct {
	db          *sqlx.DB
	helper      helper.QueryExecuter
	key         *config.Keys
	redis       *redis.Client
	redisHelper helper.RedisExecuter
}

var context *appContext

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func initDB() (*sqlx.DB, error) {

	db, err := sqlx.Open("postgres", config.GetConnectionString())
	check(err)
	err = db.Ping()
	check(err)
	return db, nil
}

func initRedis() (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := redisClient.Ping().Result()
	check(err)
	return redisClient, nil
}

func initKey() (*config.Keys, error) {
	keys, err := config.GetKeys()
	check(err)
	return keys, nil
}

func InitContext() {
	db, _ := initDB()
	log.Print("Database Connection Established")
	key, _ := initKey()
	log.Print("RSA Keys fetched")
	redis, _ := initRedis()
	log.Print("Redis Connection Established")
	context = &appContext{
		db:    db,
		key:   key,
		redis: redis,
		helper: &helper.QueryHelper{
			DB: db,
		},
		redisHelper: &helper.RedisHelper{
			Redis: redis,
		},
	}
}

func InitMockContext(q helper.QueryExecuter) {
	priv, pub := helper.GenerateRsaKeyPair()
	key := config.Keys{
		PrivateKey: priv,
		PublicKey:  pub,
	}
	context = &appContext{
		key:    &key,
		helper: q,
	}

}

func GetDB() *sqlx.DB {
	return context.db
}
func GetHelper() helper.QueryExecuter {
	return context.helper
}

func GetKeys() *config.Keys {
	return context.key
}

func GetRedisHelper() helper.RedisExecuter {
	return context.redisHelper
}
