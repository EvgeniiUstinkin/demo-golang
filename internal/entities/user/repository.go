package user

import (
	"context"
	"time"

	srequest "github.com/booomch/demo-golang/internal/entities/_shared/models/request"
	baserepo "github.com/booomch/demo-golang/internal/entities/_shared/repositories"
	"github.com/booomch/demo-golang/internal/entities/user/models/response"
)

type Repository interface {
	baserepo.BaseRepository

	GetUserByEmail(ctx context.Context, email string) (*response.User, error)
	GetUserByID(ctx context.Context, id int) (*response.User, error)
	GetUserByUUID(ctx context.Context, uuid string) (*response.User, error)
	GetUserIDByUUID(ctx context.Context, uuid string) (int, error)
	GetUser(ctx context.Context, prefict string, params ...interface{}) (*response.User, error)
	CheckUserDeleted(ctx context.Context, userUUID string) (bool, error)
	UpdateLastSeenUser(ctx context.Context, userUUID string, lastSeen time.Time) error

	SearchUsersCalcQuery(ctx context.Context, userID int, filters []srequest.CustomFilterItem) (string, baserepo.SortDefinitionFunc, []string, error)
	SearchUsersExecQuery(ctx context.Context, query string, params []string) ([]response.User, error)
}
