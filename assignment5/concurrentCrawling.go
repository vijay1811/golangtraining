package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

// gets data from website and writes to files
func crawl(pfunc func([]byte, *Lang), lang *Lang) {

	data, bts, timeInMS := getDataLenTime(lang.URL)
	lang.Bytes, lang.TimeInMS = bts, timeInMS/1000000
	// prints length of data and crawl time to console
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

	totalCrawlTime := 0
	totalDataLength := 0
	crawlTime := make(chan int)
	dataLength := make(chan int)
	startTime := time.Now()
	for _, url := range urlArr {
		wg.Add(1)
		go crawl(
			func(data []byte, lang *Lang) {
				defer wg.Done()
				// putting time and data length values into respective  channels
				crawlTime <- lang.TimeInMS
				dataLength <- lang.Bytes
			},
			&Lang{strings.Split(url, "//")[1], url, 0, 0.0})
	}
	for range urlArr {
		// adding crawlTime, data length to their respective cumulatives
		totalCrawlTime += <-crawlTime
		totalDataLength += <-dataLength
	}
	wg.Wait()

	finalTime := time.Now()
	concurrentTimeTaken := int(finalTime.Sub(startTime)) / 1000000

	fmt.Println("Concurrent Crawling Time : ", concurrentTimeTaken)
	fmt.Println("Total Crawl Time : ", totalCrawlTime)
	fmt.Println("Total Data Length : ", totalDataLength)
}
