package helper

import (
	"github.com/jmoiron/sqlx"
)

type QueryExecuter interface {
	doQueryRow(query string, args ...interface{}) (interface{}, error)
	doQuery(query string, structs interface{}, args ...interface{}) (struct{}, error)
}

type QueryHelper struct {
	db *sqlx.DB
}

func parseRows(rows *sqlx.Rows, type_struct interface{}) {

}

func (q *QueryHelper) doQuery(query string, args ...interface{}) (interface{}, error) {
	q.db.Query(query)
	return nil, nil
}
