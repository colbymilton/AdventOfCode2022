package main

import (
	"bufio"
	"bytes"
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
	mode := 0 // 0 - stacking mode, 1 - moving mode
	var stacks []*Stack
	for scanner.Scan() {
		line := scanner.Text()

		if mode == 0 { // if in stacking mode
			if line == "" { // switch to moving mode
				mode = 1
				continue
			}

			if !strings.Contains(line, "[") { // ignore the numbered line
				continue
			}

			parts := splitStringN(line, 4)
			if stacks == nil {
				stacks = make([]*Stack, 0)
			}
			if len(parts) > len(stacks) {
				for i := len(stacks) + 1; i <= len(parts); i++ {
					stacks = append(stacks, NewStack())
				}
			}

			// add runes to stacks
			for i, p := range parts {
				r := rune(p[1])
				if r != ' ' {
					stacks[i].PutBottom(r)
				}
			}

		} else { // if in moving mode
			parts := strings.Split(line, " ")

			// parse
			amount, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatalln(err)
			}
			from, err := strconv.Atoi(parts[3])
			if err != nil {
				log.Fatalln(err)
			}
			to, err := strconv.Atoi(parts[5])
			if err != nil {
				log.Fatalln(err)
			}

			if from <= 0 || from > len(stacks) || to <= 0 || to > len(stacks) {
				log.Fatalln("unrecognized stack")
			}

			// move
			f := stacks[from-1]
			t := stacks[to-1]
			if amount > f.Size() {
				amount = f.Size()
			}
			rs := f.PopMany(amount)
			t.PutMany(rs)

			//for i := 0; i < amount; i++ {
			//	if f.Size() > 0 {
			//		r := f.Pop()
			//		t.Put(r)
			//	} else {
			//		break
			//	}
			//}
		}
	}

	fmt.Print("The top of each stack is: ")
	for _, s := range stacks {
		fmt.Print(string(s.Pop()))
	}
	fmt.Println()
}

func splitStringN(s string, n int) []string {
	sub := ""
	subs := make([]string, 0)

	runes := bytes.Runes([]byte(s))
	l := len(runes)
	for i, r := range runes {
		sub = sub + string(r)
		if (i+1)%n == 0 {
			subs = append(subs, sub)
			sub = ""
		} else if (i + 1) == l {
			subs = append(subs, sub)
		}
	}

	return subs
}

type Stack struct {
	vals []rune
}

func NewStack() *Stack {
	return &Stack{vals: make([]rune, 0)}
}

func (s *Stack) Put(r rune) {
	s.vals = append(s.vals, r)
}

func (s *Stack) PutMany(rs []rune) {
	s.vals = append(s.vals, rs...)
}

func (s *Stack) PutBottom(r rune) {
	s.vals = append([]rune{r}, s.vals...)
}

func (s *Stack) Pop() rune {
	r := s.vals[s.Size()-1]
	s.vals = s.vals[:s.Size()-1]
	return r
}

func (s *Stack) PopMany(c int) []rune {
	runes := s.vals[s.Size()-c:]
	s.vals = s.vals[:s.Size()-c]
	return runes
}

func (s *Stack) Size() int {
	return len(s.vals)
}

func (s *Stack) Print() {
	for _, r := range s.vals {
		fmt.Print(string(r))
	}
	fmt.Println()
}
