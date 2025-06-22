package tools

import (
	"database/sql"
	"strings"
)

/*
To work quickly sometimes you just want a map of the data rather than creating a viewmodel
*/
func QueryToMaps(db *sql.DB, query string, args ...any) ([]map[string]any, error) {

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var results []map[string]any

	for rows.Next() {
		values := make([]any, len(cols))
		pointers := make([]any, len(cols))

		for i := range pointers {
			pointers[i] = &values[i]
		}

		if err := rows.Scan(pointers...); err != nil {
			return nil, err
		}

		rowMap := make(map[string]any)
		for i, col := range cols {
			key := underscoredToTitleCase(col)
			val := values[i]

			switch v := val.(type) {
			case []byte:
				rowMap[key] = string(v)
			case nil:
				rowMap[key] = ""
			default:
				rowMap[key] = v
			}

		}

		results = append(results, rowMap)

	}

	return results, rows.Err()

}

func underscoredToTitleCase(s string) string {
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_'
	})

	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = strings.ToUpper(part[:1]) + strings.ToLower(part[1:])
		}
	}

	return strings.Join(parts, "")
}
