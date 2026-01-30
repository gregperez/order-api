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
	result := [][]int{}
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
