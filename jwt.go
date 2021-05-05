package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

const Jwt_Delimiter = "."

type JWT struct {
	secretKey  string
	jwtHeader  string
	jwtPayLoad string
}

// Json parse // json url decode // string split //
func (j *JWT) Decode(token string) {

	stringContentArray := strings.Split(token, Jwt_Delimiter)
	//jwtHeader := make([]string, 2)

	j.jwtHeader = j.decode(stringContentArray[0])
	j.jwtPayLoad = j.decode(stringContentArray[1])

	fmt.Println(j.jwtHeader)
	fmt.Println(j.jwtPayLoad)

	// data, err := base64.StdEncoding.DecodeString(stringContentArray[0])
	// if (err != nil) {
	// 	j.jwtHeader = string(data)
	// }

	// for _, str := range stringContentArray[0:2] {
	// 	fmt.Println(str)
	// 	data, err := base64.StdEncoding.DecodeString(str)
	// 	if err != nil {
	// 		fmt.Println("error decoding the given base64")
	// 		//return []strings{}
	// 	}

	// 	jwtHeader = append(jwtHeader, string(data))
	// 	fmt.Printf("%s - decoded\n", data)
	// 	//fmt.Print(str)
	// }
	// return jwtHeader
}

func (j *JWT) decode(token string) string {
	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		fmt.Println("error")
	}
	return string(data)
}

// HMACSHA256(
// 	base64UrlEncode(header) + "." +
// 	base64UrlEncode(payload),
// 	1000
// )
func (j *JWT) VerifyToken() [32]byte {

	jwtHeader := base64.StdEncoding.EncodeToString([]byte(j.jwtHeader))
	fmt.Println(jwtHeader)

	jwtPayLoad := base64.StdEncoding.EncodeToString([]byte(j.jwtPayLoad))
	fmt.Println(jwtPayLoad)

	jwtContentAsString := jwtHeader + "." +
		jwtPayLoad + "," +
		j.secretKey

	validToken := sha256.Sum256([]byte(jwtContentAsString))
	return validToken
}
