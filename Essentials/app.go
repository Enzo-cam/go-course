// This indicate that this PACKAGE will be the main entry point of the APP
// A module consist of multiple packages
// To create a module we use 'go mod init ${PATH}'

package mainEssential

import (
	"fmt"
	"math"
)

func mainEssential() {
	const inflationRate = 3.9
	var investmentAmount, years float64
	expectedRetunRate := 5.5

	// With the &, we indicate that we have to point in that variable and is going to have the value that the user inputs
	fmt.Print("Invest Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Years that you are going to expect: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedRetunRate/100, years)
	futureValueInflation := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureValueInflation)
}
