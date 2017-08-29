package main

import (
	"testing"
)

func Test_getFreqMap(t *testing.T) {
	input := "hi hi Hi HI"
	expectedOutput := map[string]int{"hi": 2, "Hi": 1, "HI": 1}
	testedOutput := getFreqMap(input)
	for key, val := range expectedOutput {
		if testedOutput[key] != val {
			t.Error("Expected Output k : v", key, val, "Tested Output  k : v", key, testedOutput[key])
		}
	}
}

func Test_isValid1(t *testing.T) {
	input1 := "I am fine1"
	testedOutput1 := isValid(input1)
	expectedOutput1 := false
	if testedOutput1 != expectedOutput1 {
		t.Error("Expected Output", expectedOutput1, "Tested Output", testedOutput1)
	}
}

func Test_isValid1failure(t *testing.T) {
	t.Fatalf("this test case is failing")
	input1 := "I am fine1"
	testedOutput1 := isValid(input1)
	expectedOutput1 := false
	if testedOutput1 != expectedOutput1 {
		t.Error("Expected Output", expectedOutput1, "Tested Output", testedOutput1)
	}
}

func Test_isValid2(t *testing.T) {
	input1 := "i am fine"
	testedOutput1 := isValid(input1)
	expectedOutput1 := true
	if testedOutput1 != expectedOutput1 {
		t.Error("Expected Output", expectedOutput1, "Tested Output", testedOutput1)
	}
}
