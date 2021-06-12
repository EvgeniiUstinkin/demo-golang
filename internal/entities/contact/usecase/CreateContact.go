package service

import (
	"context"
	"net/http"

	"github.com/booomch/demo-golang/internal/pkg/codes"
	"github.com/booomch/demo-golang/internal/pkg/httperr"

	"github.com/booomch/demo-golang/internal/entities/contact/models"
	"github.com/booomch/demo-golang/internal/utils"
)

func (svc *usecase) CreateContact(ctx context.Context, userID, contactID int) (*models.Contact, error) {
	contactUserExists, err := svc.repo.ValidateUserIdExists(ctx, contactID)
	if err != nil {
		return nil, err
	}
	if contactUserExists {
		return nil, httperr.New(codes.Omit, http.StatusBadRequest, "contact already exists")
	}

	newUUID := utils.NewUUID()
	err = svc.repo.CreateContact(ctx, userID, contactID, newUUID)
	if err != nil {
		return nil, err
	}
	res, err := svc.repo.GetUserByContactUUID(ctx, newUUID)
	if err != nil {
		return nil, err
	}
	return res, nil
}
