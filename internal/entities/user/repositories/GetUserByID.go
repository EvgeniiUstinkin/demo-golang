package repositories

import (
	"context"

	"github.com/booomch/demo-golang/internal/entities/user/models/response"
)

func (s *repository) GetUserByID(ctx context.Context, id int) (*response.User, error) {
	return s.GetUser(ctx, "id = ?", id)
}
