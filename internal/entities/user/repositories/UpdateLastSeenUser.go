package repositories

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/booomch/demo-golang/internal/utils"
	"github.com/sirupsen/logrus"
)

func (s *repository) UpdateLastSeenUser(ctx context.Context, userUUID string, lastSeen time.Time) error {
	query := squirrel.Update("core.users").
		Set("last_seen", lastSeen).
		Where("firebase_id=?", userUUID).
		PlaceholderFormat(squirrel.Question).RunWith(s.MasterNode())

	_, err := query.ExecContext(ctx)
	if err != nil {
		q, p, _ := query.ToSql()
		logrus.Error(utils.SqlErrLogMsg(err, q, p))
		return err
	}

	return nil
}
