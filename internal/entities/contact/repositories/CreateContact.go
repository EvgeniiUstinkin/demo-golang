package repositories

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/booomch/demo-golang/internal/utils"
	"github.com/sirupsen/logrus"
)

func (s *repository) CreateContact(ctx context.Context, userID, contactID int, uuid string) error {
	query := squirrel.Insert("users_contacts").
		Columns(
			"uuid",
			"user_id",
			"contact_id",
		).
		Values(
			uuid,
			userID,
			contactID,
		).
		PlaceholderFormat(squirrel.Question).RunWith(s.MasterNode())

	_, err := query.ExecContext(ctx)
	if err != nil {
		q, p, _ := query.ToSql()
		logrus.Error(utils.SqlErrLogMsg(err, q, p))
		return err
	}

	return nil
}
