package domain

import (
	"math/big"
)

// 2X2行列同士の掛け算を定義
func matrixMul(a, b [2][2]*big.Int) [2][2]*big.Int {
	return [2][2]*big.Int{
		{
			new(big.Int).Add(
				new(big.Int).Mul(a[0][0], b[0][0]),
				new(big.Int).Mul(a[0][1], b[1][0]),
			),
			new(big.Int).Add(
				new(big.Int).Mul(a[0][0], b[0][1]),
				new(big.Int).Mul(a[0][1], b[1][1]),
			),
		},
		{
			new(big.Int).Add(
				new(big.Int).Mul(a[1][0], b[0][0]),
				new(big.Int).Mul(a[1][1], b[1][0]),
			),
			new(big.Int).Add(
				new(big.Int).Mul(a[1][0], b[0][1]),
				new(big.Int).Mul(a[1][1], b[1][1]),
			),
		},
	}
}

// 行列の累乗を計算する関数：計算量 O(log n)
func MatrixPow(m [2][2]*big.Int, n *big.Int) [2][2]*big.Int {
	ans := [2][2]*big.Int{{big.NewInt(1), big.NewInt(0)}, {big.NewInt(0), big.NewInt(1)}}
	pm := m
	// n.BitLen() は n を 2 進数で表したときの長さ
	for i := 0; i < n.BitLen(); i++ {
		if n.Bit(i) == 1 {
			ans = matrixMul(ans, pm)
		}
		pm = matrixMul(pm, pm)
	}
	return ans
}

func fibDoubling(n *big.Int) (*big.Int, *big.Int) {
	if n.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0), big.NewInt(1)
	}
	if n.Cmp(big.NewInt(1)) == 0 {
		return big.NewInt(1), big.NewInt(1)
	}

	k := new(big.Int).Rsh(n, 1)

	fk, fk1 := fibDoubling(k) // F_k, F_{k+1}

	// twoFk1  = 2 * F_{k+1}
	twoFk1 := new(big.Int).Lsh(fk1, 1)
	// t = (2 * F_{k+1} - F_k)
	t := new(big.Int).Sub(twoFk1, fk)
	// c = F_{2k} = F_k * (2*F_{k+1} – F_k)
	c := new(big.Int).Mul(fk, t)

	squaredFk := new(big.Int).Mul(fk, fk)
	squaredFk1 := new(big.Int).Mul(fk1, fk1)
	// d = F_{2k+1} = F_k^2 + F_{k+1}^2
	d := new(big.Int).Add(squaredFk, squaredFk1)

	if n.Bit(0) == 0 { // n が偶数の場合
		return c, d
	}

	return d, new(big.Int).Add(c, d)
}
