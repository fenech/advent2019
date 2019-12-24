package day05

import (
	"advent2019/day02"
	"io"
	"os"
	"strings"
)

func Part1(r io.Reader) {
	computer := day02.NewIntcode(r, strings.NewReader("1\n"), os.Stdout)
	computer.Run()
	computer.Out.Flush()
}
