package main

import (
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

	output.PrintGraph(selic.FutureValue, saving.FutureValue)
}
