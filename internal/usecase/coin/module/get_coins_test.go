package module

import (
	"crypto-tracker/internal/model/coin"
	"crypto-tracker/internal/model/coincap"
	mockCoincap "crypto-tracker/internal/repo/coincap/http/_mock"
	mockSqlite "crypto-tracker/internal/repo/sqlite/_mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_usecase_GetCoins(t *testing.T) {
	type args struct {
		userID int64
	}
	tests := []struct {
		name      string
		args      args
		mock      func(ctrl *gomock.Controller) *usecase
		wantCoins []coin.Coin
		wantErr   bool
	}{
		{
			name: "success",
			args: args{
				userID: 1,
			},
			mock: func(ctrl *gomock.Controller) *usecase {
				mCoincap := mockCoincap.NewMockRepo(ctrl)
				mDB := mockSqlite.NewMockSqlite(ctrl)

				mCoincap.EXPECT().FindRate("indonesian-rupiah").Return(coincap.Rate{
					CoinID:  "indonesian-rupiah",
					RateUsd: 1,
				}, nil)

				mDB.EXPECT().GetCoins(gomock.Any()).Return([]coin.CoinDB{
					{
						ID:      1,
						UserID:  2,
						CoindID: "bitcoin",
					},
				}, nil)

				mCoincap.EXPECT().FindRate(gomock.Any()).Return(coincap.Rate{
					CoinID:  "bitcoin",
					RateUsd: 1,
				}, nil).AnyTimes()

				return &usecase{
					rDB:      mDB,
					rCoincap: mCoincap,
				}
			},
			wantCoins: []coin.Coin{
				{
					Name:    "bitcoin",
					RateIdr: 1,
				},
			},
			wantErr: false,
		},
		{
			name: "success_empty",
			args: args{
				userID: 1,
			},
			mock: func(ctrl *gomock.Controller) *usecase {
				mCoincap := mockCoincap.NewMockRepo(ctrl)
				mDB := mockSqlite.NewMockSqlite(ctrl)

				mCoincap.EXPECT().FindRate("indonesian-rupiah").Return(coincap.Rate{
					CoinID:  "indonesian-rupiah",
					RateUsd: 1,
				}, nil)

				mDB.EXPECT().GetCoins(gomock.Any()).Return([]coin.CoinDB{}, nil)

				return &usecase{
					rDB:      mDB,
					rCoincap: mCoincap,
				}
			},
			wantCoins: nil,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			u := tt.mock(ctrl)
			gotCoins, err := u.GetCoins(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCoins() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equalf(t, tt.wantCoins, gotCoins, "GetCoins() = %v, want %v", gotCoins, tt.wantCoins)
		})
	}
}
