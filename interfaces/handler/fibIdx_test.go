package handler

import (
	"fib_api/usecase"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Cside/jsondiff"
	"github.com/labstack/echo/v4"
)

func TestHandlerCalcFibNum(t *testing.T) {
	tests := []struct {
		name       string
		fibIdxStr  string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "valid input 0",
			fibIdxStr:  "0",
			wantStatus: http.StatusOK,
			wantBody:   `{"result":0}`,
		},
		{
			name:       "invalid input not integer",
			fibIdxStr:  "5.5",
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"message":"n must be a non-negative integer"}`,
		},
		{
			name:       "too large input",
			fibIdxStr:  "200001",
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"message":"n is too large, please use less than 200000"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/fib?n="+tt.fibIdxStr, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			fu := usecase.NewFibIdxUsecase()
			fh := NewFibIdxHandler(fu)

			if err := fh.HandlerCalcFibNum(c); err != nil {
				t.Errorf("HandlerCalcFibNum() error = %v", err)
				return
			}

			//ステータスコードを確認
			if rec.Code != tt.wantStatus {
				t.Errorf("HandlerCalcFibNum() got status = %v, want %v", rec.Code, tt.wantStatus)
			}

			//レスポンスの中身を確認
			if diff := jsondiff.Diff([]byte(rec.Body.String()), []byte(tt.wantBody)); diff != "" {
				t.Errorf("HandlerCalcFibNum() got body = %v, want %v", rec.Body.String(), tt.wantBody)
				t.Errorf("Diff: %v", diff)
			}
		})
	}
}
