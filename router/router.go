package router

import (
	"github.com/gin-gonic/gin"
)

// StartRouter - привязать хэндлеры и запусить роутер
func StartRouter() {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode) //- убрать DEBUG mode
	//
	//
	// ?? router.Static("/assets", "../assets")
	//
	r.GET("/", rootHandler)
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.File("./assets/favicon.ico")
	})
	// показать список материалов
	r.GET("/stuff", stuffListHandler)

	r.Run(":8888") // listen and serve on localhost:8888
}

//---

func rootHandler(c *gin.Context) {
	c.String(200, "main page for authorized users")
	// parse template for main page
}

//показать список материалов
func stuffListHandler(c *gin.Context) {
	c.String(200, "stuff list")
	//
	// get stuff list from DB
	//
	// parse template for stuff list
	//
	// ExecuteTemplate with param <StuffList>

}
