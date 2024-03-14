package util

import (
	"fmt"
	"math/rand"
)

var characters = "abcdefghijklmnopqrstuvwxyz"

var currencies = []string{"USD", "INR", "EUR"}

// return random int64 between min and max
func GetRandomInt(min, max int64) int64 {
	return rand.Int63n(max-min+1) + min
}

// return random string of size n
func GetRandomString(n int64) string {
	k := len(characters)
	var str string
	var i int64
	for i = 0; i < n; i++ {
		ch := characters[rand.Intn(k)]
		str += string(ch)
	}
	return str
}

func GetRandomOwner() string {
	return GetRandomString(5)
}

func GetRandomMoney() int64 {
	return GetRandomInt(0, 1000)
}

func GetRandomCurrency() string {
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func GetRandomEmail() string {
	return fmt.Sprintf("%s@mail.com", GetRandomString(10))
}
