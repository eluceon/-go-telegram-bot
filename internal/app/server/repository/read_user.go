package repository

import (
	"context"
	"errors"
	"github.com/eluceon/quizwithmebot/internal/app/server/models"
	"github.com/jackc/pgx/v4"
)

func (r *repository) ReadUser(ctx context.Context, ID int64) (user models.User, err error) {
	const query = `
		SELECT user_id,
			username,
			correct_answers,
			total_answers,
			is_passing,
			registered_at
		FROM users
		WHERE user_id = $1;
	`

	err = r.pool.QueryRow(ctx, query, ID).Scan(
		&user.ID,
		&user.Username,
		&user.CorrectAnswers,
		&user.TotalAnswers,
		&user.IsPassing,
		&user.RegisteredAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		err = ErrNotFound
		return
	}

	return
}
