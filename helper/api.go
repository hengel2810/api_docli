package helper

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
)

func UserIdFromRequest(r *http.Request) string {
	user := r.Context().Value("user")
	claims, ok := user.(*jwt.Token).Claims.(jwt.MapClaims)
	if !ok || claims["sub"] == nil {
		return ""
	}
	userId := claims["sub"].(string)
	return userId
}