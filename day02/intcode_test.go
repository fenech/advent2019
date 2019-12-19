package day02

import (
	"reflect"
	"testing"
)

func TestIntcode_Compute(t *testing.T) {
	type fields struct {
		pointer int
		state   []int
	}
	tests := []struct {
		name      string
		fields    fields
		wantState []int
	}{

		{
			"1,0,0,0,99 becomes 2,0,0,0,99 (1 + 1 = 2).",
			fields{0, []int{1, 0, 0, 0, 99}},
			[]int{2, 0, 0, 0, 99},
		},
		{
			"2,3,0,3,99 becomes 2,3,0,6,99 (3 * 2 = 6).",
			fields{0, []int{2, 3, 0, 3, 99}},
			[]int{2, 3, 0, 6, 99},
		},
		{
			"2,4,4,5,99,0 becomes 2,4,4,5,99,9801 (99 * 99 = 9801).",
			fields{0, []int{2, 4, 4, 5, 99, 0}},
			[]int{2, 4, 4, 5, 99, 9801},
		},
		{
			"1,1,1,4,99,5,6,0,99 becomes 30,1,1,4,2,5,6,0,99.",
			fields{0, []int{1, 1, 1, 4, 99, 5, 6, 0, 99}},
			[]int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Intcode{
				pointer: tt.fields.pointer,
				State:   tt.fields.state,
			}
			if gotState := c.Compute(); !reflect.DeepEqual(gotState, tt.wantState) {
				t.Errorf("Intcode.Compute() = %v, want %v", gotState, tt.wantState)
			}
		})
	}
}
