package utils

import (
	"os"
	"testing"
)

func TestDemo(t *testing.T) {
	tests := [][]string{
		{"cmd", "1", "value1"},
		{"cmd", "2", "value2"},
		{"cmd", "abc", "value3"}, // Invalid key test case
		{"cmd", "Z"},             // Insufficient arguments test case
	}

	for _, args := range tests {
		t.Run("Args_"+args[1], func(t *testing.T) {
			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()

			os.Args = args
			Demo()
		})
	}
}
