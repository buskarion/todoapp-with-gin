package main

import (
	"github.com/buskarion/todoapp-with-gin/db"
	"github.com/buskarion/todoapp-with-gin/router"
)

func main() {
	db.BuildDB()
	r := router.SetupRouter()
	r.Run(":3000")
}
