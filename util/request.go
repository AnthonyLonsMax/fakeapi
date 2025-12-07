package util

import (
	"net/http"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

func GetIntFromQuery(r *http.Request, key string, defaultValue int) (int, error) {
	query := r.URL.Query().Get(key)
	if query == "" {
		return defaultValue, nil
	}
	parsed, err := strconv.Atoi(query)
	if err != nil {
		return -1, err
	}
	return parsed, nil
}

func GetStringFromQuery(r *http.Request, key, defaultValue string) string {
	query := r.URL.Query().Get(key)
	if query == "" {
		return defaultValue
	}
	return query
}

func SortMap(values []map[string]any, key string) {
	slices.SortFunc(values, func(a, b map[string]any) int {
		// contains the key
		v1, ok1 := a[key]
		v2, ok2 := b[key]
		if !(ok1 && ok2) {
			// if the field is not present sort by key count
			return len(a) - len(b)
		}
		type1 := reflect.TypeOf(v1)
		type2 := reflect.TypeOf(v2)

		if type1.Kind() != type2.Kind() {
			// invalid types
			return len(a) - len(b)
		}

		switch type1.Kind() {
		case reflect.Float64:
			return int(v1.(float64)) - int(v2.(float64))
		case reflect.Int:
			return int(v1.(int)) - int(v2.(int))
		case reflect.String:
			return strings.Compare(v1.(string), v2.(string))
		case reflect.Bool:
			bool1, _ := v1.(bool)
			bool2, _ := v2.(bool)
			vbool1 := 0
			vbool2 := 0
			if bool1 {
				vbool1 = 1
			}
			if bool2 {
				vbool2 = 1
			}
			return vbool2 - vbool1
		default:
			return len(a) - len(b)
		}
	})
}
