package day02

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Intcode struct {
	Ptr   int
	State []int
	In    *bufio.Reader
	Out   *bufio.Writer
}

func NewIntcode(r io.Reader, i io.Reader, o io.Writer) *Intcode {
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

	return &Intcode{
		State: state,
		In:    bufio.NewReader(i),
		Out:   bufio.NewWriter(o),
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
	instruction := c.State[c.Ptr]
	op := instruction % 100
	modes := int(instruction / 100)

	switch op {
	case 1:
		c.Add(modes)
	case 2:
		c.Multiply(modes)
	case 3:
		c.Input()
	case 4:
		c.Output(modes)
	case 99:
		stop = true
	}
	return
}

func (c *Intcode) Source(modes int, ptr int) (s int) {
	s = c.State[ptr]
	if modes%10 == 0 {
		s = c.State[s]
	}
	return
}

func (c *Intcode) Add(modes int) {
	s1 := c.Source(modes, c.Ptr+1)
	s2 := c.Source(modes/10, c.Ptr+2)
	c.State[c.State[c.Ptr+3]] = s1 + s2
	c.Ptr += 4
}

func (c *Intcode) Multiply(modes int) {
	s1 := c.Source(modes, c.Ptr+1)
	s2 := c.Source(modes/10, c.Ptr+2)
	c.State[c.State[c.Ptr+3]] = s1 * s2
	c.Ptr += 4
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

func (c *Intcode) Output(modes int) {
	s1 := c.Source(modes, c.Ptr+1)
	c.Out.WriteString(strconv.Itoa(s1))
	c.Ptr += 2
}
