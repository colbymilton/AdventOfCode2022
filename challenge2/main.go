package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	cRockScore     int = 1
	cPaperScore    int = 2
	cScissorsScore int = 3
	cLossScore     int = 0
	cDrawScore     int = 3
	cWinScore      int = 6
)

const (
	cOppRock     string = "A"
	cOppPaper    string = "B"
	cOppScissors string = "C"
	cYouRock     string = "X"
	cYouPaper    string = "Y"
	cYouScissors string = "Z"
	cYouLose     string = "X"
	cYouDraw     string = "Y"
	cYouWin      string = "Z"
)

func main() {
	// load in file
	f, err := os.OpenFile("./input.txt", os.O_RDONLY, 777)
	if err != nil {
		log.Fatalln(err)
	}

	// scan through file and calculate score
	totalScore := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			log.Fatalln("could not parse line:", line)
		}

		op := parts[0]
		out := parts[1]

		// determine your choice based on needed outcome
		you := ""
		switch out {
		case cYouLose:
			switch op {
			case cOppRock:
				you = cYouScissors
			case cOppPaper:
				you = cYouRock
			case cOppScissors:
				you = cYouPaper
			}
		case cYouDraw:
			switch op {
			case cOppRock:
				you = cYouRock
			case cOppPaper:
				you = cYouPaper
			case cOppScissors:
				you = cYouScissors
			}
		case cYouWin:
			switch op {
			case cOppRock:
				you = cYouPaper
			case cOppPaper:
				you = cYouScissors
			case cOppScissors:
				you = cYouRock
			}
		}

		// score for your selection
		switch you {
		case cYouRock:
			totalScore += cRockScore
		case cYouPaper:
			totalScore += cPaperScore
		case cYouScissors:
			totalScore += cScissorsScore
		}

		// score for outcome
		if op == cOppRock {
			switch you {
			case cYouRock:
				totalScore += cDrawScore
			case cYouPaper:
				totalScore += cWinScore
			case cYouScissors:
				totalScore += cLossScore
			}
		} else if op == cOppPaper {
			switch you {
			case cYouRock:
				totalScore += cLossScore
			case cYouPaper:
				totalScore += cDrawScore
			case cYouScissors:
				totalScore += cWinScore
			}
		} else if op == cOppScissors {
			switch you {
			case cYouRock:
				totalScore += cWinScore
			case cYouPaper:
				totalScore += cLossScore
			case cYouScissors:
				totalScore += cDrawScore
			}
		}
	}

	fmt.Println("By following the strategy guide, you will get a score of:", totalScore)
}
