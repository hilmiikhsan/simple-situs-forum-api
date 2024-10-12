package posts

const (
	queryGetUserActivity = `
		SELECT
			id,
			post_id,
			user_id,
			is_liked,
			created_at,
			created_by,
			updated_at,
			updated_by
		FROM user_activities
		WHERE post_id = ? AND user_id = ?
	`

	queryInsertUserActivity = `
		INSERT INTO user_activities
		(
			post_id,
			user_id,
			is_liked,
			created_at,
			created_by,
			updated_at,
			updated_by
		) VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	queryUpdateUserActivity = `
		UPDATE user_activities
		SET
			is_liked = ?,
			updated_at = ?,
			updated_by = ?
		WHERE post_id = ? AND user_id = ?
	`

	queryCountLikeByPostID = `
		SELECT COUNT(id) FROM user_activities
		WHERE post_id = ? AND is_liked = true
	`
)
