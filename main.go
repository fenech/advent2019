package main

import (
	"advent2019/day02"
	"advent2019/day05"
	"advent2019/day06"
	"log"
)

func main() {
	log.Println("Day 2, part 1")
	day02.Part1("day02/input.txt")

	log.Println("Day 5, part 1")
	day05.Part1("day05/input.txt")
	log.Println("Day 5, part 2")
	day05.Part2("day05/input.txt")

	log.Println("Day 6, part 1")
	day06.Part1("day06/input.txt")
	log.Println("Day 6, part 2")
	day06.Part2("day06/input.txt")
}
