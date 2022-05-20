package repository

import (
	"context"
	"github.com/eluceon/quizwithmebot/internal/app/server/models"
)

func (r *repository) CreateUser(ctx context.Context, user models.User) (err error) {

	const query = `
		INSERT INTO users (
			user_id,
			username,
			correct_answers,
			total_answers,
			registered_at
		) VALUES (
			$1, $2, $3, $4, now()
		)
	`

	_, err = r.pool.Exec(ctx, query,
		user.ID,
		user.Username,
		user.CorrectAnswers,
		user.TotalAnswers,
	)

	return
}
