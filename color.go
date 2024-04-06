package main

import (
	"fmt"
)

const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorTeal   = "\033[36m"
	color1      = "\033[37m"
	colorReset  = "\033[0m"
)

func printRed(s string) {
	fmt.Printf("%s%s%s\n", colorRed, s, colorReset)
}
func printRedErr(e error) {
	fmt.Printf("%s%v%s\n", colorRed, e, colorReset)
}
func printGreen(s string) {
	fmt.Printf("%s%s%s\n", colorGreen, s, colorReset)
}
func printYellow(s string) {
	fmt.Printf("%s%s%s\n", colorYellow, s, colorReset)
}
func printCyan(s string) {
	fmt.Printf("%s%s%s\n", colorTeal, s, colorReset)
}
func printPurple(s string) {
	fmt.Printf("%s%s%s\n", colorPurple, s, colorReset)
}
func printBlue(s string) {
	fmt.Printf("%s%s%s\n", colorBlue, s, colorReset)
}
