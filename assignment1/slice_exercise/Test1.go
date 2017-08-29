package main

import (
	"fmt"
	"regexp"
)

var (
	re1, _ = regexp.Compile("^[0-9]{1,15}$")
	re2, _ = regexp.Compile("^[0-9]*$")
	re3, _ = regexp.Compile("^{1,15}$")
)

// func containsInvalidInput(str string) (bool, string) {
// 	if validString := re1.MatchString(str); validString {
// 		return false, ""
// 	} else if validChars := re2.MatchString(str); validChars {
// 		return true, " Enter a numeral having length less than or equal to 15 "
// 	} else if invalid := re3.MatchString(str); invalid {
// 		return true, "Enter Valid Characters "
// 	} else {
// 		return true, "Enter Valid Characters with a maximum length of 15 "
// 	}
// }

func containsInvalidInput(str string) (bool, string) {
	if validString, _ := regexp.MatchString("^[0-9]{1,15}$", str); validString {
		return false, ""
	} else if valid, _ := regexp.MatchString("^[0-9]*$", str); valid {
		return true, " Enter a numeral having length less than or equal to 15 "
	} else if valid, _ := regexp.MatchString("^.{1,15}$", str); valid {
		return true, "Enter Valid Characters "
	} else {
		return true, "Enter Valid Characters with a maximum length of 15 "
	}
}

func getFrequency(str string) []int {
	retval := make([]int, 10)
	// suggested : change trivial for loop to for range loop, wherever possible
	for _, v := range str {
		retval[int(v)-48]++
	}

	// for i := 0; i < len(str); i++ {
	// 	retval[int(str[i])-48]++
	// }
	return retval
}

func main() {

	var str string
	fmt.Scanln(&str)

	if condition, msg := containsInvalidInput(str); condition {
		fmt.Println(msg)
	} else {
		s := getFrequency(str)
		// suggested : change trivial for loop to for range loop, wherever possible
		for _, val := range s {
			fmt.Print(val, " ")
		}
	}
}
