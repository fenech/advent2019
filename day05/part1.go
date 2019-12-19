package day05

import (
	"advent2019/day02"
	"log"
	"os"
	"strings"
)

func Part1(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	computer := day02.NewIntcode(file, strings.NewReader("1\n"), os.Stdout)
	computer.Run()
	computer.Out.Flush()
}
