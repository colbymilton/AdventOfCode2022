package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	contained := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			log.Fatalln("could not parse input line into 2 parts")
		}

		// split each part
		xs := strings.Split(parts[0], "-")
		ys := strings.Split(parts[1], "-")
		if len(xs) != 2 || len(ys) != 2 {
			log.Fatalln("could not parse line part into 2 parts")
		}

		// convert each part
		x0, err := strconv.Atoi(xs[0])
		if err != nil {
			log.Fatalln("could not parse as int")
		}
		x1, err := strconv.Atoi(xs[1])
		if err != nil {
			log.Fatalln("could not parse as int")
		}
		y0, err := strconv.Atoi(ys[0])
		if err != nil {
			log.Fatalln("could not parse as int")
		}
		y1, err := strconv.Atoi(ys[1])
		if err != nil {
			log.Fatalln("could not parse as int")
		}

		// check if one contains the other
		if (x0 <= y0 && x1 >= y1) || (y0 <= x0 && y1 >= x1) {
			contained++
		}
	}

	fmt.Printf("There are %v assignements that contain eachother.\n", contained)
}
