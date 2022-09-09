package auth

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestJWTTokenMaker(t *testing.T) {

	tests := []struct {
		name string
		args utilArgs
	}{
		{
			name: "Should create token without error",
			args: utilArgs{
				accountUUID:          "account-123",
				sessionUUID:          "session-123",
				accessTokenDuration:  time.Second,
				refreshTokenDuration: time.Second * 2,
			},
		},
		{
			name: "Should return error with an invalid privateKey",
			args: utilArgs{
				withoutPrivateKey: true,
				wantErr:           true,
				wantErrValue:      errInvalidPrivateKey,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.tokenType = tokenTypeJWT
			validateTokenMaker(t, tt.args)
		})
	}
}

func Test_jwtAuth_VerifyToken(t *testing.T) {

	tests := []struct {
		name    string
		args    utilArgs
		wantErr bool
	}{
		{
			name: "Should pass without error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := getConfig(t, tt.args)
			tt.args.tokenType = tokenTypeJWT
			maker, err := NewAuthToken(cfg.App.Auth)
			require.NoError(t, err)

			token, tokenPayload := createTestAccessToken(t, maker, tt.args)

			gotPayload, err := maker.VerifyToken(token)
			if (err != nil) != tt.wantErr {
				t.Errorf("jwtAuth.VerifyToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, tt.args.sessionUUID, gotPayload.SessionUUID)
			require.Equal(t, tt.args.accountUUID, gotPayload.AccountUUID)
			require.WithinDuration(t, tokenPayload.IssuedAt, gotPayload.IssuedAt, 1*time.Second)
			require.WithinDuration(t, tokenPayload.ExpiredAt, gotPayload.ExpiredAt, 1*time.Second)
		})
	}
}
