package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

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

func main() {
	clearTerminal()

	var initialContribution, monthlyContribution, investimentTime float64

	fmt.Println("Welcome to investiment simulator!")

	fmt.Println("Please, enter the initial contribution:")
	fmt.Scan(&initialContribution)
	fmt.Println("Please, enter the monthly contribution:")
	fmt.Scan(&monthlyContribution)
	fmt.Println("Please, enter the investiment time in years:")
	fmt.Scan(&investimentTime)

	selic := selic.New(initialContribution, monthlyContribution, investimentTime)
	selic.Calculate()

	fmt.Printf("The future value of your investment is R$: %.2f", selic.FutureValue)

}
