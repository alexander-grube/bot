package cron

import (
	"log"
	"net/http"
	"time"

	"github.com/alexander-grube/bot/model"
	"github.com/alexander-grube/bot/store"
)

type Cron struct {
	store store.WebsiteStore
}

func NewWebsiteCron(websiteStore store.WebsiteStore) *Cron {
	return &Cron{store: websiteStore}
}

func (c *Cron) CheckIfWebsitesAreUp() {
	websites, err := c.store.FindAll()
	if err != nil {
		return
	}

	for _, w := range websites {
		uptime := &model.Uptime{}
		log.Printf("Checking if %s is up", w.Name)
		start := time.Now()
		_, err := http.Get(w.Name)
		if err != nil {
			uptime.Up = false
			uptime.Ping = "down"
			w.Uptime = append(w.Uptime, *uptime)
			log.Printf("%s is down\n", w.Name)
			c.store.Update(w)
			continue
		}
		uptime.Up = true
		uptime.Ping = time.Since(start).String()
		w.Uptime = append(w.Uptime, *uptime)
		log.Printf("%s is up in %s\n", w.Name, time.Since(start))
		c.store.Update(w)
	}
}
