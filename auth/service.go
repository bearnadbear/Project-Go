package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userID int) (string, error)
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
