package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

const cookieName = "jwtoken"

// CheckToken проверяет наличие JWT и сверяет имя пользователя
// jwtkey - ключ верификации токенов (ключ сервера)  --- вынести выше?
func CheckToken(c *gin.Context, jwtKey []byte) gin.HandlerFunc {
	tokenStr, err := getToken(c)
	if err != nil {
		return c.Redirect(http.StatusMovedPermanently, "/signin")
	}
	//
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaimsExample{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !token.Valid || err != nil {
		//todo: write to log
		return c.Status(http.StatusUnauthorized) //error.New("Authentification erro. JWT token expired or invalid.")
	}
	// сверить имя польз
	//claims := token.Claims.(*CustomClaimsExample)
	//
	//return func(c *gin.Context) {
	//	c.Next()
	//}
}

// GetToken получить токен из cookie
func getToken(c *gin.Context) (string, error) {
	co, err := c.Cookie(cookieName)
	if err != nil {
		//	если нет куки или токена - возвратить пустую строку
		return "", err
	}
	return co, nil
}
