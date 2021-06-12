package repositories

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/booomch/demo-golang/internal/utils"
	"github.com/sirupsen/logrus"
)

func (s *repository) GetUserIDByUUID(ctx context.Context, uuid string) (int, error) {
	res := 0
	q := squirrel.Select(
		"id",
	).From("core.users").Where("firebase_id = ?", uuid).
		PlaceholderFormat(squirrel.Question).RunWith(s.MasterNode())
	row := q.QueryRowContext(ctx)
	err := row.Scan(
		&res,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, nil
		}
		q1, p, _ := q.ToSql()
		logrus.Error(utils.SqlErrLogMsg(err, q1, p))
		return res, err
	}
	return res, nil
}
