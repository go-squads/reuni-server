package helper

type QueryMockHelper struct {
	Data []map[string]interface{}
	Err  error
}

func (q *QueryMockHelper) DoQuery(query string, args ...interface{}) ([]map[string]interface{}, error) {
	if q.Err != nil {
		return nil, q.Err
	}
	return q.Data, nil
}
