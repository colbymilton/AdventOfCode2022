package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	cCMD string = "$"
	cCD  string = "cd"
	cLS  string = "ls"
)

const (
	cRootDir string = "\\"
	cUpDir   string = ".."
	cDir     string = "dir"
)

type Dir struct {
	name   string
	files  map[string]int  // filename to size
	dirs   map[string]*Dir // dirname to dir
	parent *Dir
}

func NewDir(name string) *Dir {
	return &Dir{
		name:  name,
		files: make(map[string]int),
		dirs:  make(map[string]*Dir),
	}
}

func (d *Dir) AddFile(filename string, filesize int) {
	d.files[filename] = filesize
}

func (d *Dir) AddDir(dir *Dir) {
	if _, ok := d.dirs[dir.name]; ok {
		return
	}

	dir.parent = d
	d.dirs[dir.name] = dir
}

func (d *Dir) Size() int {
	sum := 0
	for _, s := range d.files {
		sum += s
	}
	for _, dir := range d.dirs {
		sum += dir.Size()
	}
	return sum
}

func main() {
	// load in file
	f, err := os.OpenFile("./input.txt", os.O_RDONLY, 777)
	if err != nil {
		log.Fatalln(err)
	}

	// scan through file and build directory structure
	scanner := bufio.NewScanner(f)
	d := NewDir(cRootDir)
	root := d
	lastCMD := ""
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		if parts[0] == cCMD { // line starts with $
			if len(parts) <= 1 {
				log.Fatalln("invalid line:", line)
			}

			lastCMD = parts[1]

			if parts[1] == cCD { // is cd line
				if len(parts) <= 2 {
					log.Fatalln("invalid line:", line)
				}

				newdir := parts[2]

				if newdir == cRootDir {
					continue
				}

				if newdir == cUpDir {
					if d.parent == nil {
						log.Fatalln("cannot go above root dir")
					}
					d = d.parent
				} else {                             // going down
					if _, ok := d.dirs[newdir]; ok { // dir exists
						d = d.dirs[newdir]
					} else { // dir does not exist
						d.AddDir(NewDir(newdir))
					}
				}
			} else if parts[1] == cLS { // is ls line
				continue
			}
		} else { // line does not start with $ and is likely the result of "ls"
			if lastCMD != cLS {
				log.Fatalln("list of files not after ls")
			}

			if len(parts) <= 1 {
				log.Fatalln("invalid line:", line)
			}

			if parts[0] == cDir {
				newdir := parts[1]
				d.AddDir(NewDir(newdir))
			} else {
				s, err := strconv.Atoi(parts[0])
				if err != nil {
					log.Fatalln(err)
				}

				d.AddFile(parts[1], s)
			}
		}
	}

	fmt.Println("Total size:", root.Size())
	fmt.Println("Total size of all directories with at most 100000:", totalSizeOfAtMost(root, 100000))
}

func totalSizeOfAtMost(d *Dir, maxSize int) int {
	sum := 0

	// escape
	dirSize := d.Size()
	if dirSize <= maxSize {
		sum += dirSize
	}

	// continue
	for _, sd := range d.dirs {
		sum += totalSizeOfAtMost(sd, maxSize)
	}

	return sum
}
