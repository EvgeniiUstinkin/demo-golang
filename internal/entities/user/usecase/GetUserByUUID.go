package service

import (
	"context"

	"github.com/booomch/demo-golang/internal/entities/user/models/response"
)

func (svc *usecase) GetUserByUUID(ctx context.Context, uuid string) (*response.User, error) {
	res, err := svc.repo.GetUserByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}
	return res, nil
}
