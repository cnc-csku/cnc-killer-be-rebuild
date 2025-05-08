package utils

import (
	"github.com/cnc-csku/cnc-killer-be-rebuild/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mitchellh/mapstructure"
)

type JWTToken struct {
	Email string `mapstructure:"email"`
	Role  string `mapstructure:"role"`
	Exp   int64  `mapstructure:"exp"`
}

func JWTDecode(tokenStr string, cfg *config.Config) (*JWTToken, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}
	var jwtToken JWTToken
	err = mapstructure.Decode(token.Claims.(jwt.MapClaims), &jwtToken)
	if err != nil {
		return nil, err
	}

	return &jwtToken, nil
}
