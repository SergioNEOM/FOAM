package main

import (
	"fmt"

	"github.com/SergioNEOM/FOAM/router"
)

func main() {

	//
	fmt.Println("FOAM: listen and serve on: http://localhost:8888")

	router.StartRouter()

}
