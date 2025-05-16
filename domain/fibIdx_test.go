package domain

import (
	"math/big"
	"reflect"
	"testing"
)

func TestCalcFibNum(t *testing.T) {
	tests := []struct {
		input    *big.Int
		expected *big.Int
	}{
		{big.NewInt(0), big.NewInt(0)},
		{big.NewInt(1), big.NewInt(1)},
		{big.NewInt(2), big.NewInt(1)},
		{big.NewInt(3), big.NewInt(2)},
		{big.NewInt(4), big.NewInt(3)},
		{big.NewInt(5), big.NewInt(5)},
		{big.NewInt(6), big.NewInt(8)},
	}

	for _, test := range tests {
		result, err := CalcFibNum(test.input)
		if err != nil {
			t.Errorf("Error calculating Fibonacci number: %v", err)
			continue
		}
		if result.Cmp(test.expected) != 0 {
			t.Errorf("Expected %v, but got %v", test.expected, result)
		}
	}
}

func TestMatrixMul(t *testing.T) {
	m1 := [2][2]*big.Int{{big.NewInt(1), big.NewInt(2)}, {big.NewInt(3), big.NewInt(4)}}
	m2 := [2][2]*big.Int{{big.NewInt(5), big.NewInt(6)}, {big.NewInt(7), big.NewInt(8)}}
	expected := [2][2]*big.Int{{big.NewInt(19), big.NewInt(22)}, {big.NewInt(43), big.NewInt(50)}}

	result := matrixMul(m1, m2)

	if !reflect.DeepEqual(result, expected) {
		// DeepEqualを使って2つの行列が等しいか比較
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestMatrixPow(t *testing.T) {
	m := [2][2]*big.Int{{big.NewInt(1), big.NewInt(1)}, {big.NewInt(1), big.NewInt(0)}}
	n := big.NewInt(5)
	expected := [2][2]*big.Int{{big.NewInt(8), big.NewInt(5)}, {big.NewInt(5), big.NewInt(3)}}

	result := MatrixPow(m, n)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
