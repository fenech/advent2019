package day06

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type Orbit struct {
	Next  string
	Count int
}

func Recurse(orbits map[string]Orbit, planet string, count int) int {
	if planet == "COM" {
		return 0
	}

	c := orbits[orbits[planet].Next].Count
	if c == 0 {
		c = Recurse(orbits, orbits[planet].Next, count+1)
	}

	orbits[planet] = Orbit{
		Next:  orbits[planet].Next,
		Count: c + 1,
	}

	return c + 1
}

func Checksum(r io.Reader) (c int) {
	s := bufio.NewScanner(r)
	lines := make(map[string]Orbit)
	for s.Scan() {
		parts := strings.Split(s.Text(), ")")
		lines[parts[1]] = Orbit{Next: parts[0]}
	}

	for planet := range lines {
		c += Recurse(lines, planet, 0)
	}

	return
}

func Part1(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(Checksum(file))
}
