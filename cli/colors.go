package cli

import "github.com/fatih/color"

var (
	Alert *color.Color = color.New(color.FgRed, color.Bold)
	Cyan *color.Color = color.New(color.FgCyan)
	Yellow *color.Color = color.New(color.FgYellow)
	Confirm *color.Color = color.New(color.FgGreen)
)