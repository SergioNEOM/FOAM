package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const cookieName = "jwtoken"
const secretFOAMKey = "*&~FOAMSecretKey~&*"

type myClaims struct {
	Login string `json:"login"`
	jwt.StandardClaims
}

// CheckToken проверяет наличие JWT и сверяет имя пользователя (middleware)
// jwtkey - ключ верификации токенов (ключ сервера)  --- вынести выше?
func CheckToken(c *gin.Context) bool {
	// 1.проверить, есть ли в куке токен. Нет - редирект на форму логина
	tokenStr, err := getToken(c)
	if err != nil {
		// c.Redirect(http.StatusMovedPermanently, "/signin")
		c.Error(errors.New("No cookies with token")) // записать в журнал
		c.String(403, "No cookies with token")       //todo: сделать шаблон с выводом ошибки
		c.AbortWithStatus(403)
		return false
	}
	//
	// 2.проверить подпись на токене. Нет - - редирект на форму логина
	ttt := &myClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, ttt, func(to *jwt.Token) (interface{}, error) {
		return []byte(secretFOAMKey), nil
	})
	if err != nil || !token.Valid { //проверка !token.Valid паниковала, когда стояла раньше err!=nil в случае ошибки
		//todo: write to log
		c.String(403, "Token sign isn't valid")
		c.AbortWithStatus(403) //c.Status(http.StatusUnauthorized)
		return false
	}
	// payload claims
	claims, ok := token.Claims.(*myClaims)
	if !ok || !token.Valid {
		c.Error(errors.New("Token not valid (claims not matched)"))
		c.String(403, "Token not valid (claims not matched)")
		c.AbortWithStatus(403)
		return false
	}
	// 3.сверит имя в токене с введённым в логин-форме. Не совпадают - редирект на форму логина
	// ?

	// 4.проверит срок действия токена. Истёк - ПЕРЕВЫПУСК (с соотв.проверками refresh-token)

	tcla := time.Unix(claims.StandardClaims.ExpiresAt, 0)
	t := time.Now()
	if t.After(tcla) {
		c.String(401, "Token expired! Go to refresh")
		c.AbortWithStatus(401) //todo: redirect !!!!!!
		return false
	}
	// 5.предоставит доступ к ресурсу (return или next.ServeHTTP()? ).
	return true
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
