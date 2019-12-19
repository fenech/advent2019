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
			"123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Intcode{
				&day02.Intcode{State: tt.fields.state},
				bufio.NewReader(strings.NewReader(tt.fields.input)),
				bufio.NewWriter(tt.fields.output),
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
