package repositories

import (
	"context"

	"github.com/booomch/demo-golang/internal/entities/user/models/response"
)

func (s *repository) GetUserByEmail(ctx context.Context, email string) (*response.User, error) {
	return s.GetUser(ctx, "email = ?", email)
}
