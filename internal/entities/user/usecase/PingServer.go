package service

import "context"

func (svc *usecase) PingServer(ctx context.Context) error {
	err := svc.UpdateLastSeenForCurrentUser(ctx)
	if err != nil {
		return err
	}
	return nil
}
