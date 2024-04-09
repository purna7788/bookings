package helpers

import "math/rand"

func RandomNum(n int) int {

	myVar := rand.Intn(n)
	return myVar
}