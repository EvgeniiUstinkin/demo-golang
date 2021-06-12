package service

import (
	"context"
	"net/http"

	"github.com/booomch/demo-golang/internal/entities/user/models/response"
	"github.com/booomch/demo-golang/pkg/codes"
	"github.com/booomch/demo-golang/pkg/httperr"
)

func (svc *usecase) GetCurrentUser(ctx context.Context) (*response.User, error) {
	userUUID := ""
	if ctx.Value("CURRENT_USER_UUID") != nil {
		userUUID = ctx.Value("CURRENT_USER_UUID").(string)
	}
	if userUUID == "" {
		return nil, httperr.New(codes.Omit, http.StatusUnauthorized, "failed to retrieve uid from request context")
	}
	user, err := svc.repo.GetUserByUUID(ctx, userUUID)
	if err != nil {
		return nil, err
	}
	if user.DeletedAt != nil {
		return nil, httperr.New(codes.Omit, http.StatusForbidden, "user is deleted")
	}

	return user, nil
}
