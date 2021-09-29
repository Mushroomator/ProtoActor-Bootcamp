package consolepainter

import (
	"fmt"
	"runtime"
)

type ConsolePainter struct{}

// Colors for bash (DOES NOT WORK ON WINDOWS!)
var (
	Reset  string = "\033[0m"
	Red    string = "\033[31m"
	Green  string = "\033[32m"
	Yellow string = "\033[33m"
	Blue   string = "\033[34m"
	Purple string = "\033[35m"
	Cyan   string = "\033[36m"
	Gray   string = "\033[37m"
	White  string = "\033[97m"
)

// Create a console painter which allows to print messages in different colors to the console.
func NewConsolePainter() *ConsolePainter {
	Init()
	return &ConsolePainter{}
}

// Windows console does not support bash escape sequences to color output but would instead print out the values.
// If on windows set all colors to empty string so no unwanted output is generated.
func Init() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""
	}
}

// Print yellow message
func (cp *ConsolePainter) PrintfYellow(msg string, vals ...interface{}) {
	fmt.Printf(Yellow+msg+Reset, vals...)
}

// Print red message
func (cp *ConsolePainter) PrintfRed(msg string, vals ...interface{}) {
	fmt.Printf(Red+msg+Reset, vals...)
}

// Print green message
func (cp *ConsolePainter) PrintfGreen(msg string, vals ...interface{}) {
	fmt.Printf(Green+msg+Reset, vals...)
}

// Print blue message
func (cp *ConsolePainter) PrintfBlue(msg string, vals ...interface{}) {
	fmt.Printf(Blue+msg+Reset, vals...)
}

// Print purple message
func (cp *ConsolePainter) PrintfPurple(msg string, vals ...interface{}) {
	fmt.Printf(Purple+msg+Reset, vals...)
}

// Print cyan message
func (cp *ConsolePainter) PrintfCyan(msg string, vals ...interface{}) {
	fmt.Printf(Cyan+msg+Reset, vals...)
}

// Print gray message
func (cp *ConsolePainter) PrintfGray(msg string, vals ...interface{}) {
	fmt.Printf(Gray+msg+Reset, vals...)
}

// Print white message
func (cp *ConsolePainter) PrintfWhite(msg string, vals ...interface{}) {
	fmt.Printf(White+msg+Reset, vals...)
}
