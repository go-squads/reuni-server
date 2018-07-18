package appcontext
import (
	"database/sql"
	_ "github.com/lib/pq"
)
type appContext struct {
	db          *sql.DB
}

func InitDB() {
	db, err := sql.Open("postgres", "")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}
}


