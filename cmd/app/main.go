package main

import (
	"github.com/buskarion/todoapp-with-gin/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":3000")
}
