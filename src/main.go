package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"runtime"

	"agls.com.br/src/investiments-categories/savings"
	"agls.com.br/src/investiments-categories/selic"
)

func clearTerminal() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printInitialQuestions() (float64, float64, float64) {
	fmt.Println("Welcome to investiment simulator!")

	var initialContribution, monthlyContribution, investimentTime float64

	fmt.Println("Please, enter the initial contribution:")
	fmt.Scan(&initialContribution)
	fmt.Println("Please, enter the monthly contribution:")
	fmt.Scan(&monthlyContribution)
	fmt.Println("Please, enter the investiment time in years:")
	fmt.Scan(&investimentTime)

	return initialContribution, monthlyContribution, investimentTime
}

func printGraph(selicValue, savingValue float64) {
	maxValue := math.Max(selicValue, savingValue)
	scale := 5000 / maxValue

	selicBars := int(selicValue * scale)
	savingBars := int(savingValue * scale)

	fmt.Println("Investment Future Value Graph:")
	fmt.Printf("Selic:   |%s\n", string('#'*selicBars))
	fmt.Printf("Savings: |%s\n", string('#'*savingBars))
}

func main() {
	clearTerminal()
	initialContribution, monthlyContribution, investimentTime := printInitialQuestions()

	selic := selic.New(initialContribution, monthlyContribution, investimentTime)
	selic.Calculate()
	saving := savings.New(initialContribution, monthlyContribution, investimentTime)
	saving.Calculate()

	fmt.Println("The future value of your investment will be: ")
	fmt.Printf("Selic: R$ %.2f\n", selic.FutureValue)
	fmt.Printf("Savings: R$ %.2f\n", saving.FutureValue)

	printGraph(selic.FutureValue, saving.FutureValue)
}
