package main

import (
	"fmt"

	_ "github.com/SergioNEOM/FOAM/config"
	"github.com/SergioNEOM/FOAM/router"
)

func main() {

	//
	fmt.Println("FOAM: listen and serve on: http://localhost:8888")

	router.StartRouter()

}
