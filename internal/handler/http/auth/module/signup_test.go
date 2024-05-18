package module

import (
	"bytes"
	"crypto-tracker/internal/model/auth"
	ucAuth "crypto-tracker/internal/usecase/auth"
	mockAuth "crypto-tracker/internal/usecase/auth/_mock"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_handler_SignUp(t *testing.T) {
	type args struct {
		value map[string]string
	}
	tests := []struct {
		name string
		args args
		mock func(ctrl *gomock.Controller) ucAuth.UseCase
		want int
	}{
		{
			name: "success",
			args: args{
				value: map[string]string{
					"email":                 "test@test.com",
					"password":              "secret123456!",
					"password_confirmation": "secret123456!",
				},
			},
			mock: func(ctrl *gomock.Controller) ucAuth.UseCase {
				mAuth := mockAuth.NewMockUseCase(ctrl)

				mAuth.EXPECT().SignUp(gomock.Any()).Return(auth.AuthResponse{
					Email: "test@test.com",
					Token: "secret",
					Time:  time.Now().Unix(),
				}, nil)

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
			body, _ := json.Marshal(tt.args.value)
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, url.String(), bytes.NewReader(body))

			h.SignUp(w, r)

			got := w.Result()
			wg.Wait()
			assert.EqualValues(t, tt.want, got.StatusCode)
		})
	}
}
