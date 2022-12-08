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

	// check scenic score
	rows := len(trees)
	cols := len(trees[0])
	highest := 0
	for row, _ := range trees {
		for col, _ := range trees {
			height := trees[row][col]

			up := 0
			right := 0
			down := 0
			left := 0

			// up
			for j := row - 1; j >= 0; j-- {
				up++
				other := trees[j][col]
				if other >= height {
					break
				}
			}

			// right
			for j := col + 1; j < cols; j++ {
				right++
				other := trees[row][j]
				if other >= height {
					break
				}
			}

			// down
			for j := row + 1; j < rows; j++ {
				down++
				other := trees[j][col]
				if other >= height {
					break
				}
			}

			// left
			for j := col - 1; j >= 0; j-- {
				left++
				other := trees[row][j]
				if other >= height {
					break
				}
			}

			score := up * right * down * left
			fmt.Println(up, right, down, left, score)
			if score > highest {
				highest = score
			}
		}
	}

	fmt.Printf("There are highest scenic score is %v!\n", highest)
}
