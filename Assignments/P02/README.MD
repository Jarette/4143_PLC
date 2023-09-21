## P02 - Baby Steps
### Jarette Greene
### Description

    In this Project we perform some image manipulation on a presaved image
    and also demonstrate how to create a psuedo class in GOLANG. First we create 
    a package called imageManipulator that contains a struct with a method that
    functions as a constructor. We used this package in our image module to create 
    a png of a rectangle then we updated the imageManipulator struct attributes and
    methods so that we take a saved image and place that rectangle on top of that image

### Files
|   #   | File             | Description                                        |
| :---: | ---------------  | -------------------------------------------------- |
|   1   | [imagemod](https://github.com/Jarette/4143_PLC/tree/main/Assignments/P02/imagemod)| Image Module|
|   2   | [main.go](https://github.com/Jarette/4143_PLC/tree/main/Assignments/P02/imagemod/main.go)| Source code|
|   3   | [imageManipulator](https://github.com/Jarette/4143_PLC/tree/main/Assignments/P02/imagemod/imageManipulator)| imageManipulator Package|
|   4   | [imageManipulator.go](https://github.com/Jarette/4143_PLC/blob/main/Assignments/P02/imagemod/imageManipulator/imageManipulator.go)|imageManipulator package code|

### Instruction 

- in terminal write "go run main.go" to execute program

### Example Command

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

   

