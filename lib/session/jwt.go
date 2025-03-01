package session

import (
	"os"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

func GenSessionToken(id uint64) (string, error) {
	return lockWithExp(id, time.Now().Add(time.Minute*5).Unix())
}

func GenRefreshToken(id uint64) (string, error) {
	return lockWithExp(id, time.Now().Add(time.Hour).Unix())
}

func lockWithExp(id uint64, exp int64) (string, error) {
	key := os.Getenv("JWT_SIGNATURE_KEY")
	if key == "" {
		key = "AchmadRifai"
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.StandardClaims{Subject: strconv.FormatUint(id, 10), ExpiresAt: exp})
	return token.SignedString([]byte(key))
}
