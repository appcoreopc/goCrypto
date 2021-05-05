package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

const Jwt_Delimiter = "."

type JWT struct {
	Token      string
	SecretKey  string
	JwtHeader  string
	JwtPayLoad string
}

func NewJWT(tokenString string) *JWT {
	return &JWT{Token: tokenString}
}

// Json parse // json url decode // string split //
func (j *JWT) Decode(token string) {

	stringContentArray := strings.Split(token, Jwt_Delimiter)
	//jwtHeader := make([]string, 2)

	j.JwtHeader = j.decode(stringContentArray[0])
	j.JwtPayLoad = j.decode(stringContentArray[1])

	fmt.Println(j.JwtHeader)
	fmt.Println(j.JwtPayLoad)

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
func (j *JWT) VerifyToken() string {

	jwtHeader := base64.StdEncoding.EncodeToString([]byte(j.JwtHeader))
	fmt.Println(jwtHeader)

	jwtPayLoad := base64.StdEncoding.EncodeToString([]byte(j.JwtPayLoad))
	fmt.Println(jwtPayLoad)

	jwtContentAsString := jwtHeader + "." +
		jwtPayLoad + ","

	validToken := j.ComputeHmac256(jwtContentAsString, j.SecretKey)
	return validToken
}

func (j *JWT) ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
