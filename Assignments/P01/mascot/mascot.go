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
*		This is the mascot package that will return the name of Go Lang
*		mascot , the Go Gopher.
*
* 
*  Usage:
*       -import into main using go get function
*		
*  Files:           
*     N/A
*****************************************************************************/
// Package mascot contains methods to display mascot name
package mascot
/**
* BestMascot 
*
* 	Description 
*	 	returns the name of the Go Lang masoct
*
*	Param: 
*		N/A
*	
*	Returns:
*		string : mascot name 
*		
*/
//BestMascot Returns the name of the best mascot
func BestMascot() string {
	return "Go Gopher"
}
