package cron

import (
	"log"
	"net"

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
		log.Printf("Checking if %s is up", w.Name)
		_, err := net.Dial("tcp", w.Name+":80")
		if err != nil {
			w.Up = false
			c.store.Update(w)
			continue
		}
		w.Up = true
		c.store.Update(w)
	}
}
