/*****************************************************************************
*                    
*  Author:           Jarette Greene
*  Email:            jkgreene0406@my.msutexas.edu / jarettegreene09@gmail.com
*  Label:            P02
*  Title:            Baby Steps
*  Course:           CMPS 4143
*  Semester:         Fall 2023
* 
*  Description:
*		
*		This program used to show how to locally manage package files in a  
*		Go workspace. This program also shows the extensive libraries that  
*		available in Go Lang by using some basic image Manipulation functions 
*		import from a package taken from the internet then used to perform
*		simple image manipulation on a presaved image and resave said image 
*		with a rectangle drawn over it. This program also displays Go Lang's 
*		version of constructors.
* 
*  Usage:
*       - make sure you have image saved by approipriate name saved in files
*		- run go run main.go to run program  
*		
*  Files:           
*      go.mod			:		mod file to handle dependencies
*	   go.sum			:		controls versioning of different packages
*	   mustang.jpg		:		Image to be manipulated
*	   mustang.png 		:		Image after manipulation 
*	   imageManipulator :		package used to manipulate the saved image
*****************************************************************************/
//Main Driver Package
package main

//necessay imports
import (
	// fmt package : allows to input and output from and to the screen respectively 
	"fmt"

	// myimageapp/imageManipulator: local package that contains methods to draw a rectangle on an image
	"myimageapp/imageManipulator"
)

func main() {
	/**
	Code that simply creates a PNG of a rectangle 
	// Only Draws a rectangle on a PNG 
	// // Create an ImageManipulator instance
	// im := imageManipulator.NewImageManipulator(800, 600)

	// // Draw a rectangle
	// im.DrawRectangle(150, 50, 560, 411)

	// // Save the image to a file
	// im.SaveToFile("mustangs.png")
	*/

	// Draws a rectangle on an image pulled from files 
	// Create an ImageManipulator instance with an existing image
    im, err := imageManipulator.NewImageManipulatorWithImage("mustangs.jpg")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Draw a rectangle
    im.DrawRectangle(150, 50, 560, 411)

    // Save the modified image to a file
    im.SaveToFile("mustangs.png")
}