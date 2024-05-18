package module

import (
	"crypto-tracker/internal/common/session"
	"testing"
)

func Test_usecase_SignOut(t *testing.T) {
	session.Sessions["session-test"] = session.Session{}
	type args struct {
		sessionKey string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				sessionKey: "session-test",
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				sessionKey: "session-test-2",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{}
			if err := u.SignOut(tt.args.sessionKey); (err != nil) != tt.wantErr {
				t.Errorf("SignOut() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
