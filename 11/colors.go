package main

import "fmt"

type Color string

const (
	Black       Color = "\033[0;30m"
	Red         Color = "\033[0;31m"
	Green       Color = "\033[0;32m"
	BrownOrange Color = "\033[0;33m"
	Blue        Color = "\033[0;34m"
	Purple      Color = "\033[0;35m"
	Cyan        Color = "\033[0;36m"
	LightGray   Color = "\033[0;37m"
	DarkGray    Color = "\033[1;30m"
	LightRed    Color = "\033[1;31m"
	LightGreen  Color = "\033[1;32m"
	Yellow      Color = "\033[1;33m"
	LightBlue   Color = "\033[1;34m"
	LightPurple Color = "\033[1;35m"
	LightCyan   Color = "\033[1;36m"
	White       Color = "\033[1;37m"
	Plain       Color = "\033[0m"
)

func asColor(s string, c Color) string {
	return fmt.Sprintf("%v%v%v", c, s, Plain)
}
