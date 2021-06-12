package service

import "context"

func (svc *usecase) GetUserIDByUUID(ctx context.Context, uuid string) (int, error) {
	return svc.repo.GetUserIDByUUID(ctx, uuid)
}
