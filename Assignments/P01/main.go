// Package main is the Main Driver
package main

// importing neccesary packages
import (
	//Package fmt used to output on the screen
	"fmt"
	// created package to display mascot name
	"github.com/4143_PLC/P01/mascot"
	// Package Quote that generates a quote
	"rsc.io/quote"
)

// Main Driver function
func main() {
	// displaying the name of the best mascot on the screen
	fmt.Println(mascot.BestMascot())
	// displaying a quote to the screen
	fmt.Println(quote.Go())
}
