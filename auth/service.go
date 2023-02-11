package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct{}

func NewService() *jwtService {
	return &jwtService{}
}

var SK = []byte("Hohihuheho")

func (s *jwtService) GenerateToken(userID int) (string, error) {
	var claim = jwt.MapClaims{}

	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signToken, err := token.SignedString(SK)
	if err != nil {
		return signToken, err
	}

	return signToken, nil
}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	tokens, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(SK), nil
	})

	if err != nil {
		return tokens, err
	}

	return tokens, nil
}
