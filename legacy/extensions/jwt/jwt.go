package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jonloureiro/tiny-bank/legacy/extensions/id"
)

type Token struct {
	AccountId      id.ID
	Token          string
	ExpirationTime time.Time
	IssuedAt       time.Time
}

const (
	_ttl = 15 * time.Minute
)

var ErrTokenUnknownError = errors.New("token unknown error")

func New(AccountId id.ID, privateKey string) (*Token, error) {
	issuedAt := time.Now()
	expirationTime := issuedAt.Add(_ttl)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"account_id": AccountId,
		"exp":        expirationTime.Unix(),
		"iat":        issuedAt.Unix(),
	})
	token, err := jwtToken.SignedString([]byte(privateKey))
	if err != nil {
		return nil, ErrTokenUnknownError
	}
	return &Token{
		AccountId:      AccountId,
		Token:          token,
		ExpirationTime: expirationTime,
		IssuedAt:       issuedAt,
	}, nil
}

func Parse(token, privateKey string) (*Token, error) {
	t, err := jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(privateKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := t.Claims.(*jwt.MapClaims)
	if !ok {
		return nil, ErrTokenUnknownError
	}
	accountId, ok := (*claims)["account_id"].(string)
	if !ok {
		return nil, ErrTokenUnknownError
	}
	expirationTime, ok := (*claims)["exp"].(float64)
	if !ok {
		return nil, ErrTokenUnknownError
	}
	issuedAt, ok := (*claims)["iat"].(float64)
	if !ok {
		return nil, ErrTokenUnknownError
	}
	return &Token{
		AccountId:      id.ID(accountId),
		Token:          token,
		ExpirationTime: time.Unix(int64(expirationTime), 0),
		IssuedAt:       time.Unix(int64(issuedAt), 0),
	}, nil
}

func (t *Token) IsExpired() bool {
	return t.ExpirationTime.Before(time.Now())
}
