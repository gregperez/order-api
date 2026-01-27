package leetcode

import "strings"

/*
FirstUniqueChar encuentra el primer carácter único en un string
LeetCode #387: First Unique Character in a String
*/
func FirstUniqueChar(s string) *string {
	// Mapa para contar frecuencias
	count := make(map[rune]int)

	// Convertir a minúsculas para ignorar mayúsculas/minúsculas
	strLower := strings.ToLower(s)

	// Primera pasada: contar ocurrencias
	for _, char := range strLower {
		count[char]++
	}

	// Segunda pasada: encontrar el primer único
	for _, char := range strLower {
		if count[char] == 1 {
			result := string(char)
			return &result
		}
	}

	return nil
}
