package main

import (
	"fmt"
	"os"
	"time"

	"github.com/alexander-grube/bot/cron"
	"github.com/alexander-grube/bot/db"
	"github.com/alexander-grube/bot/handler"
	"github.com/alexander-grube/bot/store"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
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

	s := gocron.NewScheduler(time.UTC)

	cron := cron.NewWebsiteCron(*ws)

	s.Every(10).Seconds().Do(func() {
		fmt.Println("Running Cron")
		cron.CheckIfWebsitesAreUp()
	})

	h.RegisterRoutes(r)

	
	r.Run(PORT)
}