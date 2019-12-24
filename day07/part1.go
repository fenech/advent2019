package day07

import (
	"advent2019/day02"
	"bytes"
	"io"
	"log"
	"strconv"
	"strings"
)

func permutations(arr []string) [][]string {
	var helper func([]string, int)
	res := [][]string{}

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func Amplify(program io.ReadSeeker, perm []string) int {
	var pr *io.PipeReader
	var pw *io.PipeWriter
	var buf *bytes.Buffer

	for i, phase := range perm {
		var in io.Reader
		if i == 0 {
			in = strings.NewReader(phase + "\n0\n")
		} else {
			in = strings.NewReader(phase + "\n" + buf.String())
		}

		program.Seek(0, 0)
		pr, pw = io.Pipe()

		go func() {
			defer pw.Close()
			amp := day02.NewIntcode(program, in, pw)
			amp.Run()
			amp.Out.Flush()
		}()

		buf = new(bytes.Buffer)
		io.Copy(buf, pr)
	}

	val, _ := strconv.Atoi(strings.TrimRight(buf.String(), "\n"))
	return val
}

func Part1(program io.ReadSeeker) int {
	phases := []string{"0", "1", "2", "3", "4"}
	var maxVal int

	for _, perm := range permutations(phases) {
		val := Amplify(program, perm)
		if val > maxVal {
			maxVal = val
		}
	}

	log.Println(maxVal)
	return maxVal
}
