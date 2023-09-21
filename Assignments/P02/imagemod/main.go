package main

import (
	"fmt"
	"myimageapp/imageManipulator"
)

func main() {
	
	// Only Draws a rectangle on a PNG 
	// // Create an ImageManipulator instance
	// im := imageManipulator.NewImageManipulator(800, 600)

	// // Draw a rectangle
	// im.DrawRectangle(150, 50, 560, 411)

	// // Save the image to a file
	// im.SaveToFile("mustangs.png")


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