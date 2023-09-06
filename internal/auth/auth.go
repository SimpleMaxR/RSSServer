package auth

import (
	"errors"
	"net/http"
)

var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")

func GetAPIKey(header http.Header) (string, error) {
	authHeader := header.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	//splitAuth := strings.Split(authHeader, " ")
	//if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
	//	return "", errors.New("malformed auth header")
	//}

	return authHeader, nil
}
