package module

import (
	"context"
	"crypto-tracker/internal/common/session"
	"crypto-tracker/internal/model/coin"
	ucCoin "crypto-tracker/internal/usecase/coin"
	mockCoin "crypto-tracker/internal/usecase/coin/_mock"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_handler_GetCoins(t *testing.T) {
	type args struct{}
	tests := []struct {
		name string
		args args
		mock func(ctrl *gomock.Controller) ucCoin.UseCase
		want int
	}{
		{
			name: "success",
			args: args{},
			mock: func(ctrl *gomock.Controller) ucCoin.UseCase {
				mCoin := mockCoin.NewMockUseCase(ctrl)

				mCoin.EXPECT().GetCoins(gomock.Any()).Return([]coin.Coin{}, nil)

				return mCoin
			},
			want: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			h := &handler{
				ucCoin: tt.mock(ctrl),
			}

			wg := sync.WaitGroup{}

			url := url.URL{
				Scheme: "http",
				Host:   "example.com",
			}
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, url.String(), nil)
			ctx := r.Context()
			ctx = context.WithValue(ctx, "session", session.Session{
				UserID: 1,
				Email:  "test@example.com",
				Expiry: time.Now().Add(1 * time.Hour).Unix(),
			})
			r = r.WithContext(ctx)

			h.GetCoins(w, r)

			got := w.Result()
			wg.Wait()
			assert.EqualValues(t, tt.want, got.StatusCode)
		})
	}
}
