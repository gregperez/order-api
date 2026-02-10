package leetcode

import (
	"math"
	"strconv"
	"strings"
)

/*
Reverse revierte un entero
LeetCode #7: Reverse Integer
*/
func Reverse(x int) int {
	var slc []int
	isNegative := 1
	if x < 0 {
		x = -x
		isNegative = -1
	}
	for x > 0 {
		digit := x % 10
		slc = append(slc, digit)
		x /= 10
	}
	num, err := arrayToInt(slc)
	if err != nil {
		return 0
	}

	result := num * isNegative

	// Validar overflow de 32 bits
	if result < math.MinInt32 || result > math.MaxInt32 {
		return 0
	}

	return result
}

func arrayToInt(arr []int) (int, error) {
	strArr := make([]string, len(arr))
	for i, v := range arr {
		strArr[i] = strconv.Itoa(v)
	}

	combinedStr := strings.Join(strArr, "")

	result, err := strconv.Atoi(combinedStr)
	if err != nil {
		return 0, err
	}

	return result, nil
}
