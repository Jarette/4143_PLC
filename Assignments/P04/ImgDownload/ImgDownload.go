package imgdownload
import (
    "fmt"
    "net/http"
    "io"
    "os"
)

func HttpRequest(url, filename string){
	req, err := http.NewRequest("GET",url,nil)
	if err != nil{
		fmt.Println(err)
		return
	}

	client:= &http.Client{}

	resp,err:= client.Do(req)
	if err != nil{
		fmt.Println(err)
		return
	}

	if resp.StatusCode != http.StatusOK{
		fmt.Println("Response status code:", resp.StatusCode)
		return
	}

	f,err :=os.Create(filename)
	if err != nil{
		fmt.Println(err)
		return
	}

	_, err = io.Copy(f,resp.Body)
	if err != nil{
		fmt.Println(err)
		return
	}

	f.Close()

	fmt.Println("Image saved to", filename)
}