package helper

import (
	"strconv"
	"strings"
)

func ReplaceQueryParams(query string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    = 1
		args []interface{}
	)

	for k, v := range params {
		if k != "" && strings.Contains(query, ":"+k) {
			query = strings.ReplaceAll(query, ":"+k, "$"+strconv.Itoa(i))
			args = append(args, v)
			i++
		}
	}

	return query, args
}
