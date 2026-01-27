package examples

import (
	"fmt"

	"example.com/order-api/pkg/exercises/leetcode"
	"example.com/order-api/utils"
)

// RunExercisesDemo ejecuta ejemplos de los ejercicios implementados
func RunExercisesDemo() {
	fmt.Println("=== Demo de Ejercicios ===")

	// Demo: Palíndromo
	fmt.Println("\n--- Palíndromos ---")
	testCases := []string{"aja", "aba", "setset", "abcba", "abcta", "arepera", "c", ""}
	for _, word := range testCases {
		fmt.Printf("¿Es palíndromo '%s'? → %t\n", word, utils.IsPalindrome(word))
	}

	// Demo: TwoSum
	fmt.Println("\n--- LeetCode: Two Sum ---")
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("TwoSum(%v, %d) = %v\n", nums, target, leetcode.TwoSum(nums, target))

	// Demo: FirstUniqueChar
	fmt.Println("\n--- LeetCode: First Unique Character ---")
	tests := []string{"swiss", "aabbcc", "Hello", "abacddbec", "aAbBABac"}
	for _, test := range tests {
		result := leetcode.FirstUniqueChar(test)
		if result == nil {
			fmt.Printf("FirstUniqueChar(\"%s\") → nil\n", test)
		} else {
			fmt.Printf("FirstUniqueChar(\"%s\") → \"%s\"\n", test, *result)
		}
	}

	// Demo: ATM
	fmt.Println("\n--- LeetCode: ATM Machine ---")
	atm := leetcode.Constructor()
	fmt.Println("ATM initialized")

	atm.Deposit([]int{0, 0, 1, 2, 1})
	fmt.Println("Deposited [0,0,1,2,1]")

	result := atm.Withdraw(600)
	fmt.Printf("Withdraw 600: %v\n", result)

	atm.Deposit([]int{0, 1, 0, 1, 1})
	fmt.Println("Deposited [0,1,0,1,1]")

	result = atm.Withdraw(600)
	fmt.Printf("Withdraw 600: %v\n", result)

	result = atm.Withdraw(550)
	fmt.Printf("Withdraw 550: %v\n", result)

	// Demo: MinimumAbsDifference
	fmt.Println("\n--- LeetCode: Minimum Absolute Difference ---")
	arr := []int{4, 2, 1, 3}
	fmt.Printf("MinimumAbsDifference(%v) = %v\n", arr, leetcode.MinimumAbsDifference(arr))

	fmt.Println("\n=== Fin del Demo ===")
}

// RunConcurrencyDemo ejecuta el demo de concurrencia
func RunConcurrencyDemo() {
	fmt.Println("=== Demo de Concurrencia ===")
	fmt.Println("Nota: Este demo requiere argumentos de línea de comandos")
	fmt.Println("Uso: go run cmd/server/main.go <key> <value>")
	fmt.Println("=== Fin del Demo ===")
}
