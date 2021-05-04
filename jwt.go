package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

const Jwt_Delimiter = "."

type JWT struct {
}

// Json parse // json url decode // string split //
func (j *JWT) Decode(token string) []string {

	stringContentArray := strings.Split(token, Jwt_Delimiter)
	jwtHeader := make([]string, 2)

	for _, str := range stringContentArray[0:2] {
		fmt.Println(str)
		//data, err := base64.StdEncoding.DecodeString("eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ")
		data, err := base64.StdEncoding.DecodeString(str)
		if err != nil {
			fmt.Println("error decoding the given base64")
			//return []strings{}
		}
		jwtHeader = append(jwtHeader, string(data))
		fmt.Printf("%s - decoded\n", data)
		//fmt.Print(str)
	}
	return jwtHeader
}
