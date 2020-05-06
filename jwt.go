package hive_go

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"strings"
)

type JWTBackend struct {
	repository SecretsRepository
}

func InitJWTBackend(repository *Repository) *JWTBackend {
	return &JWTBackend{repository: repository}
}

type JWTUser struct {
	jwt.StandardClaims
	IsAdmin  bool      `json:"isAdmin"`
	Roles    []string  `json:"roles"`
	UserID   uuid.UUID `json:"userID"`
	SecretID uuid.UUID `json:"secretID"`
}

func (user JWTUser) GetIsAdmin() bool {
	return user.IsAdmin
}

func (user JWTUser) GetRoles() []string {
	return user.Roles
}

func (user JWTUser) GetUserID() uuid.UUID {
	return user.UserID
}

func (backend JWTBackend) DecodeAccessTokenWithoutValidation(_ context.Context, tokenValue string) (int, *JWTUser) {
	parser := jwt.Parser{
		SkipClaimsValidation: false,
	}

	token, _, err := parser.ParseUnverified(tokenValue, &JWTUser{})

	if err != nil {
		return IncorrectToken, nil
	}

	if claims, ok := token.Claims.(*JWTUser); ok {
		return Ok, claims
	} else {
		return IncorrectToken, nil
	}
}

func (backend JWTBackend) DecodeAccessToken(_ context.Context, tokenValue string, secret uuid.UUID) (int, *JWTUser) {

	token, err := jwt.ParseWithClaims(tokenValue, &JWTUser{}, func(token *jwt.Token) (interface{}, error) {
		bytes, _ := secret.MarshalBinary()
		return bytes, nil
	})

	var e *jwt.ValidationError

	if err != nil {
		if errors.As(err, &e) && strings.Contains(e.Error(), "expired") {
			return InvalidToken, nil
		} else {
			return IncorrectToken, nil
		}
	}

	if !token.Valid {
		return InvalidToken, nil
	} else if claims, ok := token.Claims.(*JWTUser); ok {
		return Ok, claims
	} else {
		return IncorrectToken, nil
	}
}

func (backend JWTBackend) GetUser(ctx context.Context, token string) (int, IAuthenticationBackendUser) {

	status, unverifiedPayload := backend.DecodeAccessTokenWithoutValidation(ctx, token)
	if status != Ok {
		return status, nil
	}

	secret := backend.repository.GetSecret(unverifiedPayload.SecretID)
	if secret == nil {
		return SecretNotFound, nil
	}

	status, payload := backend.DecodeAccessToken(ctx, token, secret.Value)
	if status != Ok {
		return status, nil
	}

	return Ok, payload
}
