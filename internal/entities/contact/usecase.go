package contact

import (
	"context"

	"github.com/booomch/demo-golang/internal/entities/contact/models"
)

type Usecase interface {
	CreateContact(ctx context.Context, userID, contactID int) (*models.Contact, error)
	DeleteContact(ctx context.Context, userID, contactID int) error
}
