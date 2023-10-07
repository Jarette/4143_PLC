## P03 - Image Ascii Art
### Jarette Greene
### Description

    This Project was to demonstrate the use of the "go get" command
    by creating 4 different packages and uploading them to individual
    repositories and importing them using the "go get" command to import 
    these packages into a singular main.go file and use all these packages
    to manipulate a downloaded image and user entered text.

### Files
|   #   | File             | Description                                        |
| :---: | ---------------  | -------------------------------------------------- |
|   1   | [main.go](https://github.com/Jarette/4143_PLC/blob/main/Assignments/P03/main.go)| source code|
|   2   | [go.mod](https://github.com/Jarette/4143_PLC/blob/main/Assignments/P03/go.mod)| go mod file|
|   3   | [go.sum](https://github.com/Jarette/4143_PLC/blob/main/Assignments/P03/go.sum)|package version managing file|
|   4   | [Img_get](https://github.com/Jarette/Img_get)|Img_get package repo|
|   5   | [Img_colors](https://github.com/Jarette/Img_colors)| Img_color package repo|
|   6   | [Img_text](https://github.com/Jarette/Img_text)| Img_text package repo|
|   7   | [Img_gray_scale](https://github.com/Jarette/Img_gray_scale)| Img_gray_scale package repo|
|   8   | [downloaded_image.jpg](https://github.com/Jarette/4143_PLC/blob/main/Assignments/P03/downloaded_image.jpg)|image downloaded using a url|
|   9   | [gray_image.png](https://github.com/Jarette/4143_PLC/blob/main/Assignments/P03/gray_image.png)|the grayscaled version of the image|
|   10  | [hello.png](https://github.com/Jarette/4143_PLC/blob/main/Assignments/P03/hello.png)|text displayed placed on png|

### Instruction 
- using "go get" functions to import the neccesary packages 
- use the repo urls to access the functions 
- in terminal write "go run main.go" to execute program
- follow prompts on the screen 

### Example Command

    // ImgMan Image Manipulator object
    type ImgMan struct {
	URL  string
	Text string
    }

    image := &ImgMan{URL: url} //overloaded constructor to fill with the url of the image
    image := &ImgMan{Text: text} //overloaded constructor to fill with the text to place on image or screen
    image.MakeGray() // grayscales an image downloaded from the previous url
    image.DisplayPixels() // displays the RGB values of the downloaded image
    image.Colortext() // displays color text on the screen thats saved in the image object
    image.Imgtxt() // places text from image object to a png






