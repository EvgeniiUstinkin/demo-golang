package controller

import (
	"context"

	srequest "github.com/booomch/demo-golang//internal/entities/_shared/models/request"
	"github.com/booomch/demo-golang/internal/entities/user"
)

func PingServer(uc interface{}, ctx context.Context, incomeRequest interface{}) (interface{}, error) {
	svc := uc.(user.Usecase)

	err := svc.PingServer(ctx)

	return nil, err
}

func GetUsers(uc interface{}, ctx context.Context, incomeRequest interface{}) (interface{}, error) {
	svc := uc.(user.Usecase)
	req := incomeRequest.(srequest.NewGridList)

	res, err := svc.SearchUsers(ctx, req)

	return res, err
}
