package main

import "strings"

type JWT struct {
}

// Json parse // json url decode // string split //
func (j *JWT) Decode(token string) {

	strings.Split(token, ".")

}
