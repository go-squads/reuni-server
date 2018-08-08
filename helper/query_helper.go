package helper

import (
	"github.com/jmoiron/sqlx"
)

type QueryExecuter interface {
	DoQuery(query string, args ...interface{}) ([]map[string]interface{}, error)
	DoQueryRow(query string, args ...interface{}) (map[string]interface{}, error)
}

type QueryHelper struct {
	DB *sqlx.DB
}

func parseRows(rows *sqlx.Rows) ([]map[string]interface{}, error) {
	var data []map[string]interface{}
	for rows.Next() {
		datum := make(map[string]interface{})
		err := rows.MapScan(datum)
		if err != nil {
			return nil, err
		}
		data = append(data, datum)
	}
	return data, nil
}

func (q *QueryHelper) DoQuery(query string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := q.DB.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	data, err := parseRows(rows)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (q *QueryHelper) DoQueryRow(query string, args ...interface{}) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	err := q.DB.QueryRowx(query, args...).MapScan(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
