package domain

import "math/big"

func CalcFibNum(fibIdx *big.Int) (*big.Int, error) {
	//[[1, 1], [1, 0]]^fibIdx の行列対角成分がフィボナッチ数列の一般項となる
	fib_matrix := [2][2]*big.Int{{big.NewInt(1), big.NewInt(1)}, {big.NewInt(1), big.NewInt(0)}}
	fib_matrix_pow := MatrixPow(fib_matrix, fibIdx)

	fibNum := fib_matrix_pow[0][1]

	return fibNum, nil
}
