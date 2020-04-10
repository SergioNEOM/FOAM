package auth

import (
	"errors"
	"fmt"
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
	tokenStr, err := c.Cookie(cookieName)
	if err != nil {
		//	если нет куки или токена -
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
		c.Error(err)
		c.String(403, "Token signature isn't valid")
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
	legal := claims.StandardClaims.VerifyExpiresAt(time.Now().Unix(), true)
	if !legal {
		c.String(401, "Token expired! Go to refresh")
		c.AbortWithStatus(401) //todo: redirect !!!!!!
		return false
	}
	// 5.предоставит доступ к ресурсу (return или next.ServeHTTP()? ).
	return true
}

// NewToken create new JSON Web token
//todo: если хранить несколько параметров в токене, предусмотреть их передачу вместо логина
func NewToken(login string, expire int64) string {
	// Устанавливаем набор параметров для токена
	claims := &myClaims{}
	claims.Login = login
	claims.StandardClaims.ExpiresAt = expire
	claims.StandardClaims.Issuer = "FOAM"
	// Создаем новый токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(secretFOAMKey))

	// Отдаем токен клиенту
	//	w.Write([]byte(tokenString))
	if err != nil {
		//todo: в журнал надо
		fmt.Printf("Func NewToken ERROR: %v\n", err)
		return ""
	}
	return tokenStr
}

func SetNewToken(c *gin.Context) {
	//todo: заменить литерал "admin" на реальное значение из формы login-form
	//todo: domain from config
	//todo: expired: time.Now().Add(time.Minute * 15).Unix()
	//exp := time.Now().Add(time.Hour).Unix()
	exp := time.Now().Add(time.Minute * 5).Unix()
	tokenStr := NewToken("admin", exp)
	if tokenStr != "" {
		c.SetCookie(cookieName, tokenStr, 3600 /*todo: ????*/, "/", "0.0.0.0", false, true)
	}
}
