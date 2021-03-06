package handler

import (
	"github.com/alexander-grube/bot/model"
	"github.com/gin-gonic/gin"
)

type WebsiteUpRequest struct {
	Name string `json:"name"`
}

func (r *WebsiteUpRequest) bind(c *gin.Context, w *model.Website) error {
	if err := c.ShouldBindJSON(&w); err != nil {
		return err
	}

	return nil
}
