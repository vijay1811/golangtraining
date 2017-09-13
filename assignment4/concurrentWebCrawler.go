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

// Lang contains info about each website to be crawled
type Lang struct {
	Name     string
	URL      string
	Bytes    int
	TimeInMS int
}

// TODO : Apply all coments from assignment 3

// gets data from website and writes to files
func crawl(pfunc func([]byte, *Lang), lang *Lang) {

	data, bts, timeInMS, err := getDataLenTime(lang.URL)
	// when you get an error panic - Done
	if err != nil {
		panic(err)
	}
	lang.Bytes, lang.TimeInMS = bts, timeInMS/1000000
	// writes data into files
	pfunc(data, lang)
}

//TODO return error and handle it in crawl - Done
// returns data, length of data, time taken and error for get request
func getDataLenTime(url string) ([]byte, int, int, error) {

	startTime := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, 0, 0, err
	}

	timeInMS := time.Now().Sub(startTime)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, 0, 0, err
	}
	return data, len(data), int(timeInMS), nil
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
		// be ready for the panic here - Done
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main, Error : ", r)
			}
		}()

		go crawl(
			func(data []byte, lang *Lang) {
				// creates a file named goFormatted-<URL-Index>-<URL-Name>.html
				fptr1, err := os.Create("./goFormatted-" + strconv.Itoa(index+1) + "-" + strings.Split(strings.Split(websiteURL, "//")[1], "/")[0] + ".html")
				if err != nil {
					// panic - Done
					panic(err)
				}
				// creates a file named jsonFormatted-<URL-Index>-<URL-Name>.txt
				fptr2, err := os.Create("./jsonFormatted-" + strconv.Itoa(index+1) + "-" + strings.Split(strings.Split(websiteURL, "//")[1], "/")[0] + ".txt")
				if err != nil {
					// panic - Done
					panic(err)
				}
				bytesWritten, err := fptr1.Write(data)
				if err != nil {
					// panic - Done
					panic(err)
				}
				if bytesWritten != len(data) {
					fmt.Println("Some data was not written to file")
				}
				js, err := json.Marshal(*lang)
				if err != nil {
					// panic - Done
					panic(err)
				}
				bytesWritten, err = fptr2.Write(js)
				if err != nil {
					// panic - Done
					panic(err)
				}
				if bytesWritten != len(js) {
					fmt.Println("Some data was not written to file")
				}
				fptr1.Close()
				fptr2.Close()
			},
			&Lang{strings.Split(url, "//")[1], url, 0, 0.0})
		wg.Done()
	}
	wg.Wait()
}
