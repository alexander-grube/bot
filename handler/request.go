package handler

import (
	"log"

	"github.com/alexander-grube/bot/model"
	"github.com/gin-gonic/gin"
)

type WebsiteUpRequest struct {
	Name string `json:"name" binding:"required"`
}

func (r *WebsiteUpRequest) bind(c *gin.Context, w *model.Website) error {
	if err := c.ShouldBindJSON(&w); err != nil {
		return err
	}
	log.Println("WebsiteUpRequest:", r.Name)
	w.Name = r.Name

	return nil
}
