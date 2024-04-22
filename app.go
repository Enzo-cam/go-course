// This indicate that this PACKAGE will be the main entry point of the APP
// A module consist of multiple packages
// To create a module we use 'go mod init ${PATH}'

package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount, years float64 = 1000, 10
	expectedRetunRate := 5.5

	futureValue := investmentAmount * math.Pow(1+expectedRetunRate/100, years)

	fmt.Println(futureValue)
}
