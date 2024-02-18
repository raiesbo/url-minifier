package utils

import (
	"testing"
)

// TestCreateURLKey test the lengths of the results of CreateURLKey
func TestCreateURLKey(t *testing.T) {
	tests := []int{10, 5, 6, 7}
	for _, length := range tests {
		key := CreateURLKey(length)
		if len(key) != length {
			t.Fatalf(`The length of key is %v and should have been %v`, len(key), length)
		}
	}
}
