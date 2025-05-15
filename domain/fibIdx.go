package domain

import "math/big"

func CalcFibNum(fibIdx *big.Int) (*big.Int, error) {
	//n = 0, 1のときのフィボナッチ数列の値を返す
	if fibIdx == big.NewInt(0) {
		return big.NewInt(0), nil
	}
	if fibIdx == big.NewInt(1) {
		return big.NewInt(1), nil
	}

	//n = 2以上のときのフィボナッチ数列の値を返す
	fib0 := big.NewInt(0)
	fib1 := big.NewInt(1)

	fib_matrix := [2][2]*big.Int{{big.NewInt(1), big.NewInt(1)}, {big.NewInt(1), big.NewInt(0)}}
	fib_matrix_pow := MatrixPow(fib_matrix, fibIdx)

	fibNum := new(big.Int)
	fibNum.Add(fibNum, fib0.Mul(fib0, fib_matrix_pow[1][0]))
	fibNum.Add(fibNum, fib1.Mul(fib1, fib_matrix_pow[1][1]))

	return fibNum, nil
}
