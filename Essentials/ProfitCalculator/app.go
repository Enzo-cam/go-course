//Profit calculator
// Get revenue, expenses and tax rate from user
// Calculate earnings before TAX(EBT) and earning after TAX
// calculate ratio (EBT/profit)
// Output EBT, profit and the ratio

package main

import (
	"fmt"
)

func main() {
	const taxRate = 0.075 // Tax rate as a decimal (7.5%)

	var revenue, expenses float64

	fmt.Print("Enter the total revenue: $")
	fmt.Scan(&revenue)

	fmt.Print("Enter the total expenses for the month: $")
	fmt.Scan(&expenses)

	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate)

	var ebtToProfitRatio float64

	if profit > 0 {
		ebtToProfitRatio = earningsBeforeTax / profit
	} else {
		fmt.Println("Profit is zero or negative. Ratio calculation skipped.")
	}

	// Format output with descriptive labels and appropriate precision
	fmt.Printf("Earnings Before Tax (EBT):  $%.2f\n", earningsBeforeTax)
	fmt.Printf("Profit After Tax:          $%.2f\n", profit)

	if profit > 0 {
		fmt.Printf("Earnings Before Tax (EBT) to Profit Ratio: %.2f\n", ebtToProfitRatio)
	}
}
