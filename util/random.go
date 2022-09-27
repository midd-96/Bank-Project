package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "qwertyuiopasdfghjklzxcvbnm"

func init() {
	rand.Seed(time.Now().UnixNano())
}

//RandomInt generates a random integer
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

//RandomString generates
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

//RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

//RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

//RandomCurrency select a currency mode

func RandomCurrency() string {
	currencies := []string{"EUR", "USD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
