package store

import (
	"github.com/alexander-grube/bot/model"
	"gorm.io/gorm"
)

type WebsiteStore struct {
	db *gorm.DB
}

func NewWebsiteStore(db *gorm.DB) *WebsiteStore {
	return &WebsiteStore{db: db}
}

func (ws *WebsiteStore) Create(w *model.Website) error {
	return ws.db.Create(w).Error
}

func (ws *WebsiteStore) Find(name string) (*model.Website, error) {
	var w model.Website
	if err := ws.db.Where("name = ?", name).First(&w).Error; err != nil {
		return nil, err
	}
	return &w, nil
}

func (ws *WebsiteStore) FindAll() ([]*model.Website, error) {
	var websites []*model.Website
	if err := ws.db.Find(&ws).Error; err != nil {
		return nil, err
	}
	return websites, nil
}

func (ws *WebsiteStore) Update(w *model.Website) error {
	return ws.db.Save(w).Error
}

func (ws *WebsiteStore) Delete(w *model.Website) error {
	return ws.db.Delete(w).Error
}