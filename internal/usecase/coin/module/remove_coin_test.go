package module

import (
	"crypto-tracker/internal/model/coin"
	mockSqlite "crypto-tracker/internal/repo/sqlite/_mock"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_usecase_RemoveCoin(t *testing.T) {
	type args struct {
		userID int64
		coin   string
	}
	tests := []struct {
		name    string
		args    args
		mock    func(ctrl *gomock.Controller) *usecase
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				userID: 1,
				coin:   "bitcoin",
			},
			mock: func(ctrl *gomock.Controller) *usecase {
				mDB := mockSqlite.NewMockSqlite(ctrl)

				mDB.EXPECT().GetCoins(gomock.Any()).Return([]coin.CoinDB{
					{
						ID:      1,
						UserID:  1,
						CoindID: "bitcoin",
					},
				}, nil)
				mDB.EXPECT().DeleteCoin(gomock.Any(), gomock.Any()).Return(nil)

				return &usecase{
					rDB:      mDB,
					rCoincap: nil,
				}
			},
			wantErr: false,
		},
		{
			name: "fail_no_coin_for_user",
			args: args{
				userID: 1,
				coin:   "bitcoin",
			},
			mock: func(ctrl *gomock.Controller) *usecase {
				mDB := mockSqlite.NewMockSqlite(ctrl)
				mDB.EXPECT().GetCoins(gomock.Any()).Return([]coin.CoinDB{}, nil)

				return &usecase{
					rDB:      mDB,
					rCoincap: nil,
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			u := tt.mock(ctrl)
			if err := u.RemoveCoin(tt.args.userID, tt.args.coin); (err != nil) != tt.wantErr {
				t.Errorf("RemoveCoin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
