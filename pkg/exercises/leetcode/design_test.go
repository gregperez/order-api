package leetcode

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want ATM
	}{
		{
			name: "Test Constructor",
			want: ATM{
				banknotes: [5]int64{0, 0, 0, 0, 0},
				values:    [5]int{20, 50, 100, 200, 500},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); got != tt.want {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestATM_Deposit(t *testing.T) {
	tests := []struct {
		name              string
		initialBanknotes  [5]int64
		deposit           []int
		expectedBanknotes [5]int64
	}{
		{
			name:              "Deposit Test 1",
			initialBanknotes:  [5]int64{0, 0, 0, 0, 0},
			deposit:           []int{1, 2, 3, 4, 5},
			expectedBanknotes: [5]int64{1, 2, 3, 4, 5},
		},
		{
			name:              "Deposit Test 2",
			initialBanknotes:  [5]int64{1, 1, 1, 1, 1},
			deposit:           []int{2, 2, 2, 2, 2},
			expectedBanknotes: [5]int64{3, 3, 3, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			atm := ATM{banknotes: tt.initialBanknotes}
			atm.Deposit(tt.deposit)
			if atm.banknotes != tt.expectedBanknotes {
				t.Errorf("After Deposit(%v), banknotes = %v, want %v", tt.deposit, atm.banknotes, tt.expectedBanknotes)
			}
		})
	}
}

func TestATM_Withdraw(t *testing.T) {
	initialBanknotesValues := [5]int{20, 50, 100, 200, 500}
	tests := []struct {
		name              string
		initialBanknotes  [5]int64
		initialValues     [5]int
		withdrawAmount    int
		expectedResult    []int
		expectedBanknotes [5]int64
	}{
		{
			name:              "Withdraw Test 1",
			initialBanknotes:  [5]int64{1, 1, 1, 1, 1},
			initialValues:     initialBanknotesValues,
			withdrawAmount:    600,
			expectedResult:    []int{0, 0, 1, 0, 1},
			expectedBanknotes: [5]int64{1, 1, 0, 1, 0},
		},
		{
			name:              "Withdraw Test 2 - Insufficient Funds",
			initialBanknotes:  [5]int64{0, 0, 0, 0, 0},
			withdrawAmount:    100,
			expectedResult:    []int{-1},
			expectedBanknotes: [5]int64{0, 0, 0, 0, 0},
		},
		{
			name:              "Withdraw Test 3 - Exact Change",
			initialBanknotes:  [5]int64{2, 2, 2, 2, 2},
			initialValues:     initialBanknotesValues,
			withdrawAmount:    370,
			expectedResult:    []int{1, 1, 1, 1, 0},
			expectedBanknotes: [5]int64{1, 1, 1, 1, 2},
		},
		{
			name:              "Withdraw Test 4 - Amount Not Possible exceeding available",
			initialBanknotes:  [5]int64{2, 0, 0, 0, 0},
			initialValues:     initialBanknotesValues,
			withdrawAmount:    60,
			expectedResult:    []int{-1},
			expectedBanknotes: [5]int64{2, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			atm := ATM{banknotes: tt.initialBanknotes, values: tt.initialValues}
			result := atm.Withdraw(tt.withdrawAmount)
			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("Withdraw(%d) = %v, want %v", tt.withdrawAmount, result, tt.expectedResult)
			}
			if atm.banknotes != tt.expectedBanknotes {
				t.Errorf("After Withdraw(%d), banknotes = %v, want %v", tt.withdrawAmount, atm.banknotes, tt.expectedBanknotes)
			}
		})
	}
}
