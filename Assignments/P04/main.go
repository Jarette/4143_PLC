/*****************************************************************************
*                    
*  Author:           Jarette Greene
*  Email:            jkgreene0406@my.msutexas.edu / jarettegreene09@gmail.com
*  Label:            P04
*  Title:			 Concurrent Image Download           
*  Course:           CMPS 4143
*  Semester:         Fall 2023
* 
*  Description:
*     
*	 Go is known for it ease of concurrency through its use of the Goroutines and 
*    go channels. In this program we show this ease of use by accepting a number of URLS 
*    of royalty free images online and using a go HTTP request to download these images
*    then this program will compare the speed by downloading each image Sequentially and 
*    Concurrently and displaying the time it takes to do both as a way to compare the effeciency 
*    of both methods and also performs necessay. 
* 
*  Usage:
*   - Place copy of image URLs in url slice 
* 	- enter go run main.go into the thr terminal 
*	- watch the image be downlaoded 
*   
* 
*  Packages:           
*        Fmt		:		formatting tools built into go
*		 strconv	:		string conversion package
*		 sync		: 		package used to create mutual exclusion for concurrency 
*		 time		:		used to time processes to check effiency 
*****************************************************************************/
// Main Package
package main

// Necessary Packages 
import (
	// format package that contains standard io methods
	"fmt"
	//used to convert string types
	"strconv"
	// contains concurrency primitives that are used for mutual exclusion
	"sync"
	// used to check the time it takes for a process to run
	"time"
	// contains methods that will can download an image from a url
	imgdownload "github.com/Jarette/4143_PLC/tree/main/Assignments/P04/ImgDownload"
)

/*
*	downloadImagesSequential
*
*	Description: takes a slice of urls and downloads each image 
*			 sequentially
*	
*	Paramenters:
*			
*			urls []string	: 	A slice of strings
*	
*	Output:
*			
*			void 		:		returns nothing
*
*/
// Sequential version of the image downloader.
func downloadImagesSequential(urls []string) {
	//name of the image file
	filename := "imageseq"
	for i := 0; i < len(urls); i++ {
		//increments name of the file so that image being downloaded does not keep getting overwritten
		filename = filename + strconv.Itoa(i) + ".jpg"
		//running function to downlaod image and also checking for error
		err := downloadImage(urls[i], filename)
		if err != nil {
			//displaying error message incase there is an error 
			fmt.Println(err)
			return
		}
		//resetting file name
		filename = "imageseq"
	}
}
/*
*	downloadImagesConcurrent
*
*	Description: takes a slice of urls and downloads each image 
*			 concurrently
*	
*	Paramenters:
*			
*			urls []string	: 	A slice of strings
*	
*	Output:
*			
*			void 		:		returns nothing
*
*/
// Concurrent version of the image downloader.
func downloadImagesConcurrent(urls []string) {
	//concurrency primitives
	var wg sync.WaitGroup
	//name of the image file
	filename := "imagecon"
	for i := 0; i < len(urls); i++ {
		// using primitives to achieve mutual exclusion
		wg.Add(1)
		// go routine that runs the download image function concurrently
		go func() {
			// incrementing file name to prevent corruption and overwritting of file names
			filename = filename + strconv.Itoa(i) + ".jpg"
			// checking for errors in downlaods
			err := downloadImage(urls[i], filename)
			if err != nil {
				//displaying error message
				fmt.Println(err)
				return
			}
			filename = "imagecon"
			wg.Done()
		}()
		// to ensure that each routine is completed before another
		wg.Wait()
	}
}

func main() {
	//slice containing urls of images to be downloaded
	urls := []string{
		"https://unsplash.com/photos/hvdnff_bieQ/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://unsplash.com/photos/HQaZKCDaax0/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://images.unsplash.com/photo-1698778573682-346d219402b5?ixlib=rb-4.0.3&q=85&fm=jpg&crop=entropy&cs=srgb&w=640",
		"https://unsplash.com/photos/Bs2jGUWu4f8/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://images.pexels.com/photos/18937801/pexels-photo-18937801/free-photo-of-wanna-play-football-or-drone.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1",
		"https://cdn.stocksnap.io/img-thumbs/960w/rose-flowers_UHDDRHM5UZ.jpg",
		"https://cdn.stocksnap.io/img-thumbs/960w/flat%20lay-thanksgiving_9BOCPBLGRV.jpg",
		"https://images.unsplash.com/photo-1699024571491-f7088816e8ba?auto=format&fit=crop&q=80&w=1935&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		"https://images.unsplash.com/photo-1699014446393-a1e0f2e15336?auto=format&fit=crop&q=80&w=1974&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		"https://images.unsplash.com/photo-1496354265829-17b1c7b7c363?q=80&w=1454&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		// Add more image URLs
	}

	// Sequential download
	//starting timer to check how long this process takes to complete
	start := time.Now()
	//downloading process
	downloadImagesSequential(urls)
	//displaying the time it took to run this process
	fmt.Printf("Sequential download took: %v\n", time.Since(start))

	// Concurrent download
	// starting timer
	start = time.Now()
	// running concurrent process
	downloadImagesConcurrent(urls)
	// displaying the time it took to run process concurrently 
	fmt.Printf("Concurrent download took: %v\n", time.Since(start))

}
/*
*	downloadImage
*
*	Description: calls method from imported package to download images using 
*				 a url that is passed in.
*	
*	Paramenters:
*			
*			urls string	: 	url of an image
*			filename	:	string used as file name
*	
*	Output:
*			
*			error 		:		error checking variables
*
*/
// Helper function to download and save a single image.
func downloadImage(url, filename string) error {
	//function call from imported package.
	imgdownload.HTTPRequest(url, filename)
	return nil
}
