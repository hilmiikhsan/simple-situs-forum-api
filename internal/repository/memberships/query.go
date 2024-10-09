package memberships

const (
	queryGetUserByEmailOrUsername = `
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
		WHERE email = ? OR username = ?
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
)
