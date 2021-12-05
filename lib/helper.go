package lib

import (
	"io"
	"net/url"
	"strings"
)

func CreateFormReader(data map[string]string) io.Reader {
	form := url.Values{}
	for k, v := range data {
		form.Add(k, v)
	}
	return strings.NewReader(form.Encode())
}

func ArrayContains(array []string, item string) bool {
	for _, a := range array {
		if a == item {
			return true
		}
	}
	return false
}
