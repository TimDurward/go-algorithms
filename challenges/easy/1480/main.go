package main

// RunningSum returns the running sum of the values in a slice, adding the current value to the preceding values.
func RunningSum(nums []int) []int {
	sums := []int{0}
	for _, n := range nums {
		sums = append(sums, sums[len(sums)-1]+n)
	}

	return sums[1:]
}
