package handler

import (
	"github.com/alexander-grube/bot/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateWebsiteEntry(c *gin.Context) {
	r := &WebsiteUpRequest{}
	w := model.Website{}

	w.Uptime = []model.Uptime{} 

	if err := r.bind(c, &w); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.websiteStore.Create(&w); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"message": "ok"})
}
