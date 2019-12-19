package day02

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"strconv"
)

type Intcode struct {
	Ptr   int
	State []int
}

func NewIntcode(r io.Reader) *Intcode {
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

func (c *Intcode) Run() (state []int) {
	var stop bool
	for !stop && c.Ptr < len(c.State) {
		stop = c.Compute()
	}

	return c.State
}

func (c *Intcode) Compute() (stop bool) {
	switch c.State[c.Ptr] {
	case 1:
		c.State[c.State[c.Ptr+3]] = c.Add(c.Ptr+1, c.Ptr+2)
		c.Ptr += 4
	case 2:
		c.State[c.State[c.Ptr+3]] = c.Multiply(c.Ptr+1, c.Ptr+2)
		c.Ptr += 4
	case 99:
		stop = true
	}
	return
}

func (c *Intcode) Add(o1, o2 int) int {
	return c.State[c.State[o1]] + c.State[c.State[o2]]
}

func (c *Intcode) Multiply(o1, o2 int) int {
	return c.State[c.State[o1]] * c.State[c.State[o2]]
}
