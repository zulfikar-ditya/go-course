package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func GetAPIKey(header *http.Header) (string, error) {
	authorization := header.Get("Authorization")

	if authorization == "" {
		fmt.Println("authorization header is empty");
		return "", errors.New("missing authorization header")
	}

	vals := strings.Split(authorization, " ")

	if len(vals) != 2 {
		fmt.Println("authorization header is invalid")
		return "", errors.New("invalid authorization header")
	}

	if vals[0] != "Bearer" {
		fmt.Println("authorization header is invalid bearer")
		return "", errors.New("invalid authorization header")
	}

	return vals[1], nil
}