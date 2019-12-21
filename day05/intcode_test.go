package day05

import (
	"advent2019/day02"
	"bufio"
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestIntcode_Run(t *testing.T) {
	type fields struct {
		state  []int
		input  string
		output *bytes.Buffer
	}
	tests := []struct {
		name       string
		fields     fields
		wantState  []int
		wantOutput string
	}{
		{
			"3,0,4,0,99 outputs whatever it gets as input, then halts.",
			fields{state: []int{3, 0, 4, 0, 99}, input: "123\n", output: &bytes.Buffer{}},
			[]int{123, 0, 4, 0, 99},
			"123\n",
		},
		{
			"Integers can be negative: 1101,100,-1,4,0 is a valid program (find 100 + -1, store the result in position 4).",
			fields{state: []int{1101, 100, -1, 4, 0}, input: "", output: &bytes.Buffer{}},
			[]int{1101, 100, -1, 4, 99},
			"",
		},
		{
			"3,9,8,9,10,9,4,9,99,-1,8 outputs 1 if the input is equal to 8",
			fields{state: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, input: "8\n", output: &bytes.Buffer{}},
			[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, 1, 8},
			"1\n",
		},
		{
			"3,9,8,9,10,9,4,9,99,-1,8 outputs 1 if the input is not equal to 8",
			fields{state: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, input: "7\n", output: &bytes.Buffer{}},
			[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, 0, 8},
			"0\n",
		},
		{
			"3,9,7,9,10,9,4,9,99,-1,8 outputs 1 if the input is less than 8",
			fields{state: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, input: "7\n", output: &bytes.Buffer{}},
			[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, 1, 8},
			"1\n",
		},
		{
			"3,9,7,9,10,9,4,9,99,-1,8 outputs 0 if the input is not less than 8",
			fields{state: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, input: "8\n", output: &bytes.Buffer{}},
			[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, 0, 8},
			"0\n",
		},
		{
			"3,3,1108,-1,8,3,4,3,99,-1,8 outputs 1 if the input is equal to 8",
			fields{state: []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, input: "8\n", output: &bytes.Buffer{}},
			[]int{3, 3, 1108, 1, 8, 3, 4, 3, 99},
			"1\n",
		},
		{
			"3,3,1108,-1,8,3,4,3,99,-1,8 outputs 0 if the input is not equal to 8",
			fields{state: []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, input: "7\n", output: &bytes.Buffer{}},
			[]int{3, 3, 1108, 0, 8, 3, 4, 3, 99},
			"0\n",
		},
		{
			"3,3,1107,-1,8,3,4,3,99 outputs 1 if the input is less than 8",
			fields{state: []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, input: "7\n", output: &bytes.Buffer{}},
			[]int{3, 3, 1107, 1, 8, 3, 4, 3, 99},
			"1\n",
		},
		{
			"3,3,1107,-1,8,3,4,3,99 outputs 0 if the input is not less than 8",
			fields{state: []int{3, 3, 1107, 0, 8, 3, 4, 3, 99}, input: "8\n", output: &bytes.Buffer{}},
			[]int{3, 3, 1107, 0, 8, 3, 4, 3, 99},
			"0\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := day02.Intcode{
				State: tt.fields.state,
				In:    bufio.NewReader(strings.NewReader(tt.fields.input)),
				Out:   bufio.NewWriter(tt.fields.output),
			}
			gotState := c.Run()
			if !reflect.DeepEqual(gotState, tt.wantState) {
				t.Errorf("Intcode.Run() = %v, want %v", gotState, tt.wantState)
			}

			c.Out.Flush()
			if tt.fields.output.String() != tt.wantOutput {
				t.Errorf("Intcode.Run() output = %v, want %v", tt.fields.output.String(), tt.wantOutput)
			}
		})
	}
}

func TestIntcode_Jump(t *testing.T) {
	type fields struct {
		state  []int
		input  string
		output *bytes.Buffer
	}
	tests := []struct {
		name       string
		fields     fields
		wantOutput string
	}{
		{
			"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9 outputs 0 if the input is 0",
			fields{state: []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, input: "0\n", output: &bytes.Buffer{}},
			"0\n",
		},
		{
			"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9 outputs 1 if the input is not 0",
			fields{state: []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, input: "2\n", output: &bytes.Buffer{}},
			"1\n",
		},
		{
			"3,3,1105,-1,9,1101,0,0,12,4,12,99,1 outputs 0 if the input is 0",
			fields{state: []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, input: "0\n", output: &bytes.Buffer{}},
			"0\n",
		},
		{
			"3,3,1105,-1,9,1101,0,0,12,4,12,99,1 outputs 1 if the input is not 0",
			fields{state: []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, input: "2\n", output: &bytes.Buffer{}},
			"1\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := day02.Intcode{
				State: tt.fields.state,
				In:    bufio.NewReader(strings.NewReader(tt.fields.input)),
				Out:   bufio.NewWriter(tt.fields.output),
			}
			c.Run()
			c.Out.Flush()
			if tt.fields.output.String() != tt.wantOutput {
				t.Errorf("Intcode.Run() output = %v, want %v", tt.fields.output.String(), tt.wantOutput)
			}
		})
	}
}
