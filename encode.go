package main

import (
	"bufio"
	"os"
	"path/filepath"
)

const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func EncodeString(data string) string { return encode([]byte(data)) }

func EncodeFile() {
	path, _ := filepath.Abs("input.txt")
	readfile := readFile(path)
	defer readfile.Close()

	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		encode([]byte(fileScanner.Text()))
	}
}

func encode(data []byte) string {
	runes := make([]rune, 0)

	i := 0
	for ; i < len(data)-2; i += 3 {
		char := data[i] >> 2
		runes = append(runes, rune(characters[char]))

		char = ((data[i] & 0x03) << 4) | ((data[i+1] & 0xF0) >> 4)
		runes = append(runes, rune(characters[char]))

		char = ((data[i+1] & 0x0f) << 2) | ((data[i+2] & 0xC0) >> 6)
		runes = append(runes, rune(characters[char]))

		char = data[i+2] & 0x3F
		runes = append(runes, rune(characters[char]))
	}

	if i < len(data) {
		nextToLast := data[i]

		if (i + 1) < len(data) {
			last := data[i+1]
			char := nextToLast >> 2
			runes = append(runes, rune(characters[char]))

			char = ((nextToLast & 0x03) << 4) | ((last & 0xF0) >> 4)
			runes = append(runes, rune(characters[char]))

			char = (last & 0x0F) << 2
			runes = append(runes, rune(characters[char]))
		} else {
			char := nextToLast >> 2
			runes = append(runes, rune(characters[char]))

			char = (nextToLast & 0x03) << 4
			runes = append(runes, rune(characters[char]))
			runes = append(runes, '=')
		}
		runes = append(runes, '=')
	}

	return string(runes)
}

// Utility functions.

func readFile(path string) *os.File {
	readfile, err := os.Open(path)
	if err != nil {
		os.Exit(1)
	}
	return readfile
}
