package util

import (
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// init func will run automatically when call this package import --> const --> var --> init()

func init() {

	t := time.Now().UnixNano()
	rand.New(rand.NewSource(t))
}

// RandomInt generate a random int
func RandomInt(min, max int) int64 {
	return int64((max - min + 1) + min)

}

// RandomString generates a random string of length
func RandomString(length int) string {
	randomStr := make([]byte, length)

	for i := 0; i < length; i++ {

		randomStr[i] = alphabet[rand.Intn(len(alphabet))]
	}
	//Before random String value is bfapr
	// fmt.Printf("Before random String value is %v\n", string(randomStr))

	//After random String value is [98 102 97 112 114]
	// fmt.Printf("After random String value is %v\n", randomStr)
	return string(randomStr)
}

// RandomOwner return random string
func RandomOwner() string {
	return RandomString(7)
}

// RandomAmount return random amount of money
func RandomAmount() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generate random currency
func RandomCurrency() string {
	currencies := []string{"USD", "EURO", "SGD"}

	l := len(currencies)
	return currencies[rand.Intn(l)]

}
