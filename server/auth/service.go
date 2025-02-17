package auth

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type Service interface {
	GenereteToken(userId int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenereteToken(userId int) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using default environment variables")
	}
	secretKey := os.Getenv("SECRETKEY_JWT")
	claim := jwt.MapClaims{}
	claim["user_id"] = userId
	claim["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Println("Generate token failed")
		return signedToken, err
	}
	return signedToken, nil

}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using default environment variables")
	}
	secretKey := os.Getenv("SECRETKEY_JWT")
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
