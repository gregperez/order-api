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

/*
LengthOfLongestSubstring encuentra la longitud de la subcadena sin caracteres repetidos más larga
LeetCode #3: Longest Substring Without Repeating Characters
*/
func LengthOfLongestSubstring(s string) int {
	// Mapa para guardar la última posición vista de cada carácter
	charIndex := make(map[rune]int)

	maxLen := 0
	start := 0 // Inicio de la ventana

	for i, char := range s {
		// Si encontramos el carácter Y está dentro de la ventana actual
		if lastIndex, exists := charIndex[char]; exists && lastIndex >= start {
			// Mover el inicio de la ventana justo después del carácter duplicado
			start = lastIndex + 1
		}

		// Actualizar la última posición de este carácter
		charIndex[char] = i

		// Calcular la longitud de la ventana actual [start, i]
		currentLen := i - start + 1

		// Actualizar máximo si encontramos una ventana más larga
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}

	return maxLen
}

/*
LongestPalindrome encuentra la longitud del palíndromo más largo que se puede construir con las letras dadas
LeetCode #5: Longest Palindrome Substring
*/
func LongestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLen := 0, 0

	expandAroundCenter := func(left, right int) int {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		return right - left - 1
	}

	for i := 0; i < len(s); i++ {
		// Palíndromos impares (centro en un carácter)
		len1 := expandAroundCenter(i, i)
		// Palíndromos pares (centro entre dos caracteres)
		len2 := expandAroundCenter(i, i+1)

		currentLen := maxLens(len1, len2)
		if currentLen > maxLen {
			maxLen = currentLen
			start = i - (currentLen-1)/2
		}
	}

	return s[start : start+maxLen]
}

func maxLens(a, b int) int {
	if a > b {
		return a
	}
	return b
}
