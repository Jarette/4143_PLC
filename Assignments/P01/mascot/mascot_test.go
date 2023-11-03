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
*		masoct tester file that test if the masoct package
*		returns the correct masoct name
*
* 
*  Usage:
*       -click run test button 
*		
*  Files:           
*     mascot.go 	: contains the Best mascot function to be tested
*****************************************************************************/
// Mascot Test File
package mascot_test

// importing test package and the mascot package
import (
	//import package to allow go to test package functions
	"testing"

	//importing the BestMascot function 
	"github.com/4143_PLC/P01/mascot"
)

// TestMascot function tests the mascot package
func TestMascot(t *testing.T) {
	// checks if the mascot package's "BestMascot" function
	// returns the correct output
	if mascot.BestMascot() != "Go Gopher" {
		//display this message if the wrong output is produced
		t.Fatal("Wrong Mascot")
	}
}
