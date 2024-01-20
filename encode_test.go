package main

import (
	"testing"
)

func TestEncodeString(t *testing.T) {
	encoding := EncodeString("Hello, World!")
	if encoding != "SGVsbG8sIFdvcmxkIQ==" {
		t.Fatalf("Wrong encoding")
	}

	encoding = EncodeString("S")
	if encoding != "Uw==" {
		t.Fatalf("Wrong encoding")
	}

	encoding = EncodeString("Su")
	if encoding != "U3U=" {
		t.Fatalf("Wrong encoding")
	}

}
