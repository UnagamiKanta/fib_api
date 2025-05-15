package domain

import "math/big"

// 2X2行列同士の掛け算を定義
func matrixMul(m1 [2][2]*big.Int, m2 [2][2]*big.Int) [2][2]*big.Int {
	m00 := new(big.Int)
	m01 := new(big.Int)
	m10 := new(big.Int)
	m11 := new(big.Int)

	m00.Add(m00, m1[0][0].Mul(m1[0][0], m2[0][0]))
	m00.Add(m00, m1[0][1].Mul(m1[0][1], m2[1][0]))
	m01.Add(m01, m1[0][0].Mul(m1[0][0], m2[0][1]))
	m01.Add(m01, m1[0][1].Mul(m1[0][1], m2[1][1]))
	m10.Add(m10, m1[1][0].Mul(m1[1][0], m2[0][0]))
	m10.Add(m10, m1[1][1].Mul(m1[1][1], m2[1][0]))
	m11.Add(m11, m1[1][0].Mul(m1[1][0], m2[0][1]))
	m11.Add(m11, m1[1][1].Mul(m1[1][1], m2[1][1]))

	return [2][2]*big.Int{{m00, m01}, {m10, m11}}
}

// 2^30乗までの行列の累乗を計算する
func MatrixPow(m [2][2]*big.Int, n *big.Int) [2][2]*big.Int {
	pm := m
	ans := [2][2]*big.Int{{big.NewInt(1), big.NewInt(0)}, {big.NewInt(0), big.NewInt(1)}}

	for i := 0; i < 30; i++ {
		if n.Bit(i) == 1 {
			ans = matrixMul(ans, pm)
		}
		pm = matrixMul(pm, pm)
	}

	return ans
}
