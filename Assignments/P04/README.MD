## P04 - Concurrent Image Downloader 
### Jarette Greene
### Description

    Go is known for it ease of concurrency through its use of the Goroutines and 
    go channels. In this program we show this ease of use by accepting a number of URLS 
    of royalty free images online and using a go HTTP request to download these images
    then this program will compare the speed by downloading each image Sequentially and 
    Concurrently and displaying the time it takes to do both as a way to compare the effeciency 
    of both methods and also performs necessay. 

### Files
|   #   | File             | Description                                        |
| :---: | ---------------  | -------------------------------------------------- |
|   1   | [main.go](https://github.com/Jarette/4143_PLC/blob/main/Assignments/P04/main.go)| source code|
|   2   | [go.mod](https://github.com/Jarette/4143_PLC/blob/main/Assignments/P04/go.mod)| go mod file|
|   3   | [go.work](https://github.com/Jarette/4143_PLC/blob/main/Assignments/P04/go.work)|used to manage dependencies|
|   4   | [ImgDownload](https://github.com/Jarette/4143_PLC/tree/main/Assignments/P04/ImgDownload)|contains package with functions used to downnload an image|

### Instruction 
    - copy URL of image you would like to download 
    - place URL in slice named url 
    - open up a new terminal
    - enter go run main.go 
    - the program shall download each image

### Example Command

    





