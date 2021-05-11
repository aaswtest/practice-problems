package main

import (
	"testing"
	"unicode"

	"github.com/sirupsen/logrus"
)

/*
 * Complete the 'camelcase' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func camelcase(s string) int32 {
	// Write your code here

	wC := 1

	for _, char := range s {
		if unicode.IsUpper(char) {
			wC++
		}
	}
	return int32(wC)
}

func TestCamel(t *testing.T) {
	le := logrus.NewEntry(logrus.New())
	le.Logger.SetLevel(logrus.FatalLevel)

	tests := []struct {
		testName string
		str      string
		expected int32
	}{
		{testName: "success",
			str:      "saveChangesInTheEditor",
			expected: 5},
		{testName: "Changes",
			str:      "oneTwoThree",
			expected: 3},
	}
	for _, tt := range tests {
		got := camelcase(tt.str)
		if got != tt.expected {
			t.Errorf("Reduce got string %v, want %v", got, tt.expected)
			return
		}

	}

}
