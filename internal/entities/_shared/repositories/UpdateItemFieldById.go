package repositories

import (
	"context"

	"github.com/booomch/demo-golang/internal/utils"

	"github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"
)

func UpdateItemFieldById(s BaseRepository, ctx context.Context, tableName string, id int, field string, value interface{}) error {
	{
		query := squirrel.Update(tableName).
			Set(field, value).
			Where("id=?", id).
			PlaceholderFormat(squirrel.Question).RunWith(s.MasterNode())

		_, err := query.ExecContext(ctx)
		if err != nil {
			q, p, _ := query.ToSql()
			logrus.Error(utils.SqlErrLogMsg(err, q, p))
			return err
		}
	}

	return nil
}
