package router

import (
	"fmt"

	"github.com/SergioNEOM/FOAM/auth"
	"github.com/SergioNEOM/FOAM/database"

	"github.com/gin-gonic/gin"
)

// StartRouter - привязать хэндлеры и запусить роутер
func StartRouter() {
	gin.SetMode(gin.ReleaseMode) //- убрать DEBUG mode
	r := gin.Default()
	//
	//
	// ?? router.Static("/assets", "../assets")
	//
	// not authorized group
	r.GET("/", rootHandler)
	r.GET("/refreshtoken", refreshToken)
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.File("./assets/favicon.ico")
	})
	// Статику (картинки, скрипти, стили) будем раздавать
	// по определенному роуту /static/:file
	// если файл не найден, вернется статус 404
	r.Static("/static/", "./assets")
	// конкретные файлы
	/*
		r.GET("/static/:file", func(c *gin.Context) {
			fname := "./assets/" + c.Param("file")
			fmt.Printf("requeststatic file: %v\n", fname)
			fi, err := os.Lstat(fname)
			if err == nil {
				if fi.Mode().IsRegular() {
					c.File(fname)
				}
			} else {
				c.Error(err)
				//log.Fatal(err)
			}
		})
	*/
	//---------- load templates
	//r.LoadHTMLFiles("./templates/meta.tmpl", "./templates/stufflist.tmpl", "./templates/stuffform.tmpl")
	r.LoadHTMLGlob("./templates/*.tmpl")
	//----------
	// stuff group (authorized):
	//	stuff := r.Group("/stuff", gin.BasicAuth(gin.Accounts{"admin": "admin"}))
	stuff := r.Group("/stuff", authFunc(auth.CheckToken)) // authFunc - обёртка для auth.CheckToken
	// показать список материалов
	//stuff.Use(auth)
	stuff.GET("", stuffListHandler)
	stuff.GET("/add", stuffAddFormHandler)
	stuff.POST("/add", stuffAddHandler)
	//-------------

	r.Run(":8888") // listen and serve on localhost:8888
}

//---

func rootHandler(c *gin.Context) {
	c.String(200, "main page for authorized users")
	// parse template for main page
}

//----
func refreshToken(c *gin.Context) {
	auth.SetNewToken(c)
}

// Заглушка: потом поставить вызов соотв.функции из модуля auth, которая:
// 1.проверит, есть ли в куке токен. Нет - редирект на форму логина
// 2.проверит подпись на токене. Нет - - редирект на форму логина
// 3.сверит имя в токене с введённым в логин-форме. Не совпадают - редирект на форму логина
// 4.проверит срок действия токена. Истёк - ПЕРЕВЫПУСК (с соотв.проверками refresh-token)
// 5.предоставит доступ к ресурсу (return или next.ServeHTTP()? ).

func authFunc(myhandler func(c *gin.Context) bool) func(*gin.Context) {
	return func(c *gin.Context) {
		fmt.Println("--- auth ---")
		if !myhandler(c) {
			fmt.Println("--- auth error ---")
			//todo: ??? как обработать
			// return nothing ?
			//c.Redirect(307, "/")
		}
	}
}

//показать список материалов
func stuffListHandler(c *gin.Context) {
	obj := *database.Dbase.GetStuffList()
	c.HTML(200, "stufflist.tmpl", obj) // !! Предварительно требует LoadHTMLFiles(...)
}

// stuffAddFormHandler отображает форму для добавления записи о материале
func stuffAddFormHandler(c *gin.Context) {
	c.HTML(200, "stuffform.tmpl", nil) // !! Предварительно требует LoadHTMLFiles(...)
}

// stuffAddHandler добавляет запись о материале
func stuffAddHandler(c *gin.Context) {
	// get values from form params
	sn, ok := c.GetPostForm("shortname")
	if !ok {
		// error ?
		// redirect to form ?
	}
	ds, ok := c.GetPostForm("description")
	if !ok {
		// error ?
		// redirect to form ?
	}
	// save values to DB
	/*	err := database.Dbase.AddStuff(&models.Stuff{ShortName: sn, Description: ds})
		if err != nil {
			// error ?
			// redirect to form ?

		}
	*/
	c.JSON(200, gin.H{"shortname": sn, "description": ds})
	// redirect to /stuff
	c.Redirect(302, "/stuff")
}
