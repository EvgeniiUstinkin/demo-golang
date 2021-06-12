package repositories

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/booomch/demo-golang/internal/entities/user/models/response"
	"github.com/booomch/demo-golang/internal/utils"
	"github.com/sirupsen/logrus"
)

func (s *repository) GetUser(ctx context.Context, prefict string, params ...interface{}) (*response.User, error) {
	res := response.User{}

	query := squirrel.Select(
		"id",
		"is_active",
		"firebase_id",
		"stripe_id",
		"first_name",
		"last_name",
		"phone_number",
		"country_code",
		"email",
		"username",
		"is_verified",
		"user_type",
		"profile_avatar",
		"last_seen",
		"created_at",
		"updated_at",
		"deleted_at",
	).From("users").Where(prefict, params...).
		PlaceholderFormat(squirrel.Question).RunWith(s.MasterNode())
	row := query.QueryRowContext(ctx)
	err := row.Scan(
		&res.ID,
		&res.IsActive,
		&res.FirebaseID,
		&res.StripeID,
		&res.FirstName,
		&res.LastName,
		&res.Phonenumber,
		&res.CountryCode,
		&res.Email,
		&res.Username,
		&res.IsVerified,
		&res.Type,
		&res.ProfileAvatar,
		&res.LastSeen,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		q, p, _ := query.ToSql()
		logrus.Error(utils.SqlErrLogMsg(err, q, p))
		return nil, err
	}

	return &res, err
}
