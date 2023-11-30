/*****************************************************************************
*                    
*  Author:           Jarette Greene
*  Email:            jkgreene0406@my.msutexas.edu / jarettegreene09@gmail.com
*  Label:            P03
*  Title:			 Image Ascii Art           
*  Course:           CMPS 4143
*  Semester:         Fall 2023
* 
*  Description:

*     This Project was to demonstrate the use of the "go get" command
*    by creating 4 different packages and uploading them to individual
*    repositories and importing them using the "go get" command to import 
*    these packages into a singular main.go file and use all these packages
*    to manipulate a downloaded image and user entered text.
* 
*  Usage:
*    - using "go get" functions to import the neccesary packages 
*	- use the repo urls to access the functions 
*	- in terminal write "go run main.go" to execute program
*	- follow prompts on the screen 
* 
*  Packages:           
*        Img_color	:	 gets RGB values
*		 Img_get	:	 downloads image from the internet
*        Img_gray_scale      :  grayscales downloaded image
*        Img_Text    : displays colored image and creates text png
*****************************************************************************/
package main

import (
	//fmt package: allows for input and output from and to the console respectively
	"fmt"

	// github.com/Jarette/Img_colors package: allows to get the RGB values of an image
	"github.com/Jarette/Img_colors"

	// github.com/Jarette/Img_get package: contains function to download an image using URL
	"github.com/Jarette/Img_get"

	// github.com/Jarette/Img_gray_scale: contains functions to apply grayscale to an image
	"github.com/Jarette/Img_gray_scale"

	// github.com/Jarette/Img_text : contains fucntions to print colored text to screen 
	"github.com/Jarette/Img_text"
)

// ImgMan Image Manipulator object
type ImgMan struct {
	URL  string
	Text string
}
/*
*	MakeGray
*
*	Description: Method used to take an image and make it grayscale
*	
*	Paramenters:
*			
*			n/a
*	
*	Output:
*			
*			void 		:		returns nothing
*
*/
// MakeGray method to make grayscale
func (i ImgMan) MakeGray() {
	Img_get.Getpic(i.URL)
	Img_gray_scale.Grayscale()
}
/*
*	Imgtxt
*
*	Description: places image on colored text on an image
*	
*	Paramenters:
*			
*			n/a
*	
*	Output:
*			
*			void 		:		returns nothing
*
*/
// Imgtxt creates image of text
func (i ImgMan) Imgtxt() {
	Img_text.Color_imgtext(i.Text)
}

// Colortext displays color text to the terminal
/*
*	Colortext
*
*	Description: prints colored text to terminal
*	
*	Paramenters:
*			
*			n/a
*	
*	Output:
*			
*			void 		:		returns nothing
*
*/
func (i ImgMan) Colortext() {
	Img_text.Screen_text(i.Text)
}

// DisplayPixels displays the RGB values of every pixel in downloaded image
/*
*	DisplayPixels
*
*	Description: Displays the RGB values of an image
*	
*	Paramenters:
*			
*			n/a
*	
*	Output:
*			
*			void 		:		returns nothing
*
*/
func (i ImgMan) DisplayPixels() {
	Img_get.Getpic(i.URL)
	Img_colors.Colors()
}

// SetURL setter method for URL
/*
*	SetURL
*
*	Description: setter function to get the url 
*	
*	Paramenters:
*			
*			url string 	: 	the url of the image to be manipulated 
*	
*	Output:
*			
*			void 		:		returns nothing
*
*/
func (i ImgMan) SetURL(url string) {
	i.URL = url
}

// SetText setter method for text
/*
*	SetText
*
*	Description: setter function to get the text
*	
*	Paramenters:
*			
*			url text 	: 	the text to be put on the image 
*	
*	Output:
*			
*			void 		:		returns nothing
*
*/
func (i ImgMan) SetText(text string) {
	i.Text = text
}
func main() {
	// variables
	var choice int
	var url string
	var text string
	// first choice for the user to download an image and manipulate it or just manipulate text
	// of thier choice
	fmt.Println("Enter 1. to Manipulate Image. Enter 2. to Manipulate text")
	fmt.Scan(&choice)
	if choice == 1 {
		// getting the URL. NB: I thought that you can copy and paste into the terminal but you cant so you
		// may have to type of the URL.
		println("Enter URL for the image you would like to download ")
		fmt.Scan(&url)
		// initializing ImgMan object with a URL
		image := &ImgMan{URL: url}
		// Picking if you would like the Graysclae or RGB values
		println("Enter 1. to use grayscale, Enter 2. to get the RGB of every pixel")
		fmt.Scan(&choice)
		if choice == 1 {
			//downloads an image from the internet using a url and grayscales that image 
			image.MakeGray() 
		} else if choice == 2 {
			// downloads an image from the internet using a url and displays the RGB values 
			// of each pixels 
			image.DisplayPixels() 
		} else {
			// incase of invalid input
			fmt.Println("Error")
			return
		}
		//similar to previous section but manipulating user entered text
	} else if choice == 2 {
		println("Enter the text you would like to manipulate ")
		fmt.Scan(&text)
		// initializing ImgMan object with a Text
		image := &ImgMan{Text: text}
		println("Enter 1. to use print to screen, Enter 2. print to a png")
		fmt.Scan(&choice)
		if choice == 1 {
			// displays text stored in the image object to the screen in color
			image.Colortext()
		} else if choice == 2 {
			// displays text stored in the image object to a png.
			image.Imgtxt()
		} else {
			fmt.Println("Error")
			return
		}
	} else {
		fmt.Println("Error")
		return
	}
}
