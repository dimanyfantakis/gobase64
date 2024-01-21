package main

import (
	"testing"
)

func TestDecodeString(t *testing.T) {
	decoding := DecodeString("Uw==")
	if decoding != "S" {
		t.Errorf("Wrong decoding")
	}

	decoding = DecodeString("U3U=")
	if decoding != "Su" {
		t.Errorf("Wrong decoding")
	}

	decoding = DecodeString("U3Vu")
	if decoding != "Sun" {
		t.Errorf("Wrong decoding")
	}

}

func FuzzDecodeString(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev := EncodeString(orig)
		doubleRev := DecodeString(rev)
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
	})
}
