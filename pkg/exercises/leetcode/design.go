package leetcode

/*
ATM simula una máquina de cajero automático
LeetCode #2241: Design an ATM Machine
*/

type ATM struct {
	banknotes [5]int64 // Cantidad de billetes: [20, 50, 100, 200, 500]
	values    [5]int   // Valores de los billetes
}

// Constructor crea una nueva instancia del ATM
func Constructor() ATM {
	return ATM{
		banknotes: [5]int64{0, 0, 0, 0, 0},
		values:    [5]int{20, 50, 100, 200, 500},
	}
}

// Deposit deposita billetes en el ATM
func (a *ATM) Deposit(banknotesCount []int) {
	for i := 0; i < 5; i++ {
		a.banknotes[i] += int64(banknotesCount[i])
	}
}

// Withdraw intenta retirar el monto especificado
func (a *ATM) Withdraw(amount int) []int {
	result := [5]int{0, 0, 0, 0, 0}
	remaining := amount

	// Intentar retirar desde el billete de mayor denominación (500) hasta el menor (20)
	for i := 4; i >= 0; i-- {
		if remaining >= a.values[i] && a.banknotes[i] > 0 {
			// Calcular cuántos billetes de esta denominación necesitamos
			count := remaining / a.values[i]
			// No podemos usar más billetes de los que tenemos
			if int64(count) > a.banknotes[i] {
				count = int(a.banknotes[i])
			}

			result[i] = count
			remaining -= count * a.values[i]
		}
	}

	// Si no pudimos completar el monto exacto, rechazar
	if remaining > 0 {
		return []int{-1}
	}

	// Actualizar los billetes en el ATM
	for i := 0; i < 5; i++ {
		a.banknotes[i] -= int64(result[i])
	}

	return result[:]
}
