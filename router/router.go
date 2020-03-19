package router

import (
	"github.com/SergioNEOM/FOAM/database"
	"github.com/SergioNEOM/FOAM/models"

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
	r.GET("/", rootHandler)
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.File("./assets/favicon.ico")
	})
	//---------- load templates
	r.LoadHTMLFiles("./templates/stufflist.tmpl", "./templates/stuffform.tmpl")
	//----------
	// stuff group:
	stuff := r.Group("/stuff")
	// показать список материалов
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

//показать список материалов
func stuffListHandler(c *gin.Context) {

	obj := *database.Dbase.GetStuffList()
	c.HTML(200, "stufflist.tmpl", obj) // !! Предварительно требует LoadHTMLFiles(...)

	/*	tmpl := template.Must(template.ParseFiles("./templates/stufflist.tmpl"))

		err := tmpl.Execute(c.Writer, obj)
		if err != nil {
			log.Fatalf("template execution: %s", err)
		}
	*/
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
	err := database.Dbase.AddStuff(&models.Stuff{"ShortName": sn, "Description": ds})
	if err != nil {
		// error ?
		// redirect to form ?

	}
	//c.JSON(200, gin.H{"shortname": sn, "description": ds})
	// redirect to /stuff
	c.Redirect(302, "/stuff")
}
