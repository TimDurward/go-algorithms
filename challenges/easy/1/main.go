package main

func TwoSumBruteForce(nums []int, target int) []int {
	for idx1, num1 := range nums {
		for idx2, num2 := range nums {
			if (num1+num2) == target && (idx1 != idx2) {
				return []int{idx1, idx2}
			}
		}
	}
	return nil
}

// TODO - reduce time-complexity to O(n) compared to brute force example O(n2)
// try adding incorporating a Hash Table and pass through nums only once
