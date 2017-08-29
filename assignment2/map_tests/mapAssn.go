package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func getFreqMap(str string) map[string]int {
	slice := strings.Fields(str)
	mp := make(map[string]int)
	for _, v := range slice {
		mp[v]++
	}
	return mp
}
func isValid(str string) bool {
	if valid, _ := regexp.MatchString("^[a-z ]*$", str); valid {
		return true
	}
	return false
}

func main() {
	var str string
	scr := bufio.NewScanner(os.Stdin)
	scr.Scan()
	str = scr.Text()

	str = strings.ToLower(str)
	if isValid(str) {
		freqMap := getFreqMap(str)
		for k, v := range freqMap {
			fmt.Println(k, " : ", v)
		}
	} else {
		fmt.Println("Enter a Valid String")
	}
}
