package authsvc

import (
	"errors"
	"log"
	"time"
	authdomain "todo-app/internal/domain/auth"
	userdomain "todo-app/internal/domain/user"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTService struct {
	accessSecret    []byte
	refreshSecret   []byte
	accesTokenTTL   time.Duration
	refreshTokenTTL time.Duration
}

// JWT payload data
// Claims types - Registered, public and private
type Claims struct {
	UserID uuid.UUID
	Role   userdomain.Role
	jwt.RegisteredClaims
}

func NewJWTService(accessSecret, refreshSecret []byte, accessTokenTTL, refreshTokenTTL time.Duration) *JWTService {
	return &JWTService{
		accessSecret:    accessSecret,
		refreshSecret:   refreshSecret,
		accesTokenTTL:   accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}

func (j *JWTService) generateAccessToken(userID uuid.UUID, role userdomain.Role) (string, error) {
	now := time.Now()

	log.Println("access now time", now)
	log.Println("access now add time", now.Add(j.accesTokenTTL))
	log.Println("jwt numeric date", jwt.NewNumericDate(now))

	accessClaims := Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(j.accesTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
			Subject:   userID.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)

	return token.SignedString(token)
}

func (j *JWTService) generateRefreshToken(userID uuid.UUID, role userdomain.Role) (string, error) {
	now := time.Now()

	refreshClaims := Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(j.refreshTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
			Subject:   userID.String(),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(j.refreshSecret)
	if err != nil {
		return "", err
	}

	return refreshTokenString, nil
}

func (j *JWTService) ValidateToken(tokenStr string, whichtoken authdomain.TokenType) (*Claims, error) {
	claims := &Claims{}
	var secret []byte

	if whichtoken == 0 {
		secret = j.accessSecret
	} else {
		secret = j.refreshSecret
	}

	token, err := jwt.ParseWithClaims(
		tokenStr,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, authdomain.ErrInvalidToken
			}

			return secret, nil
		},
	)

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, authdomain.ErrExpiredToken
		}

		return nil, authdomain.ErrInvalidToken
	}

	log.Println("Claims before:", claims)

	parsedClaims, ok := token.Claims.(*Claims)
	if !token.Valid || !ok {
		return nil, authdomain.ErrInvalidToken
	}

	log.Println("Claims after", parsedClaims)

	return claims, nil
}
