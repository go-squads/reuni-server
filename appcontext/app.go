package appcontext

import (
	"log"

	"github.com/go-squads/reuni-server/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type appContext struct {
	db  *sqlx.DB
	key *config.Keys
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
	context = &appContext{
		db:  db,
		key: key,
	}
}

func GetDB() *sqlx.DB {
	return context.db
}

func GetKeys() *config.Keys {
	return context.key
}
