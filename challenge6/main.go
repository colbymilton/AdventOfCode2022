package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const cMaxSize = 4

func main() {
	// load in file
	f, err := os.OpenFile("./input.txt", os.O_RDONLY, 777)
	if err != nil {
		log.Fatalln(err)
	}

	// scan through file
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)
	chars := make([]rune, cMaxSize)
	lastChar := 0
	for scanner.Scan() {
		char := rune(scanner.Text()[0])
		chars = chars[1:cMaxSize] // trim first character
		chars = append(chars, char)
		lastChar++

		if chars[0] != 0 {
			// check if 4 are unique
			unique := true
		L:
			for i, c := range chars {
				for j, c2 := range chars {
					if i == j {
						continue
					}
					if c == c2 {
						unique = false
						break L
					}
				}
			}

			if unique {
				break
			}
		}
	}

	fmt.Println("The first start-of-packet marker is detected at", lastChar)
}
