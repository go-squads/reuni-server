package appcontext
import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/go-squads/reuni-server/config"
	"log"
)
type appContext struct {
	db          *sql.DB
}

var context *appContext

func initDB() (*sql.DB,error) {

	db, err := sql.Open("postgres", config.GetConnectionString())

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}
	return db,nil
}

func InitContext() {
	db,_ := initDB()
	log.Print("Connection Established")
	context = &appContext{
		db : db,
	}
}

func GetDB() *sql.DB{
	return context.db
}

