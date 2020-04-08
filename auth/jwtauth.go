package auth

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const cookieName = "jwtoken"

type myClaims struct {
	Login string `json:"login"`
	jwt.StandardClaims
}

// CheckToken проверяет наличие JWT и сверяет имя пользователя
// jwtkey - ключ верификации токенов (ключ сервера)  --- вынести выше?

func CheckToken(c *gin.Context, jwtSecretKey []byte) gin.HandlerFunc {
	// 1.проверить, есть ли в куке токен. Нет - редирект на форму логина
	tokenStr, err := getToken(c)
	if err != nil {
		c.String(200, "No cookies with token") // c.Redirect(http.StatusMovedPermanently, "/signin")
		return c.Next()
	}
	//
	// 2.проверить подпись на токене. Нет - - редирект на форму логина
	token, err := jwt.ParseWithClaims(tokenStr, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})
	if !token.Valid || err != nil {
		//todo: write to log
		return c.String(200, "No token sign not valid") //c.Status(http.StatusUnauthorized)
	}
	// payload claims
	claims, ok := token.Claims.(*myClaims)
	if !ok || !token.Valid {
		return c.String(200, "Token not valid (claims not matched)")
	}
	// 3.сверит имя в токене с введённым в логин-форме. Не совпадают - редирект на форму логина
	// ?

	// 4.проверит срок действия токена. Истёк - ПЕРЕВЫПУСК (с соотв.проверками refresh-token)
	t := time.Now()
	if t.After(claims.StandardClaims.ExpiresAt) {
		return c.String(200, "Token expired! Go to refresh")
	}
	// 5.предоставит доступ к ресурсу (return или next.ServeHTTP()? ).

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
