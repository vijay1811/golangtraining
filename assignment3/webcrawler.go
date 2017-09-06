package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Lang struct {
	Name     string
	URL      string
	Bytes    int
	TimeInMS int
}

func crawl(pfunc func([]byte, *Lang), lang *Lang) {
	data, bts, timeInMS := getDataLenTime(lang.URL)
	lang.Bytes, lang.TimeInMS = bts, timeInMS/1000000
	pfunc(data, lang)
}

func getDataLenTime(url string) ([]byte, int, int) {

	startTime := time.Now()
	resp, err := http.Get(url)
	timeInMS := time.Now().Sub(startTime)
	if err != nil {
		fmt.Println(err)
		return nil, 0, 0.0
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, 0, 0.0
	}
	return data, len(data), int(timeInMS)
}

func main() {
	urlArr := []string{"https://www.python.org/",
		"https://www.ruby-lang.org/en/",
		"https://golang.org/"}

	for idx, url := range urlArr {
		crawl(
			func(data []byte, lang *Lang) {
				fptr1, err := os.Create("./goFormatted-" + strconv.Itoa(idx+1) + "-" + strings.Split(strings.Split(url, "//")[1], "/")[0] + ".html")
				fptr2, err := os.Create("./jsonFormatted-" + strconv.Itoa(idx+1) + "-" + strings.Split(strings.Split(url, "//")[1], "/")[0] + ".txt")
				defer fptr1.Close()
				defer fptr2.Close()

				if err != nil {
					fmt.Println(err)
					return
				}
				_, err = fptr1.Write(data)

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
}
