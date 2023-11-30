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
*	 This package contains methods that will take a url of an image and downloads 
*	 the image locally
* 
*  Usage:
*
*    call the HTTPRequest and pass in a url and a file name to download an image 
*   
* 
*  Packages:           
*        Fmt		:		formatting tools built into go
*		 net/http	:		contains methods that are used downlaod images
*		 io			: 		contains io primitives
*		 os			:		used in error handling 
*****************************************************************************/
package imgdownload
import (
	// cointaings tools for basic io
    "fmt"
	// used to download images using uls
    "net/http"
	// contains io primitives
    "io"
	// used in error handling 
    "os"
)
/*
*	HTTPRequest
*
*	Description: takes a url and downloads it locally 
*	
*	Paramenters:
*			
*			urls string	: 	url of image to be downloaded
*			filename string :  name of the image file
*							
*	Output:
*			
*			void 		:		returns nothing
*
*/
//HTTPRequest downloads an image from the internet using a url
func HTTPRequest(url, filename string){
	// getting the image from the url and checkign for error
	req, err := http.NewRequest("GET",url,nil)
	if err != nil{
		fmt.Println(err)
		return
	}

	// performing http request
	client:= &http.Client{}

	resp,err:= client.Do(req)
	if err != nil{
		fmt.Println(err)
		return
	}
	// making sure the status of request
	if resp.StatusCode != http.StatusOK{
		fmt.Println("Response status code:", resp.StatusCode)
		return
	}
	// creating image file
	f,err :=os.Create(filename)
	if err != nil{
		fmt.Println(err)
		return
	}
	// placing image on file and checking if there is an error in copy 
	_, err = io.Copy(f,resp.Body)
	if err != nil{
		fmt.Println(err)
		return
	}
	// closing file
	f.Close()
	//indicating that the file was downloaded sucesfully 
	fmt.Println("Image saved to", filename)
}