package main

import (
	"bufio"
	"path/filepath"
)

func DecodeString(data string) string { return decode([]byte(data)) }

func DecodeFile() {
	path, _ := filepath.Abs("input.txt")
	readfile := readFile(path)
	defer readfile.Close()

	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		decode([]byte(fileScanner.Text()))
	}
}

func decode(data []byte) string {
	decodingTable := []byte{
		80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, /* 0 - 15 */
		80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, /* 16 - 31 */
		80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 62, 80, 80, 80, 63, /* 32 - 47 */
		52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 80, 80, 80, 64, 80, 80, /* 48 - 63 */
		80, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, /* 64 - 79 */
		15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 80, 80, 80, 80, 80, /* 80 - 96 */
		80, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, /* 87 - 111 */
		41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 80, 80, 80, 80, 80, /* 112 - 127 */
	}

	runes := make([]rune, 0)
	i := 0

	for ; i < len(data)-3; i += 4 {
		char := (decodingTable[data[i]] << 2) | (decodingTable[data[i+1]] >> 4)
		runes = append(runes, rune(char))

		if data[i+2] != '=' {
			char = (decodingTable[data[i+1]] << 4) | (decodingTable[data[i+2]] >> 2)
			runes = append(runes, rune(char))
		}

		if data[i+3] != '=' {
			char = (decodingTable[data[i+2]] << 6) | decodingTable[data[i+3]]
			runes = append(runes, rune(char))
		}
	}

	if i < len(data) {
		char := (decodingTable[data[i]] << 2) | (decodingTable[data[i+1]] >> 4)
		runes = append(runes, rune(char))

		if i+2 < len(data) && data[i+2] != '=' {
			char = (decodingTable[data[i+1]] << 4) | (decodingTable[data[i+2]] >> 2)
			runes = append(runes, rune(char))
		}

		if i+3 < len(data) && data[i+3] != '=' {
			char = (decodingTable[data[i+2]] << 6) | decodingTable[data[i+3]]
			runes = append(runes, rune(char))
		}
	}

	return string(runes)
}
