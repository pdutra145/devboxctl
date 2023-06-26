package utils

import "github.com/fatih/color"

var (
	Alert *color.Color = color.New(color.FgRed, color.Bold)
	Special *color.Color = color.New(color.FgCyan)
	Warning *color.Color = color.New(color.FgYellow)
	Success *color.Color = color.New(color.FgGreen, color.Bold)
)