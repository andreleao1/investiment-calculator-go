package output

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func ClearTerminal() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func PrintInitialQuestions() (float64, float64, float64) {
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

func PrintGraph(selicValue, savingValue float64) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}

	defer ui.Close()

	formattedSelicValue, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", selicValue), 64)
	formattedSavingValue, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", savingValue), 64)

	barchart := widgets.NewBarChart()
	barchart.Title = "Investment Future Value Graph"
	barchart.Labels = []string{"Selic", "Savings"}
	barchart.Data = []float64{formattedSelicValue, formattedSavingValue}
	barchart.BarWidth = 13
	barchart.BarColors = []ui.Color{ui.ColorGreen, ui.ColorYellow}
	barchart.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorWhite)}
	barchart.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}

	barchart.SetRect(0, 0, 50, 10)

	paragraph := widgets.NewParagraph()
	paragraph.Text = "Press space to exit"
	paragraph.SetRect(0, 20, 50, 15)
	paragraph.TextStyle.Fg = ui.ColorWhite

	ui.Render(barchart, paragraph)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			if e.ID == "<Space>" {
				break
			}
		}
	}
}
