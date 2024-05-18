package module

import (
	"context"
	ucAuth "crypto-tracker/internal/usecase/auth"
	mockAuth "crypto-tracker/internal/usecase/auth/_mock"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_handler_SignOut(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		args args
		mock func(ctrl *gomock.Controller) ucAuth.UseCase
		want int
	}{
		{
			name: "success",
			args: args{},
			mock: func(ctrl *gomock.Controller) ucAuth.UseCase {
				mAuth := mockAuth.NewMockUseCase(ctrl)

				mAuth.EXPECT().SignOut(gomock.Any()).Return(nil)

				return mAuth
			},
			want: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			h := &handler{
				ucAuth: tt.mock(ctrl),
			}

			wg := sync.WaitGroup{}

			url := url.URL{
				Scheme: "http",
				Host:   "test.com",
			}
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, url.String(), nil)
			ctx := r.Context()
			ctx = context.WithValue(ctx, "token", "secret")
			r = r.WithContext(ctx)

			h.SignOut(w, r)

			got := w.Result()
			wg.Wait()
			assert.EqualValues(t, tt.want, got.StatusCode)
		})
	}
}
