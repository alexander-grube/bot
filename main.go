package main

import (
	"fmt"
	"os"

	"github.com/alexander-grube/bot/db"
	"github.com/alexander-grube/bot/handler"
	"github.com/alexander-grube/bot/store"
	"github.com/gin-gonic/gin"
)

var (
	PORT string = ":" + os.Getenv("PORT")
)

func main() {
	d := db.New()
	fmt.Println("Database Connected")
	db.AutoMigrate(d)
	fmt.Println("Database Migrated")

	r := gin.Default()

	ws := store.NewWebsiteStore(d)

	h := handler.NewHandler(*ws)

	h.RegisterRoutes(r)

	
	r.Run(PORT)
}