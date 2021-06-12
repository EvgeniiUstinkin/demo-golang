package repositories

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/booomch/demo-golang/internal/utils"
	"github.com/sirupsen/logrus"
)

func (s *repository) CheckUserDeleted(ctx context.Context, userUUID string) (bool, error) {
	res := 0
	q := squirrel.Select(
		"id",
	).From("users").Where("firebase_id = ? and deleted_at IS NOT NULL", userUUID).
		PlaceholderFormat(squirrel.Question).RunWith(s.MasterNode())
	row := q.QueryRowContext(ctx)
	err := row.Scan(
		&res,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return res > 0, nil
		}
		q1, p, _ := q.ToSql()
		logrus.Error(utils.SqlErrLogMsg(err, q1, p))
		return res > 0, err
	}
	return res > 0, nil
}
