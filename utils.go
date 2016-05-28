package token_generator

import (
	"bytes"
	"math/rand"
)

// RandomInt function is used to fetch a random integer value in a range (min, max)
func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

/*
RandomString function is string generator which makes use of the "math/rand" library.

Note: If you want to use this solo, it is suggested that you seed the "rand" library
	rand.Seed(time.Now().Unix())
*/
func RandomString(length int) string {
	var buffer bytes.Buffer
	for i := 0; i < length; i++ {
		choice := RandomInt(0, 2)
		switch choice {
		case 0:
			buffer.WriteString(string(RandomInt(48, 57)))
		case 1:
			buffer.WriteString(string(RandomInt(97, 122)))
		case 2:
			buffer.WriteString(string(RandomInt(65, 90)))

		}
	}
	return buffer.String()
}
