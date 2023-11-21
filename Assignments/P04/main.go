package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	imgdownload "github.com/Jarette/4143_PLC/tree/main/Assignments/P04/ImgDownload"
)

// Sequential version of the image downloader.
func downloadImagesSequential(urls []string) {
	filename := "imageseq"
	for i := 0; i < len(urls); i++ {
		filename = filename + strconv.Itoa(i) + ".jpg"
		err := downloadImage(urls[i], filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		filename = "imageseq"
	}
}

// Concurrent version of the image downloader.
func downloadImagesConcurrent(urls []string) {
	var wg sync.WaitGroup
	filename := "imagecon"
	for i := 0; i < len(urls); i++ {
		wg.Add(1)
		go func() {
			filename = filename + strconv.Itoa(i) + ".jpg"
			err := downloadImage(urls[i], filename)
			if err != nil {
				fmt.Println(err)
				return
			}
			filename = "imagecon"
			wg.Done()
		}()
		wg.Wait()
	}
}

func main() {
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
	start := time.Now()
	downloadImagesSequential(urls)
	fmt.Printf("Sequential download took: %v\n", time.Since(start))

	// Concurrent download
	start = time.Now()
	downloadImagesConcurrent(urls)
	fmt.Printf("Concurrent download took: %v\n", time.Since(start))

}

// Helper function to download and save a single image.
func downloadImage(url, filename string) error {
	imgdownload.HTTPRequest(url, filename)
	return nil
}
