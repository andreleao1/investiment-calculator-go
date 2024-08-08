package main

import (
	"fmt"

	"agls.com.br/investiments-categories/savings"
	"agls.com.br/investiments-categories/selic"
	"agls.com.br/output"
)

func main() {
	output.ClearTerminal()
	initialContribution, monthlyContribution, investimentTime := output.PrintInitialQuestions()

	selic := selic.New(initialContribution, monthlyContribution, investimentTime)
	selic.Calculate()
	saving := savings.New(initialContribution, monthlyContribution, investimentTime)
	saving.Calculate()

	fmt.Println("The future value of your investment will be: ")
	fmt.Printf("Selic: R$ %.2f\n", selic.FutureValue)
	fmt.Printf("Savings: R$ %.2f\n", saving.FutureValue)

	output.PrintGraph(selic.FutureValue, saving.FutureValue)
}
