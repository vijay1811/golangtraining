package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// contains info about each website to be crawled
type Lang struct {
	Name     string
	URL      string
	Bytes    int
	TimeInMS int
}

// gets data from website and writes to files
func crawl(pfunc func([]byte, *Lang), lang *Lang) {

	data, bts, timeInMS := getDataLenTime(lang.URL)
	lang.Bytes, lang.TimeInMS = bts, timeInMS/1000000
	// writes data into files
	pfunc(data, lang)
}

// returns data, length of data, time taken for get request in milliseconds
func getDataLenTime(url string) ([]byte, int, int) {

	startTime := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, 0, 0.0
	}

	timeInMS := time.Now().Sub(startTime)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, 0, 0.0
	}
	return data, len(data), int(timeInMS)
}

func main() {
	var wg sync.WaitGroup
	urlArr := []string{"https://www.python.org/",
		"https://golang.org/",
		"https://www.ruby-lang.org/en/",
	}
	for idx, url := range urlArr {
		index := idx
		websiteURL := url
		wg.Add(1)
		go crawl(
			func(data []byte, lang *Lang) {
				defer wg.Done()
				// creates a file named goFormatted-<URL-Index>-<URL-Name>.html
				fptr1, err := os.Create("./goFormatted-" + strconv.Itoa(index+1) + "-" + strings.Split(strings.Split(websiteURL, "//")[1], "/")[0] + ".html")
				if err != nil {
					fmt.Println(err)
					return
				}
				defer fptr1.Close()
				// creates a file named jsonFormatted-<URL-Index>-<URL-Name>.txt
				fptr2, err := os.Create("./jsonFormatted-" + strconv.Itoa(index+1) + "-" + strings.Split(strings.Split(websiteURL, "//")[1], "/")[0] + ".txt")
				if err != nil {
					fmt.Println(err)
					return
				}
				defer fptr2.Close()
				_, err = fptr1.Write(data)
				if err != nil {
					fmt.Println(err)
					return
				}
				js, err := json.Marshal(*lang)
				if err != nil {
					fmt.Println(err)
					return
				}
				_, err = fptr2.Write(js)
				if err != nil {
					fmt.Println(err)
					return
				}
			},
			&Lang{strings.Split(url, "//")[1], url, 0, 0.0})
	}
	wg.Wait()
}
