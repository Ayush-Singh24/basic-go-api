package auth

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Ayush-Singh24/basic-go-api/config"
	"github.com/Ayush-Singh24/basic-go-api/types"
	"github.com/golang-jwt/jwt"
)

func CreateJWT(secret []byte, userId int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    strconv.Itoa(userId),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.Envs.JWTSecret), nil
	})
}

func AlreadyLoggedIn(h types.UserStore, r *http.Request) (*types.User, bool) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return nil, false
	}
	token, err := ValidateJWT(cookie.Value)

	if err != nil {
		log.Printf("failed to validate token: %v\n", token)
		return nil, false
	}

	if !token.Valid {
		log.Println("invalid token")
		return nil, false
	}

	claims := token.Claims.(jwt.MapClaims)
	str := claims["userId"].(string)

	userId, err := strconv.Atoi(str)

	if err != nil {
		log.Println("str is", str)
		return nil, false
	}

	user, err := h.GetUserById(userId)

	if err != nil {
		log.Println("user not found")
		return nil, false
	}

	return user, true
}
