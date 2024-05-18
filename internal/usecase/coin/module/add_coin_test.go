package module

import (
	"crypto-tracker/internal/model/coin"
	"crypto-tracker/internal/model/coincap"
	mockCoincap "crypto-tracker/internal/repo/coincap/http/_mock"
	mockSqlite "crypto-tracker/internal/repo/sqlite/_mock"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_usecase_AddCoin(t *testing.T) {
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
				mCoincap := mockCoincap.NewMockRepo(ctrl)
				mCoincap.EXPECT().FindRate(gomock.Any()).Return(coincap.Rate{
					CoinID: "bitcoin",
				}, nil)

				mDB := mockSqlite.NewMockSqlite(ctrl)
				mDB.EXPECT().GetCoins(gomock.Any()).Return(nil, nil)
				mDB.EXPECT().InsertCoin(gomock.Any(), gomock.Any()).Return(nil)

				return &usecase{
					rDB:      mDB,
					rCoincap: mCoincap,
				}
			},
			wantErr: false,
		},
		{
			name: "fail_coin_not_found",
			args: args{
				userID: 1,
				coin:   "testcoin",
			},
			mock: func(ctrl *gomock.Controller) *usecase {
				mCoincap := mockCoincap.NewMockRepo(ctrl)
				mCoincap.EXPECT().FindRate(gomock.Any()).Return(coincap.Rate{}, nil)

				mDB := mockSqlite.NewMockSqlite(ctrl)

				return &usecase{
					rDB:      mDB,
					rCoincap: mCoincap,
				}
			},
			wantErr: true,
		},
		{
			name: "fail_duplicate_coin",
			args: args{
				userID: 1,
				coin:   "bitcoin",
			},
			mock: func(ctrl *gomock.Controller) *usecase {
				mCoincap := mockCoincap.NewMockRepo(ctrl)
				mCoincap.EXPECT().FindRate(gomock.Any()).Return(coincap.Rate{
					CoinID: "bitcoin",
				}, nil)

				mDB := mockSqlite.NewMockSqlite(ctrl)
				mDB.EXPECT().GetCoins(gomock.Any()).Return([]coin.CoinDB{
					{
						ID:      1,
						UserID:  2,
						CoindID: "bitcoin",
					},
				}, nil)

				return &usecase{
					rDB:      mDB,
					rCoincap: mCoincap,
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
			if err := u.AddCoin(tt.args.userID, tt.args.coin); (err != nil) != tt.wantErr {
				t.Errorf("AddCoin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
