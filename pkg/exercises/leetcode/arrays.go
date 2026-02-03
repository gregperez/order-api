package leetcode

import "slices"

// TwoSum encuentra dos índices donde la suma de los números sea igual al target
// LeetCode #1: Two Sum
func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int) // valor -> índice

	for i, num := range nums {
		complement := target - num
		if idx, found := seen[complement]; found {
			return []int{idx, i}
		}
		seen[num] = i
	}

	return []int{}
}

// MinimumAbsDifference encuentra todos los pares con la diferencia absoluta mínima
// LeetCode #1200: Minimum Absolute Difference
func MinimumAbsDifference(arr []int) [][]int {
	if len(arr) < 2 {
		return [][]int{}
	}

	slices.Sort(arr)

	// Encontrar la diferencia mínima
	minDiff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if diff < minDiff {
			minDiff = diff
		}
	}

	// Recolectar todos los pares con la diferencia mínima
	var result [][]int
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] == minDiff {
			result = append(result, []int{arr[i], arr[i+1]})
		}
	}

	return result
}

// MinOperations número mínimo de operaciones para hacer la suma del arreglo divisible por K
// LeetCode #3512: Minimum Number of Operations to Make Array Sum Divisible by K
func MinOperations(nums []int, k int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}

	if sum%k == 0 {
		return 0
	}

	operations := sum % k

	if operations > sum {
		return 0
	}

	return operations
}

// MinimumCost dividir un arreglo en subarreglos con costo mínimo
// LeetCode #3010: Divide Array Into Subarrays With Minimum Cost
func MinimumCost(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	if n == 3 {
		return nums[0] + nums[1] + nums[2]
	}

	// the nums[0] is fixed. you have to find another two prefix of subarrays.
	//sort the array from index 1 and return nums[0]+nums[1]+nums[2]
	slices.Sort(nums[1:])

	totalCost := nums[0] + nums[1] + nums[2]

	return totalCost
}
