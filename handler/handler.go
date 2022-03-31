package handler

import "github.com/alexander-grube/bot/store"

type Handler struct {
	websiteStore store.WebsiteStore
}

func NewHandler(websiteStore store.WebsiteStore) *Handler {
	return &Handler{websiteStore: websiteStore}
}