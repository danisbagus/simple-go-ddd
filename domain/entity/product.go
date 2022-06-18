package entity

import (
	"errors"
	"html"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	CategoryIDs []uint64           `json:"category_ids" bson:"category_ids"`
	Price       uint64             `json:"price" bson:"price"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
}

type ProductResponse struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	CategoryIDs []uint64 `json:"category_ids"`
	Price       uint64   `json:"price"`
}

func (p *Product) BeforeSave() {
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
}

func (p *Product) Prepare() {
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.CreatedAt = time.Now()
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return errors.New("name is required!")
	}
	if len(p.CategoryIDs) <= 0 {
		return errors.New("category id is required!")
	}
	if p.Price < 1000 {
		return errors.New("price must more than 1000")
	}

	return nil
}
