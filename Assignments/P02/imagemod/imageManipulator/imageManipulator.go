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
*		imageManipulator package that imports a package from the internet 
*       with function that allow you to perform simple edits of an image 
*       in this example we are simply placing a rectangle on an image. This 
*       all uses Go's "New" naming convention to create Go's version of a 
*       constructor.
* 
*  Usage:
*       N/A 
*		
*  Files:           
*       N/A
*****************************************************************************/
// imagemod/imageManipulator/imageManipulator.go

package imageManipulator

import (
    "github.com/fogleman/gg"
)

// Old struct

 //ImageManipulator represents an image manipulation tool.
// type ImageManipulator struct {
//     Image *gg.Context
// }


//updating ImageManipulator struct
type ImageManipulator struct {
    Image *gg.Context
    ImagePath string // Add a field to store the image path
}

//Old consturctor 

// NewImageManipulator creates a new ImageManipulator instance.
// func NewImageManipulator(width, height int) *ImageManipulator {
//   img := gg.NewContext(width, height)
//    return &ImageManipulator{Image: img}
// }


// Updating constructor.

//NewImageManipulatorWithImage : finds the image in your system and creates for the program 
/*
*	NewImageManipulatorWithImage
*
*	Description: Finds an image saved on system and opens it so that it 
*                can be manipulated 
*	
*	Paramenters:
*			
*			imagePath  	: 	The name of the image
*	
*	Output:
*			
*			error 		:		returns an error if it does not succeed 
*
*/
func NewImageManipulatorWithImage(imagePath string) (*ImageManipulator, error) {
    //Creating the path to the image 
    img, err := gg.LoadImage(imagePath)
    if err != nil {
        //displaying error image 
        return nil, err
    }

    // preparing the image to be chaged and rendered 
    ctx := gg.NewContextForImage(img)
    return &ImageManipulator{Image: ctx, ImagePath: imagePath}, nil
}

// SaveToFile saves the manipulated image to a file.
/*
*	SaveToFile
*
*	Description: saves the manipulated image to your system as a PNG
*	
*	Paramenters:
*			
*			filename string  	: 	name of the file
*	
*	Output:
*			
*			error 		:		returns an error if it does not succeed 
*
*/
func (im *ImageManipulator) SaveToFile(filename string) error {
    return im.Image.SavePNG(filename)
}

// DrawRectangle draws a rectangle on the image.
/*
*	DrawRectangle
*
*	Description: places a rectangle on saved image 
*	
*	Paramenters:
*			
*			x,y,width,height float64  	: 	the size of the rectangle
*	
*	Output:
*			
*			void  		:		No Return 
*
*/
func (im *ImageManipulator) DrawRectangle(x, y, width, height float64) {
    im.Image.DrawRectangle(x, y, width, height)
    im.Image.Stroke()
}