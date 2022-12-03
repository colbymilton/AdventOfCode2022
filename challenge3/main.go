package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// load in file
	f, err := os.OpenFile("./input.txt", os.O_RDONLY, 777)
	if err != nil {
		log.Fatalln(err)
	}

	// scan through file
	scanner := bufio.NewScanner(f)
	priority := 0
	for scanner.Scan() {
		line := scanner.Text()

		// split line in 2
		l := len(line)
		if l%2 != 0 {
			log.Fatalln("line is not of even length")
		}
		c1 := line[:l/2]
		c2 := line[l/2:]

		// find which character appears in both
		match := 0
		for i := 0; i < len(c1); i++ {
			w := c1[i]
			if strings.ContainsRune(c2, rune(w)) {
				match = int(w)
				break
			}
		}
		if match == 0 {
			log.Fatalln("did not find match")
		}

		// calculate priority
		if match >= 97 { // lowercase
			priority += match - 96
		} else {
			priority += match - 38
		}
	}

	fmt.Println("The total priority of the bags is", priority)
}
