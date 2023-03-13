package main

import (
	"fmt"
	"reflect"
	"testing"
)

var twoSumTests = []struct {
	nums   []int
	target int
	output []int
}{
	{
		nums:   []int{2, 7, 11, 15},
		target: 9,
		output: []int{0, 1},
	},
	{
		nums:   []int{3, 2, 4},
		target: 6,
		output: []int{1, 2},
	},
	{
		nums:   []int{3, 3},
		target: 6,
		output: []int{0, 1},
	},
}

func TestTwoSumBruteForce(t *testing.T) {
	for _, tt := range twoSumTests {
		t.Run(fmt.Sprintf("TwoSumBruteForce(%v, %d)", tt.nums, tt.target), func(t *testing.T) {
			output := TwoSumBruteForce(tt.nums, tt.target)
			if !reflect.DeepEqual(output, tt.output) {
				t.Errorf("got %d, want %d", output, tt.output)
			}
		})
	}
}

func BenchmarkTwoSumBruteForce(b *testing.B) {
	for n := 0; n < b.N; n++ {
		TwoSumBruteForce([]int{2, 7, 11, 15}, 9)
	}
}
