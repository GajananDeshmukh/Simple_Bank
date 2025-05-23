package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().Unix()))
}

// RandomInt generates a random integrer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) //min->max-min+min

}

// Randomstring generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()

}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)

}

func RandomMoney() int64 {
	return RandomInt(0, 1000)

}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]

}
