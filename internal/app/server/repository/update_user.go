package repository

import (
	"context"
	"github.com/eluceon/quizwithmebot/internal/app/server/models"
)

func (r *repository) UpdateUser(ctx context.Context, user models.User) (err error) {

	const query = `
		UPDATE users
		SET	username = $2,
			correct_answers = $3,
			total_answers = $4,
			is_passing = $5
		WHERE user_id = $1;
	`

	cmd, err := r.pool.Exec(ctx, query,
		user.ID,
		user.Username,
		user.CorrectAnswers,
		user.TotalAnswers,
		user.IsPassing,
	)
	if cmd.RowsAffected() == 0 {
		err = ErrNotFound
		return
	}

	return
}
