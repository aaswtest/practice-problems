package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

/*
 * Complete the 'superReducedString' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func superReducedString(s string) string {
	stack := []rune{}
	for _, v := range s {
		if len(stack) == 0 || len(stack) > 0 && stack[len(stack)-1] != v {
			stack = append(stack, v)
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return "Empty String"
	}
	return string(stack)
}

func TestReduction(t *testing.T) {
	le := logrus.NewEntry(logrus.New())
	le.Logger.SetLevel(logrus.FatalLevel)

	tests := []struct {
		testName string
		str      string
		expected string
	}{
		{testName: "success",
			str:      "aaabccddd",
			expected: "abd"},
		{testName: "success",
			str:      "aa",
			expected: ""},
		{testName: "success",
			str:      "baab",
			expected: ""},
	}
	for _, tt := range tests {
		got := superReducedString(tt.str)
		if got != tt.expected {
			t.Errorf("Reduce got string %s, want %s", got, tt.expected)
			return
		}

	}

}
