package common

import "database/sql"

type querier interface {
	QueryRow(query string, args ...any) *sql.Row
}

// parameter input :
// sql dari method (r.db), query count mysql, data lainnya...
func CountRows(q querier, query string, args ...any) (int, error) {
	var total int
	if err := q.QueryRow(query, args...).Scan(&total); err != nil {
		return 0, err
	}

	return total, nil
}
