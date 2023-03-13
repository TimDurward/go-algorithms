# [Running Sum of 1d Array](https://leetcode.com/problems/running-sum-of-1d-array/) (1480)

### Problem

Given an array `nums`. We define a running sum of an array as `runningSum[i] = sum(nums[0]…nums[i])`.

Return the running sum of `nums`.

```
func RunningSum(nums []int) []int
```

Example 1:

```
Input: nums = [1,2,3,4]
Output: [1,3,6,10]
Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].
```

Example 2:

```
Input: nums = [1,1,1,1,1]
Output: [1,2,3,4,5]
Explanation: Running sum is obtained as follows: [1, 1+1, 1+1+1, 1+1+1+1, 1+1+1+1+1].
```

Example 3:

```
Input: nums = [3,1,2,10,1]
Output: [3,4,6,16,17]
```

### Solution

Create a secondary array named sums, this array will hold all of the running-sum values that are evaluated from the input (nums array). After that, we can then iterate over the input array and sum up the current index of num with the last index of sums.

1. define an array `sums`
2. initialize the first element `sums` with value of `0`
3. iterate over `nums` and evaluate `x = (sums[-1] + nums[i])`
4. append value `x` to `sums`
5. once iteration of `nums` is complete remove first element in `sums` before returning `sums`
