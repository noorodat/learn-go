package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Authorization: ApiKey {insert apikey here}
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New(("no auth info found"))
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New(("malformed auth header"))
	}
	if vals[0] != "ApiKey" {
		return "", errors.New(("malformed first part of auth header"))
	}
	return vals[1], nil
}
