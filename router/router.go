package router

//import "github.com/gin-gonic/gin"

func StartRouter() {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode) //- убрать DEBUG mode
	//
	//
	// ?? router.Static("/assets", "../assets")
	//
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.String(200,"<h1>Hello, </h1><br><h2><world</h2>"})
	})

	r.Run(":8888") // listen and serve on localhost:8888

}
