package main

import (
	"encoding/json"
)

func parser(data any) (string, error) {

	str, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		return "", err
	}

	return string(str), nil
}
