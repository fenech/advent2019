package main

import (
	"advent2019/day02"
	"advent2019/day05"
	"advent2019/day06"
	"advent2019/day07"
	"io"
	"log"
	"os"
)

func run(path string, part func(io.ReadSeeker)) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	part(file)
}

func main() {
	log.Println("Day 2, part 1")
	run("day02/input.txt", func(rs io.ReadSeeker) { day02.Part1(rs) })

	log.Println("Day 5, part 1")
	run("day05/input.txt", func(rs io.ReadSeeker) { day05.Part1(rs) })
	log.Println("Day 5, part 2")
	run("day05/input.txt", func(rs io.ReadSeeker) { day05.Part2(rs) })

	log.Println("Day 6, part 1")
	run("day06/input.txt", func(rs io.ReadSeeker) { day06.Part1(rs) })
	log.Println("Day 6, part 2")
	run("day06/input.txt", func(rs io.ReadSeeker) { day06.Part2(rs) })

	log.Println("Day 7, part 1")
	run("day07/input.txt", func(rs io.ReadSeeker) { day07.Part1(rs) })
}
