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

var (
	urlfunc = http.Get
	timeNow = time.Now
	readAll = ioutil.ReadAll
)

type Lang struct {
	Name     string
	URL      string
	Bytes    int
	TimeInMS int
}

func crawl(pfunc func([]byte, *Lang), lang *Lang) {
	data, bts, timeInMS, err := getDataLenTime(lang.URL)
	// when you get an error panic - Done
	if err != nil {
		panic(err)
	}
	lang.Bytes, lang.TimeInMS = bts, timeInMS/1000000
	pfunc(data, lang)
}

//TODO return error and handle it in crawl - Done
func getDataLenTime(url string) ([]byte, int, int, error) {

	startTime := time.Now()
	resp, err := urlfunc(url)
	if err != nil {
		fmt.Println(err)
		return nil, 0, 0, err
	}

	data, err := readAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, 0, 0, err
	}
	timeInMS := time.Now().Sub(startTime)
	return data, len(data), int(timeInMS), nil
}

func main() {
	urlArr := []string{"https://www.python.org/",
		"https://www.ruby-lang.org/en/",
		"https://golang.org/"}

	for idx, url := range urlArr {
		// be ready for the panic here - Done
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main, Error : ", r)
			}
		}()
		crawl(
			func(data []byte, lang *Lang) {
				fptr1, err := os.Create("./goFormatted-" + strconv.Itoa(idx+1) + "-" + strings.Split(strings.Split(url, "//")[1], "/")[0] + ".html")
				if err != nil {
					fmt.Println(err)
					return
				}
				// handle error immediately - Done
				fptr2, err := os.Create("./jsonFormatted-" + strconv.Itoa(idx+1) + "-" + strings.Split(strings.Split(url, "//")[1], "/")[0] + ".txt")
				if err != nil {
					fmt.Println(err)
					return
				}
				bytesWritten, err := fptr1.Write(data)
				if err != nil {
					// panic - Done
					panic(err)
				}
				if bytesWritten != len(data) {
					fmt.Println("Some data was not written to file")
				}

				// check if bytes written are equal to len(data) - Done
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
				// chek if bytes written are equal to len(data) - Done
				if bytesWritten != len(js) {
					fmt.Println("some data was not written into file")
				}
				fptr1.Close()
				fptr2.Close()
			},
			&Lang{strings.Split(url, "//")[1], url, 0, 0.0})
	}
}
