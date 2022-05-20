package repository

import (
	"context"
	"errors"
	"github.com/eluceon/quizwithmebot/internal/app/server/models"
	"github.com/jackc/pgx/v4"
)

func (r *repository) GetTopUsers(ctx context.Context) (users []models.User, err error) {
	const query = `
		SELECT user_id,
			username,
			correct_answers,
			total_answers,
			registered_at
		FROM users
		ORDER BY total_answers DESC
		LIMIT 5
	`

	rows, err := r.pool.Query(ctx, query)
	if errors.Is(err, pgx.ErrNoRows) {
		err = ErrNotFound
		return
	} else if err != nil {
		return
	}

	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Username, &user.CorrectAnswers, &user.TotalAnswers, &user.RegisteredAt)
		if err != nil {
			return
		}
		users = append(users, user)
	}

	return
}
