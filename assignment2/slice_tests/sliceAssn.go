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

func isValid(str string) (bool, string) {
	if validString, _ := regexp.MatchString("^[0-9]{1,15}$", str); validString {
		return true, ""
	} else if invalid, _ := regexp.MatchString("^[0-9]*$", str); invalid {
		return false, " Enter a numeral having length less than or equal to 15 "
	} else if invalid, _ := regexp.MatchString("^.{1,15}$", str); invalid {
		return false, "Enter Valid Characters "
	} else {
		return false, "Enter Valid Characters with a maximum length of 15 "
	}
}

func getFrequency(str string) []int {
	retval := make([]int, 10)
	for _, v := range str {
		retval[int(v)-48]++
	}
	return retval
}

func main() {

	var input string
	fmt.Scanln(&input)

	if valid, msg := isValid(input); !valid {
		fmt.Println(msg)
	} else {
		s := getFrequency(input)
		// suggested : change trivial for loop to for range loop, wherever possible
		for _, val := range s {
			fmt.Print(val, " ")
		}
	}
}
