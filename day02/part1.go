package day02

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Part1() {
	path, err := filepath.Abs("day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	computer := makeIntcode(file)
	computer.State[1] = 12
	computer.State[2] = 2
	fmt.Println(computer.Compute())
}
