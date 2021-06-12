package service

import (
	"context"
	"time"
)

func (svc *usecase) UpdateLastSeenForCurrentUser(ctx context.Context) error {
	userUUID := ""
	if ctx.Value("CURRENT_USER_UUID") != nil {
		userUUID = ctx.Value("CURRENT_USER_UUID").(string)
	}
	if userUUID == "" {
		return nil
	}
	lastSeen := time.Now().UTC()
	err := svc.repo.UpdateLastSeenUser(ctx, userUUID, lastSeen)
	if err != nil {
		return err
	}
	return nil
}
