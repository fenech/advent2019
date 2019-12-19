package day05

import (
	"advent2019/day02"
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

type Intcode struct {
	*day02.Intcode
	In  *bufio.Reader
	Out *bufio.Writer
}

func NewIntcode(r io.Reader, i io.Reader, o io.Writer) *Intcode {
	return &Intcode{
		day02.NewIntcode(r),
		bufio.NewReader(i),
		bufio.NewWriter(o),
	}
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
	case 3:
		c.Input()
	case 4:
		c.Output()
	default:
		stop = c.Intcode.Compute()
	}
	return
}

func (c *Intcode) Input() {
	s, err := c.In.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	d, err := strconv.Atoi(strings.TrimRight(s, "\n"))
	c.State[c.State[c.Ptr+1]] = d
	c.Ptr += 2
}

func (c *Intcode) Output() {
	c.Out.WriteString(strconv.Itoa(c.State[c.State[c.Ptr+1]]))
	c.Ptr += 2
}
