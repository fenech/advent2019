package day06

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func Path(orbits map[string]Orbit, planet string, path []string) []string {
	path = append(path, planet)
	if planet == "COM" {
		return path
	}
	return Path(orbits, orbits[planet].Next, path)
}

func Distance(r io.Reader) int {
	s := bufio.NewScanner(r)
	lines := make(map[string]Orbit)
	for s.Scan() {
		parts := strings.Split(s.Text(), ")")
		lines[parts[1]] = Orbit{Next: parts[0]}
	}

	you := Path(lines, "YOU", []string{})
	san := Path(lines, "SAN", []string{})

	var i int
	for i = 0; you[len(you)-i-1] == san[len(san)-i-1]; i++ {
	}

	return len(you) + len(san) - 2*(1+i)
}

func Part2(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(Distance(file))
}
