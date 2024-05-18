package module

import (
	"crypto-tracker/internal/config"
	"crypto-tracker/internal/model/auth"
	mUser "crypto-tracker/internal/model/user"
	rDB "crypto-tracker/internal/repo/sqlite"
	mockSqlite "crypto-tracker/internal/repo/sqlite/_mock"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_usecase_SignIn(t *testing.T) {
	type args struct {
		in auth.AuthRequest
	}
	tests := []struct {
		name    string
		args    args
		mock    func(ctrl *gomock.Controller) rDB.Sqlite
		want    auth.AuthResponse
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				in: auth.AuthRequest{
					Email:    "test@test.com",
					Password: "qwerty12345!",
				},
			},
			mock: func(ctrl *gomock.Controller) rDB.Sqlite {
				mockDB := mockSqlite.NewMockSqlite(ctrl)

				mockDB.EXPECT().FindUser(gomock.Any()).Return(mUser.User{
					UserID:   1,
					Email:    "test@test.com",
					Password: "$2a$10$yFYaaLlj.Zy5qmOmWfkL0ew.J0g9b/ocdhKFQ8iVYH/7oJem/yWk2",
				}, nil)

				return mockDB
			},
			want: auth.AuthResponse{
				Email: "test@test.com",
				Token: "",
				Time:  0,
			},
			wantErr: false,
		},
		{
			name: "fail_invalid_password",
			args: args{
				in: auth.AuthRequest{
					Email:    "test@test.com",
					Password: "qwerty12345",
				},
			},
			mock: func(ctrl *gomock.Controller) rDB.Sqlite {
				mockDB := mockSqlite.NewMockSqlite(ctrl)

				mockDB.EXPECT().FindUser(gomock.Any()).Return(mUser.User{
					UserID:   1,
					Email:    "test@test.com",
					Password: "$2a$10$yFYaaLlj.Zy5qmOmWfkL0ew.J0g9b/ocdhKFQ8iVYH/7oJem/yWk2",
				}, nil)

				return mockDB
			},
			want:    auth.AuthResponse{},
			wantErr: true,
		},
		{
			name: "fail_user_not_found",
			args: args{
				in: auth.AuthRequest{
					Email:    "test@test.com",
					Password: "qwerty12345",
				},
			},
			mock: func(ctrl *gomock.Controller) rDB.Sqlite {
				mockDB := mockSqlite.NewMockSqlite(ctrl)

				mockDB.EXPECT().FindUser(gomock.Any()).Return(mUser.User{}, nil)

				return mockDB
			},
			want:    auth.AuthResponse{},
			wantErr: true,
		},
		{
			name: "fail_find_user_error",
			args: args{
				in: auth.AuthRequest{
					Email:    "test@test.com",
					Password: "qwerty12345",
				},
			},
			mock: func(ctrl *gomock.Controller) rDB.Sqlite {
				mockDB := mockSqlite.NewMockSqlite(ctrl)

				mockDB.EXPECT().FindUser(gomock.Any()).Return(mUser.User{}, errors.New("test-error"))

				return mockDB
			},
			want:    auth.AuthResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			u := &usecase{
				rDB:     tt.mock(ctrl),
				confJWT: config.JWTConfig{Secret: make([]byte, 12)},
			}
			got, err := u.SignIn(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equalf(t, tt.want.Email, got.Email, "SignIn(%v)", tt.args.in)
		})
	}
}
