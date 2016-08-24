package main

import (
	"testing"
)

// Fictious testing for Travis as it would require database
func TestCreateUrl(t *testing.T) {
	short := shorten(16)
	if len(short) != 16 {
		t.Error("Shortened URL is not as short as expected!")
	}
}
