package powxn

import "fmt"

var c = 1

func myPow(x float64, n int) float64 {
	c++
	fmt.Println(c)
	if n == 0 {
		return 1
	}

	if n < 0 {
		n *= -1
		x = 1 / x
	}

	if n%2 == 0 {
		return myPow(x*x, n/2)
	}

	return x * myPow(x*x, n/2)
}
