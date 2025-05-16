package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"fib_api/usecase"

	"github.com/labstack/echo/v4"
)

// JSONレスポンスを比較できるようにする
func TestFibIdxHandler_HandlerCalcFibNum(t *testing.T) {
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
			wantBody:   `{"result": 0}`,
		},
		{
			name:       "invalid input -1",
			fibIdxStr:  "-1",
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"message":"n must be a non-negative integer"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/fib?n="+tt.fibIdxStr, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			fu := usecase.NewFibIdxUsecase()
			handler := NewFibIdxHandler(fu)
			if err := handler.HandlerCalcFibNum(c); err != nil {
				t.Errorf("HandlerCalcFibNum() error = %v", err)
				return
			}

			fmt.Println("rec = ", rec)

			if rec.Code != tt.wantStatus {
				t.Errorf("HandlerCalcFibNum() got status = %v, want %v", rec.Code, tt.wantStatus)
			}
			if rec.Body.String() != tt.wantBody {
				t.Errorf("HandlerCalcFibNum() got body = %v, want %v", rec.Body.String(), tt.wantBody)
			}
		})
	}
}
