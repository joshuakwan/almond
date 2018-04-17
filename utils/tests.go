package utils

import (
	"fmt"
	"path"
	"runtime"
	"testing"
)

func GetAlertmanagerConfigFilePath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	return path.Dir(filename) + "/resources/simple.yml"
}

// AssertEqual asserts if two objects are equal
func AssertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

// AssertNotEqual asserts if two objects are not equal
func AssertNotEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a != b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}
