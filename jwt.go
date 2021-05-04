package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

type JWT struct {
}

// Json parse // json url decode // string split //
func (j *JWT) Decode(token string) []string {

	stringContentArray := strings.Split(token, ".")

	for str := range stringContentArray {

		data, err := base64.StdEncoding.DecodeString("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")
		if err != nil {
			return []string{}
		}
		fmt.Printf("%s\n", data)
		fmt.Print(str)
	}

	return []string{}

}
