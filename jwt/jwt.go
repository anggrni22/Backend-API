package jwt

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

const minSecretSize = 32
const defaultExpire = 168 // 1 week

type Info struct {
	Secret, Issuer string
	ExpireHour     int
}

func NewJWTMaker(secretKey string, issuer string, expireHour int) (Maker, error) {
	if len(secretKey) < minSecretSize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretSize)
	}

	if expireHour == 0 {
		expireHour = defaultExpire
	}

	return &Info{secretKey, issuer, expireHour}, nil
}

func (maker *Info) CreateToken(payload interface{}) (string, error) {
	rf := reflect.ValueOf(payload)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["StandardClaims"] = jwt.StandardClaims{
		Issuer:    maker.Issuer,
		ExpiresAt: time.Now().Add(time.Duration(maker.ExpireHour) * time.Hour).Unix(),
		IssuedAt:  time.Now().Unix(),
	}

	for i := 0; i < rf.NumField(); i++ {
		claims[rf.Type().Field(i).Name] = rf.Field(i).Interface()
	}

	signedToken, err := token.SignedString([]byte(maker.Secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (maker *Info) VerifyToken(token string) (*jwt.Token, error) {
	token, _ = StripBearerPrefixFromTokenString(token)
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.Secret), nil
	})
}

func StripBearerPrefixFromTokenString(tok string) (string, error) {
	if len(tok) > 6 && strings.ToUpper(tok[0:7]) == "BEARER " {
		return tok[7:], nil
	}
	return tok, nil
}
