package main

import (
	"fmt"
	"reflect"
	"testing"
)

var runningSumTests = []struct {
	input  []int
	output []int
}{
	{
		input:  []int{1, 2, 3, 4},
		output: []int{1, 3, 6, 10},
	},
	{
		input:  []int{1, 2, 3, -4},
		output: []int{1, 3, 6, 2},
	},
}

func TestRunningSum(t *testing.T) {
	for _, tt := range runningSumTests {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			sums := RunningSum(tt.input)
			if !reflect.DeepEqual(sums, tt.output) {
				t.Errorf("got %d, want %d", sums, tt.output)
			}
		})
	}
}
