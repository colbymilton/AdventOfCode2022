package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// load in file
	f, err := os.OpenFile("./input.txt", os.O_RDONLY, 777)
	if err != nil {
		log.Fatalln(err)
	}

	// scan through file and build tree array
	scanner := bufio.NewScanner(f)
	trees := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()

		row := make([]int, len(line))

		for i, c := range line {
			x, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatalln(err)
			}

			row[i] = x
		}

		trees = append(trees, row)
	}

	// check if tree is visible from any direction
	rows := len(trees)
	cols := len(trees[0])
	visCount := 0
	for row, _ := range trees {
		for col, _ := range trees {
			height := trees[row][col]

			// determine if visible
			visible := true

			// up
			for j := row - 1; j >= 0; j-- {
				other := trees[j][col]
				if other >= height {
					visible = false
					break
				}
			}
			if visible {
				visCount++
				continue
			}

			// right
			visible = true
			for j := col + 1; j < cols; j++ {
				other := trees[row][j]
				if other >= height {
					visible = false
					break
				}
			}
			if visible {
				visCount++
				continue
			}

			// down
			visible = true
			for j := row + 1; j < rows; j++ {
				other := trees[j][col]
				if other >= height {
					visible = false
					break
				}
			}
			if visible {
				visCount++
				continue
			}

			// left
			visible = true
			for j := col - 1; j >= 0; j-- {
				other := trees[row][j]
				if other >= height {
					visible = false
					break
				}
			}
			if visible {
				visCount++
				continue
			}
		}
	}

	fmt.Printf("There are %v visible trees!\n", visCount)
}
