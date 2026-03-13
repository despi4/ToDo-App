package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func JWTSecret() []byte {
	_ = godotenv.Load()

	jwtSecret := os.Getenv("JWT_ACCESS_SECRET")

	return []byte(jwtSecret)
}

func genrateToken(UserID uuid.UUID, secretKey []byte) (string, error) {
	claims := jwt.MapClaims{
		"user_id": UserID,
		"exp":     time.Now().Add(time.Second * 15).Unix(),
	}

	fmt.Println("Claims:", claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	fmt.Println("Token", token)

	strToken, err := token.SignedString(secretKey)

	return strToken, err
}

func parseToken(tokenString string, secretKey []byte) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })
}

func main() {
	userID, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}

	secret := JWTSecret()

	token, err := genrateToken(userID, secret)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("String Token", token)

	t, err := parseToken(token, secret)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(t.Valid)
}
