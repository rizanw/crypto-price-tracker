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

func Test_usecase_SignUp(t *testing.T) {
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
					Password: "pwd123pwd",
				},
			},
			mock: func(ctrl *gomock.Controller) rDB.Sqlite {
				mockDB := mockSqlite.NewMockSqlite(ctrl)

				mockDB.EXPECT().FindUser(gomock.Any()).Return(mUser.User{}, nil)
				mockDB.EXPECT().InsertUser(gomock.Any(), gomock.Any()).Return(int64(1), nil)

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
			name: "fail_user_exist",
			args: args{
				in: auth.AuthRequest{
					Email:    "test@test.com",
					Password: "pwd123pwd",
				},
			},
			mock: func(ctrl *gomock.Controller) rDB.Sqlite {
				mockDB := mockSqlite.NewMockSqlite(ctrl)

				mockDB.EXPECT().FindUser(gomock.Any()).Return(mUser.User{
					UserID:   1,
					Email:    "test@test.com",
					Password: "**************",
				}, nil)

				return mockDB
			},
			want:    auth.AuthResponse{},
			wantErr: true,
		},
		{
			name: "fail_insert_error",
			args: args{
				in: auth.AuthRequest{
					Email:    "test@test.com",
					Password: "pwd123pwd",
				},
			},
			mock: func(ctrl *gomock.Controller) rDB.Sqlite {
				mockDB := mockSqlite.NewMockSqlite(ctrl)

				mockDB.EXPECT().FindUser(gomock.Any()).Return(mUser.User{}, nil)
				mockDB.EXPECT().InsertUser(gomock.Any(), gomock.Any()).Return(int64(0), errors.New("test-error"))

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
			got, err := u.SignUp(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equalf(t, tt.want.Email, got.Email, "SignUp(%v)", tt.args.in)
		})
	}
}
