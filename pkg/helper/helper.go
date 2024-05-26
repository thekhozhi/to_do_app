package helper

import (
	"strconv"
	"strings"
)

func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var(
		i = 1
		args []interface{}
	)

	for key, value := range params {
		if key != ""{
			oldSize := len(namedQuery)
			namedQuery = strings.ReplaceAll(namedQuery, "@"+key, "$"+strconv.Itoa(i))

			if oldSize != len(namedQuery){
				args = append(args, value)
				i++
			}
		}
	}

	return namedQuery, args
}