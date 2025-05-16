package usecase

import (
	"math/big"
	"testing"
)

func TestFibIdxUsecase_CalcFibNum(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *big.Int
		wantErr error
	}{
		{
			name:    "valid input 0",
			input:   "0",
			want:    big.NewInt(0),
			wantErr: nil,
		},
		{
			name:    "invalid input not integer",
			input:   "5.5",
			want:    nil,
			wantErr: ErrInvalidInput,
		},
		{
			name:    "invalid inpput negative integer",
			input:   "-5",
			want:    nil,
			wantErr: ErrInvalidInput,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fu := NewFibIdxUsecase()
			got, err := fu.CalcFibNum(tt.input)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("CalcFibNum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil && tt.want != nil {
				t.Errorf("CalcFibNum() got = %v, want %v", got, tt.want)
				return
			}
			if got != nil && got.Cmp(tt.want) != 0 {
				t.Errorf("CalcFibNum() got = %v, want %v", got, tt.want)
			}
		})
	}
}
