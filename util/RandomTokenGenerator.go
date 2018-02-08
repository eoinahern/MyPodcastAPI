package util

import "github.com/dchest/uniuri"

func GenerateRandomToken() string {
	return uniuri.NewLen(10)
}
