package jwt_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/jonloureiro/tiny-bank/legacy/extensions/id"
	"github.com/jonloureiro/tiny-bank/legacy/extensions/jwt"
)

const (
	privateKey     = "s3cr3t"
	fifteenMinutes = 15 * time.Minute
)

func TestNewToken(t *testing.T) {
	want := reflect.TypeOf(&jwt.Token{})
	accountID := id.New()
	token, err := jwt.New(accountID, privateKey)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	got := reflect.TypeOf(token)
	if want != got {
		t.Errorf("expected %v, got %v", want, got)
	}
	if token.AccountId != accountID {
		t.Errorf("expected %v, got %v", accountID, token.AccountId)
	}
	if token.ExpirationTime.Sub(token.IssuedAt) != fifteenMinutes {
		t.Errorf("expected %v, got %v", fifteenMinutes, token.ExpirationTime.Sub(token.IssuedAt))
	}
}

func TestParseToken(t *testing.T) {
	token, err := jwt.New(id.New(), privateKey)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	parsedToken, err := jwt.Parse(token.Token, privateKey)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if token.AccountId != parsedToken.AccountId {
		t.Errorf("[AccountId] expected %v, got %v", token.AccountId, parsedToken.AccountId)
	}
	if token.Token != parsedToken.Token {
		t.Errorf("[AccountId] expected %v, got %v", token.Token, parsedToken.Token)
	}
	if token.ExpirationTime.Unix() != parsedToken.ExpirationTime.Unix() {
		t.Errorf("[ExpirationTime] expected %v, got %v", token.ExpirationTime, parsedToken.ExpirationTime)
	}
	if token.IssuedAt.Unix() != parsedToken.IssuedAt.Unix() {
		t.Errorf("[IssuedAt] expected %v, got %v", token.IssuedAt, parsedToken.IssuedAt)
	}
}
