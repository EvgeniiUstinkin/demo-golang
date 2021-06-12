package repositories

import (
	"context"

	"github.com/booomch/demo-golang/internal/entities/user/models/response"
)

func (s *repository) GetUserByUUID(ctx context.Context, uid string) (*response.User, error) {
	return s.GetUser(ctx, "firebase_id = ?", uid)
}
