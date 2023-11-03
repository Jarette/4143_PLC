/*****************************************************************************
*                    
*  Author:           Jarette Greene
*  Email:            jkgreene0406@my.msutexas.edu / jarettegreene09@gmail.com
*  Label:            P01
*  Title:            Run a Go Program 
*  Course:           CMPS 4143
*  Semester:         Fall 2023
* 
*  Description:
*		
*		This program creates a simple go workspace by having the user create 
*		an appropriate go module (go.mod file). Then it has the user create a 
*		simple package, that displays the name of the Golang mascot (the Go Gopher) 
*		with its approipriate test file, then save it locally and import it into the
*		the main go file. This Program also imports a package from the internet using 
*		the go get function to import a package that will display a quote. 
*
* 
*  Usage:
*       - use go run main.go command in console to run 
*		
*  Files:           
*      go.mod		:		mod file to handle dependencies
*	   go.sum		:		controls versioning of different packages
*	   mascot		:		contains mascot package saved locally
*****************************************************************************/
// Package main is the Main Driver
package main

// importing neccesary packages
import (
	//Package fmt used to output on the screen
	"fmt"
	// imported package that with fucntions to display the GO gopher mascot 
	"github.com/4143_PLC/P01/mascot"
	// Imported package that returns a quote
	"rsc.io/quote"
)

// Main Driver function
func main() {
	// displaying the name of the best mascot on the screen
	fmt.Println(mascot.BestMascot())
	// displaying a quote to the screen
	fmt.Println(quote.Go())
}
