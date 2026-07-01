# 💻 Ejercicios de Código

Este directorio contiene soluciones organizadas a problemas de plataformas de coding.

## 📚 LeetCode

### Arrays
- **#1 Two Sum** - `arrays.go`
- **#1200 Minimum Absolute Difference** - `arrays.go`
- **#3010 Divide Array Into Subarrays With Minimum Cost** - `arrays.go`
- **3512 Minimum Operations to Make Array Sum Divisible by K** - `arrays.go`

### Strings
- **#387 First Unique Character in a String** - `strings.go`
- **#3 Longest Substring Without Repeating Characters** - `strings.go`
- **#5 Longest Palindromic Substring** - `strings.go`

### Design
- **#2241 Design an ATM Machine** - `design.go`

### Linked Lists
- **#2 Add Two Numbers** - `linked_lists.go`
- **3363 Convert Doubly Linked List to Array** - `linked_lists.go`

### Integers
- **#7 Reverse Integer** - `integers.go`
- **#9 IsPalindrome Number** - `integers.go`

### Tests
Todos los ejercicios tienen tests unitarios en `*_test.go`

```bash
# Ejecutar tests de LeetCode
go test ./pkg/exercises/leetcode/ -v

# Ejecutar un test específico
go test ./pkg/exercises/leetcode/ -v -run TestTwoSum
```

## 🎯 HackerRank

Preparado para agregar problemas en `hackerrank/problems.go`

## ➕ Agregar Nuevos Ejercicios

### 1. Crear el archivo apropiado
```bash
# Para LeetCode
vi pkg/exercises/leetcode/nueva_categoria.go

# Para HackerRank
vi pkg/exercises/hackerrank/nueva_categoria.go
```

### 2. Formato de la función
```go
package leetcode

// NombreDelProblema descripción breve
// LeetCode #123: Nombre Completo del Problema
// Dificultad: Easy/Medium/Hard
func NombreDelProblema(params tipo) tipo {
    // Tu implementación
    return resultado
}
```

### 3. Agregar tests
```go
package leetcode

import "testing"

func TestNombreDelProblema(t *testing.T) {
    tests := []struct {
        name string
        input tipo
        want tipo
    }{
        {"Example 1", input1, expected1},
        {"Example 2", input2, expected2},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := NombreDelProblema(tt.input)
            if got != tt.want {
                t.Errorf("got %v, want %v", got, tt.want)
            }
        })
    }
}
```

### 4. Opcional: Agregar demo
Agregar ejemplo en `examples/exercises_demo.go`

## 📊 Estadísticas

- **Total ejercicios LeetCode:** 5
- **Total ejercicios HackerRank:** 0
- **Cobertura de tests:** 100%

## 🎓 Categorías Disponibles

### LeetCode
- ✅ Arrays
- ✅ Strings  
- ✅ Design
- ⏳ Trees (por agregar)
- ✅ Linked Lists
- ⏳ Dynamic Programming (por agregar)

### HackerRank
- ⏳ Algoritmos (por agregar)
- ⏳ Estructuras de datos (por agregar)
- ⏳ Matemáticas (por agregar)

