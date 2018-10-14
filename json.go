package main

import (
	"encoding/json"
	"io"
)

func parseJSON(r io.Reader) (map[string]string, error) {
	values := map[string]string{}
	err := json.NewDecoder(r).Decode(&values)
	return values, err
}
