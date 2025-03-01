package session

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"track/lib"
	"track/lib/db"
	"track/lib/repo"

	arrayutils "github.com/AchmadRifai/array-utils"
	jwt "github.com/golang-jwt/jwt/v4"
)

func ValidationRole(w http.ResponseWriter, r *http.Request, roles []string) bool {
	log.Printf("[%s] %s", r.Method, r.RequestURI)
	next := true
	claim := ParseToken(r)
	id, err := strconv.ParseUint(claim["sub"].(string), 10, 64)
	if err != nil {
		panic(err)
	}
	if r.RequestURI == "/api/logout" {
		PutSessionToResponse(w, id)
	}
	db, err := db.DbConn()
	defer lib.CloseDb(w, db)
	if err != nil {
		panic(err)
	}
	tx, err := db.Begin()
	defer lib.TxClose(tx, w)
	if err != nil {
		panic(err)
	}
	user, err := repo.UserWithRoleByUserId(tx, id)
	if err != nil {
		panic(err)
	}
	if len(roles) > 0 && arrayutils.AllOf(roles, func(v string, _ int) bool { return v != user.Role.Nm }) {
		next = false
		panic("invalid role")
	}
	return next
}

func ParseToken(r *http.Request) jwt.MapClaims {
	authorizationHeader := r.Header.Get("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		panic("invalid token")
	}
	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		panic(err)
	}
	if !token.Valid {
		panic("token invalid")
	}
	return token.Claims.(jwt.MapClaims)
}

func keyFunc(token *jwt.Token) (any, error) {
	if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("signing method invalid")
	} else if method != jwt.SigningMethodHS512 {
		return nil, errors.New("signing method invalid")
	}
	key := os.Getenv("JWT_SIGNATURE_KEY")
	if key == "" {
		key = "AchmadRifai"
	}
	return []byte(key), nil
}

func PutSessionToResponse(w http.ResponseWriter, id uint64) {
	sessionToken, err := GenSessionToken(id)
	if err != nil {
		panic(err)
	}
	w.Header().Add("sessionToken", sessionToken)
	refreshToken, err := GenRefreshToken(id)
	if err != nil {
		panic(err)
	}
	w.Header().Add("refreshToken", refreshToken)
}

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
