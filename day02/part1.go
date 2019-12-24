package day02

import (
	"fmt"
	"io"
)

func Part1(r io.Reader) {
	computer := NewIntcode(r, nil, nil)
	computer.State[1] = 12
	computer.State[2] = 2
	fmt.Println(computer.Run())
}
