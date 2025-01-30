package domain

import (
	"errors"
	"time"
)

var (
	ErrBannerNotFound = errors.New("banner not found")
)

type Banner struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type BannerRepository interface {
	GetByID(id int) (*Banner, error)
	Create(banner *Banner) error
	Update(banner *Banner) error
	Delete(id int) error
	GetAll() []Banner
}
