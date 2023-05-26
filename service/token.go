package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/MrRytis/go-fiber-test/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"time"
)

const AccessTokenJwtExpDuration = 30 * time.Minute
const RefreshTokenExpDuration = 12 * time.Hour

func GenerateJWT(user model.User) (string, error) {
	claims := jwt.MapClaims{
		"email":   user.Email,
		"name":    user.Name,
		"surname": user.Surname,
		"uid":     user.Uid,
		"icon":    user.Icon,
		"exp":     time.Now().Add(AccessTokenJwtExpDuration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := []byte(os.Getenv("JWT_SECRET"))
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateRefreshToken(user model.User) (string, error) {
	claims := jwt.MapClaims{
		"uid": user.Uid,
		"exp": time.Now().Add(RefreshTokenExpDuration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := []byte(os.Getenv("JWT_SECRET"))
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

func IsBlacklisted(cache *redis.Client, token string) bool {
	_, err := cache.Get(context.Background(), getJWTCacheKey(token)).Result()

	if err == redis.Nil {
		return false
	}

	if err != nil {
		log.Fatal("JWT cache look up failed: ", err)
	}

	return true
}

func BlackListToken(cache *redis.Client, token string, expiresAt int64) {
	SetCache(cache, getJWTCacheKey(token), 1, time.Unix(expiresAt, 0).Sub(time.Now()))
}

func getJWTCacheKey(token string) string {
	hash := md5.New().Sum([]byte(token))

	return "TOKEN_" + string(hash[:])
}
