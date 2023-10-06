package main

import (
	"fmt"

	"github.com/Jarette/Img_colors"
	"github.com/Jarette/Img_get"
	"github.com/Jarette/Img_gray_scale"
	"github.com/Jarette/Img_text"
)

// ImgMan Image Manipulator object
type ImgMan struct {
	URL  string
	Text string
}

// MakeGray method to make grayscale
func (i ImgMan) MakeGray() {
	Img_get.Getpic(i.URL)
	Img_gray_scale.Grayscale()
}

// Imgtxt creates image of text
func (i ImgMan) Imgtxt() {
	Img_text.Color_imgtext(i.Text)
}

// Colortext displays color text to the terminal
func (i ImgMan) Colortext() {
	Img_text.Screen_text(i.Text)
}

// DisplayPixels displays the RGB values of every pixel in downloaded image
func (i ImgMan) DisplayPixels() {
	Img_colors.Colors()
}

// SetURL setter method for URL
func (i ImgMan) SetURL(url string) {
	i.URL = url
}

// SetText setter method for text
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
		// initializing ImgMan object
		image := &ImgMan{URL: url}
		// Picking if you would like the Graysclae or RGB values
		println("Enter 1. to use grayscale, Enter 2. to get the RGB of every pixel")
		fmt.Scan(&choice)
		if choice == 1 {
			image.MakeGray()
		} else if choice == 2 {
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
		image := &ImgMan{Text: text}
		println("Enter 1. to use print to screen, Enter 2. print to a png")
		fmt.Scan(&choice)
		if choice == 1 {
			image.Colortext()
		} else if choice == 2 {
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
