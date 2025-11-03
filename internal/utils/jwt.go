package utils

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

var (
	accessSecret  = []byte(os.Getenv("access_secret"))
	refreshSecret = []byte(os.Getenv("refresh_secret"))
	accessExp     = time.Minute * 15
	refreshExp    = time.Hour * 24 * 7
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

func GenerateToken(userID int) (*TokenPair, error) {
	accessClaims := &tokenClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessExp).Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   "access",
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessSigned, err := accessToken.SignedString(accessSecret)
	if err != nil {
		return nil, err
	}
	refreshClaims := &tokenClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(refreshExp).Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   "refresh",
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshSigned, err := refreshToken.SignedString(refreshSecret)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessSigned,
		RefreshToken: refreshSigned,
	}, nil
}

func ValidateAccessToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return accessSecret, nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*tokenClaims); ok && token.Valid {
		return claims.UserID, nil
	}

	return 0, errors.New("invalid access token")
}

func ValidateRefreshToken(refreshToken string) (int, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return refreshSecret, nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*tokenClaims); ok && token.Valid {
		return claims.UserID, nil
	}

	return 0, errors.New("invalid refresh token")
}

func SetAuthCookies(w http.ResponseWriter, tokens *TokenPair) {
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    tokens.AccessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   int(accessExp.Seconds()),
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    tokens.RefreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   int(refreshExp.Seconds()),
	})
}
