package day06

import (
	"io"
	"strings"
	"testing"
)

func TestChecksum(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name  string
		args  args
		wantC int
	}{
		{"example", args{strings.NewReader("COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L\n")}, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := Checksum(tt.args.r); gotC != tt.wantC {
				t.Errorf("Checksum() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
