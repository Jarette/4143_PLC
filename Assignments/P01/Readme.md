## P01 - Dev Enviroment
### Jarette Greene
### Description

    In this assignment I set up a go module and workspace 
    then used that work space to create my own package that
    displays the name of the GoLang mascot "Go Gohper" and test
    package. I also installed a package from the internet to display
    a quote

    ### Files
|   #   | File             | Description                                        |
| :---: | ---------------  | -------------------------------------------------- |
|   1   | [main.go](https://github.com/Jarette/4143_PLC/blob/main/Assignments/P01/main.go)| Source code|
|   2   | [go.mod](https://github.com/Jarette/4143_PLC/blob/main/Assignments/P01/go.mod)| go module file|
|   3   | [go.sum](https://github.com/Jarette/4143_PLC/blob/main/Assignments/P01/go.sum)| for version control|
|   4   | [mascot](https://github.com/Jarette/4143_PLC/tree/main/Assignments/P01/mascot)|contains mascot package and mascot testing file|

### Instruction 

- in terminal write "go run main.go" to execute program

### Example Command

    func main() {
	// displaying the name of the best mascot on the screen
	fmt.Println(mascot.BestMascot())
	// displaying a quote to the screen
	fmt.Println(quote.Go())
  }

