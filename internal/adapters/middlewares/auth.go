package middlewares

import (
	"back/pkg/jwt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware(ctx *gin.Context) {
	authHeader := strings.Split(ctx.GetHeader("Authorization"), "Bearer ")
	if len(authHeader) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Malformed token."})
		return
	}
	jwtToken := authHeader[1]

	claims := &jwtpackage.UserClaims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtpackage.JWT_SECRET), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
	if !tkn.Valid {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims.Destination != jwtpackage.DESTINATION_TOKEN {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userId", claims.Id)
	ctx.Next()
}
