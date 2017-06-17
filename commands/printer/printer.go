package printer

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	red   = color.New(color.FgRed).Add(color.Bold).SprintFunc()
	white = color.New(color.Bold).SprintFunc()
)

// PrintWarning prints provided label in standard font and provided msg in bold red.
func PrintWarning(label, msg string) {
	fmt.Printf("%s %s\n", label, red(msg))
}

// PrintInfo prints provided label in standard font and provided msg in bold white.
func PrintInfo(label, msg string) {
	fmt.Printf("%s %s\n", label, white(msg))
}
