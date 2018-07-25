package appcontext

import (
	"database/sql"
	"log"

	"github.com/go-squads/reuni-server/config"
	_ "github.com/lib/pq"
)

type appContext struct {
	db  *sql.DB
	key *config.Keys
}

var context *appContext

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func initDB() (*sql.DB, error) {

	db, err := sql.Open("postgres", config.GetConnectionString())
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

func GetDB() *sql.DB {
	return context.db
}
