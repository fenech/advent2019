package day02

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"strconv"
)

type Intcode struct {
	pointer int
	State   []int
}

func makeIntcode(r io.Reader) *Intcode {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal("failed to read input:", err)
	}

	split := bytes.Split(bytes.TrimRight(data, "\n"), []byte(","))
	state := make([]int, 0, len(split))
	for _, c := range split {
		p, err := strconv.Atoi(string(c))
		if err != nil {
			log.Fatal(err, c)
		}
		state = append(state, p)
	}

	return &Intcode{State: state}
}

func (c *Intcode) Compute() (state []int) {
	for c.pointer < len(c.State) {
		switch c.State[c.pointer] {
		case 1:
			c.State[c.State[c.pointer+3]] = c.Add(c.pointer+1, c.pointer+2)
			c.pointer += 4
		case 2:
			c.State[c.State[c.pointer+3]] = c.Multiply(c.pointer+1, c.pointer+2)
			c.pointer += 4
		case 99:
			return c.State
		}
	}

	return c.State
}

func (c *Intcode) Add(o1, o2 int) int {
	return c.State[c.State[o1]] + c.State[c.State[o2]]
}

func (c *Intcode) Multiply(o1, o2 int) int {
	return c.State[c.State[o1]] * c.State[c.State[o2]]
}
