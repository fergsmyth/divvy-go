package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {

	payments := paymentInformation()
	totalPaid := totalOwed(payments)
	payerCount := payerCount(payments)
	simpleAverage := averageOwed(totalPaid, int64(payerCount))
	paymentDelta(simpleAverage, payments)
}

func paymentInformation() map[string]decimal.Decimal {

	peoplePaidAmount := make(map[string]decimal.Decimal)

	for {
		var name string
		fmt.Println("Input Name or -1 to exit:")
		fmt.Scanln(&name)
	
		if name == "-1"{
			break
		}

		var paid string 
		fmt.Println("Enter amount paid")
		fmt.Scanln(&paid)

		paidDecimal, err := decimal.NewFromString(paid)
	
		if err != nil {
			fmt.Println("Not a valid number")
			continue
		}

		peoplePaidAmount[name] = paidDecimal 
	}

	return peoplePaidAmount
}

func totalOwed(payments map[string]decimal.Decimal) decimal.Decimal {

	var totalPaid decimal.Decimal
	for _, val := range payments {
		totalPaid = totalPaid.Add(val) 
	} 
	fmt.Printf("Total paid : %v\n", totalPaid.String())
	return totalPaid
} 

func payerCount(payments map[string]decimal.Decimal) int {
	peopleCount := len(payments)
        fmt.Printf("Number of payers : %v\n", peopleCount)
	return peopleCount
}

func averageOwed(totalPaid decimal.Decimal, payerCount int64) decimal.Decimal {
	simpleAverage := totalPaid.Div(decimal.NewFromInt(payerCount))
        fmt.Printf("Each owes : %v\n", simpleAverage)
	return simpleAverage
}

func paymentDelta(simpleAverage decimal.Decimal, payments map[string]decimal.Decimal) map[string]decimal.Decimal {
	delta := make(map[string]decimal.Decimal)

	for key, val := range payments {
		deltaVal := val.Sub(simpleAverage)
		delta[key] = deltaVal
		fmt.Printf("%v is owed %v\n", key, deltaVal)		
	}
	return delta
}
