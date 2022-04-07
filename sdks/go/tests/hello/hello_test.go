package hello

import (
	"testing"
)

func TestHelloName(t *testing.T) {
	phrase := "Hello World"

	if phrase != "Hello World" {
		t.Fatalf("Test failed %v", phrase)
	}
}
