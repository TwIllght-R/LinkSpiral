package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UrlMapping struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Alias       string             `bson:"alias"`
	OriginalURL string             `bson:"original_url"`
	CreatedAt   time.Time          `bson:"created_at"`
	ExpiryDate  *time.Time         `bson:"expiry_date,omitempty"`
}
