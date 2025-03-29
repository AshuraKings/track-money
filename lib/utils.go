package lib

import (
	"net/url"
	"strings"
)

func QueryToMap(query url.Values) map[string]string {
	result := map[string]string{}
	for k, v := range query {
		result[k] = strings.Join(v, ", ")
	}
	return result
}
