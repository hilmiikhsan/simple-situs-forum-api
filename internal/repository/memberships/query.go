package memberships

const (
	queryGetUser = `
		SELECT
			id,
			email,
			username,
			password,
			created_at,
			created_by,
			updated_at,
			updated_by
		FROM users
		WHERE email = ? OR username = ? OR id = ?
	`

	queryInsertUser = `
		INSERT INTP users
		(
			email,
			username,
			password,
			created_at,
			created_by,
			updated_at,
			updated_by
		) VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	queryInsertRefreshTOken = `
		INSERT INTO refresh_tokens
		(
			user_id,
			refresh_token,
			expired_at,
			created_at,
			created_by,
			updated_at,
			updated_by
		) VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	queryGetRefreshToken = `
		SELECT
			id,
			user_id,
			refresh_token,
			expired_at,
			created_at,
			created_by,
			updated_at,
			updated_by
		FROM refresh_tokens
		WHERE user_id = ? AND expired_at >= NOW()
	`
)
