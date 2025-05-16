package domain

import (
	"math/big"
	"testing"
)

func BenchmarkCalcFibNum(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CalcFibNum(big.NewInt(1000000))
	}
}
