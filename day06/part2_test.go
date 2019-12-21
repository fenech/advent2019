package day06

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestPath(t *testing.T) {
	type args struct {
		orbits map[string]Orbit
		planet string
		path   []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"returns a list of steps to reach COM",
			args{
				map[string]Orbit{"B": {Next: "A"}, "A": {Next: "COM"}},
				"B",
				[]string{},
			},
			[]string{"B", "A", "COM"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Path(tt.args.orbits, tt.args.planet, tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Path() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistance(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"example",
			args{strings.NewReader("COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L\nK)YOU\nI)SAN")},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distance(tt.args.r); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
