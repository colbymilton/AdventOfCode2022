package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	Id       int
	Calories int
}

type Elves []*Elf

func (e Elves) Len() int           { return len(e) }
func (e Elves) Less(i, j int) bool { return e[i].Calories > e[j].Calories }
func (e Elves) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

func main() {
	// load in file
	f, err := os.OpenFile("./input.txt", os.O_RDONLY, 777)
	if err != nil {
		log.Fatalln(err)
	}

	// scan through file and count calories
	scanner := bufio.NewScanner(f)
	elves := make([]*Elf, 0)
	count := 1
	currentElf := &Elf{Id: 1, Calories: 0}
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			// next elf
			elves = append(elves, currentElf)
			count++
			currentElf = &Elf{Id: count, Calories: 0}
			continue
		}

		lineVal, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalln(err)
		}

		currentElf.Calories += lineVal
	}

	// handle last elf
	elves = append(elves, currentElf)

	sort.Sort(Elves(elves))

	// first in list should have the most calories
	fmt.Printf("The elf with the most calories is Elf #%v with %v calories!\n", elves[0].Id, elves[0].Calories)
}
