package usecase

import (
	"errors"
	service "fib_api/domain"
	"math/big"
	"strconv"
)

type FibIdxUsecase interface {
	CalcFibNum(fibIdxStr string) (*big.Int, error)
}

type fibIdxUsecase struct{}

func NewFibIdxUsecase() FibIdxUsecase {
	return &fibIdxUsecase{}
}

func (fu *fibIdxUsecase) CalcFibNum(fibIdxStr string) (*big.Int, error) {
	//入力が大きすぎるものは型変換できないので、エラーを返す
	if len(fibIdxStr) > len(strconv.Itoa(MAXINPUTNUM)) {
		return nil, ErrTooLargeInput
	}

	//入力を整数に変換
	fibIdx := new(big.Int)
	if _, ok := fibIdx.SetString(fibIdxStr, 10); !ok {
		//整数に変換できない場合はエラーを返す
		return nil, ErrInvalidInput
	}

	if fibIdx.Cmp(big.NewInt(0)) < 0 {
		//fibIdxが負の整数の場合はエラーを返す
		return nil, ErrInvalidInput
	}

	if fibIdx.Cmp(big.NewInt(MAXINPUTNUM)) > 0 {
		//fibIdxがMAXINPUTNUMより大きい場合はエラーを返す
		return nil, ErrTooLargeInput
	}

	fibNum, err := service.CalcFibNum(fibIdx)
	if err != nil {
		return nil, errors.New("failed to calculate Fibonacci number")
	}

	return fibNum, nil
}
