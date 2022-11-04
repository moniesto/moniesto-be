package token

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/moniesto/moniesto-be/util"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomString(6)
	user_id := util.CreateID()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(GeneralPaylod{
		UserPayload: UserPayload{
			ID:       user_id,
			Username: username,
		},
	}, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.User.Username)
	require.Equal(t, user_id, payload.User.ID)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomString(6)
	user_id := util.CreateID()
	duration := -time.Minute

	token, err := maker.CreateToken(GeneralPaylod{
		UserPayload: UserPayload{
			ID:       user_id,
			Username: username,
		},
	}, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	username := util.RandomString(6)
	user_id := util.CreateID()
	duration := time.Minute

	payload, err := NewPayload(GeneralPaylod{
		UserPayload: UserPayload{
			ID:       user_id,
			Username: username,
		},
	}, duration)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}
