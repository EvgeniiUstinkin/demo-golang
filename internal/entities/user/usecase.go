package user

import (
	"context"

	srequest "github.com/booomch/demo-golang/internal/entities/_shared/models/request"
	"github.com/booomch/demo-golang/internal/entities/user/models/response"
)

type Usecase interface {
	SearchUsers(ctx context.Context, req srequest.NewGridList) (*response.UserList, error)
	GetUserByUUID(ctx context.Context, uuid string) (*response.User, error)
	GetUserIDByUUID(ctx context.Context, uuid string) (int, error)

	PingServer(ctx context.Context) error
	UpdateLastSeenForCurrentUser(ctx context.Context) error
	GetCurrentUser(c context.Context) (*response.User, error)
}
