package day02

import (
	"fmt"
	"log"
	"os"
)

func Part1(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	computer := NewIntcode(file, nil, nil)
	computer.State[1] = 12
	computer.State[2] = 2
	fmt.Println(computer.Run())
}
