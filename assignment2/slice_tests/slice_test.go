package main

import (
	"testing"
)

func Test_getFrequency(t *testing.T) {
	input := "123123"
	testedOutput := getFrequency(input)
	expectedOutput := make([]int, 10)
	expectedOutput[1] = 2
	expectedOutput[2] = 2
	expectedOutput[3] = 2
	for idx := range expectedOutput {
		if testedOutput[idx] != expectedOutput[idx] {
			t.Errorf("expectedOutput[%d]= %d, testedOutput[%d]= %d", idx, expectedOutput[idx], idx, testedOutput[idx])
		}
	}
}

func Test_isValid1(t *testing.T) {
	input := "123"
	testedValidation, testedMsg := isValid(input)
	expectedValidation, expectedMsg := true, ""
	if testedValidation != expectedValidation || testedMsg != expectedMsg {
		t.Errorf("testedValidation = %v testedMsg = %v expectedValidation = %v expectedMsg = %v  ", testedValidation, testedMsg, expectedValidation, expectedMsg)
	}
}

func Test_isValid2(t *testing.T) {
	input := "123344535546776768757653265"
	testedValidation, testedMsg := isValid(input)
	expectedValidation, expectedMsg := false, " Enter a numeral having length less than or equal to 15 "
	if testedValidation != expectedValidation || testedMsg != expectedMsg {
		t.Errorf("testedValidation = %v testedMsg = %v expectedValidation = %v expectedMsg = %v  ", testedValidation, testedMsg, expectedValidation, expectedMsg)
	}
}
func Test_isValid3(t *testing.T) {
	input := "13rrt"
	testedValidation, testedMsg := isValid(input)
	expectedValidation, expectedMsg := false, "Enter Valid Characters "
	if testedValidation != expectedValidation || testedMsg != expectedMsg {
		t.Errorf("testedValidation = %v testedMsg = %v expectedValidation = %v expectedMsg = %v  ", testedValidation, testedMsg, expectedValidation, expectedMsg)
	}
}
func Test_isValid4(t *testing.T) {
	input := "12387r8ru3uy88u8u8ue83u8u8u98d3b3c8nm"
	testedValidation, testedMsg := isValid(input)
	expectedValidation, expectedMsg := false, "Enter Valid Characters with a maximum length of 15 "
	if testedValidation != expectedValidation || testedMsg != expectedMsg {
		t.Errorf("testedValidation = %v testedMsg = %v expectedValidation = %v expectedMsg = %v  ", testedValidation, testedMsg, expectedValidation, expectedMsg)
	}
}

func Benchmark_containsInvalidInput(b *testing.B) {
	input := "12ascc$%23"
	for i := 0; i <= b.N; i++ {
		output, _ := isValid(input)
		if !output {
			b.Errorf("")
		}
	}
}
