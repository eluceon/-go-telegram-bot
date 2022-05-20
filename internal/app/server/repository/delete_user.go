package repository

import "context"

func (r *repository) DeleteUser(ctx context.Context, ID int64) (err error) {

	const query = `
		DELETE FROM users
		WHERE user_id = $1;
	`

	cmd, err := r.pool.Exec(ctx, query, ID)
	if cmd.RowsAffected() == 0 {
		err = ErrNotFound
		return
	}

	return
}
