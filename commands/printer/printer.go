package printer

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	red   = color.New(color.FgRed).Add(color.Bold).SprintFunc()
	white = color.New(color.Bold).SprintFunc()
)

func PrintWarning(label, msg string) {
	fmt.Printf("%s %s\n", label, red(msg))
}

func PrintInfo(label, msg string) {
	fmt.Printf("%s %s\n", label, white(msg))
}
