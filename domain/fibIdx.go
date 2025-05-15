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
	last2Num := big.NewInt(0)
	last1Num := big.NewInt(1)
	fibNum := big.NewInt(0)

	for i := big.NewInt(2); i.Cmp(fibIdx) <= 0; i.Add(i, big.NewInt(1)) {
		fibNum.Add(last1Num, last2Num)
		last2Num.Set(last1Num)
		last1Num.Set(fibNum)
	}

	return fibNum, nil
}
