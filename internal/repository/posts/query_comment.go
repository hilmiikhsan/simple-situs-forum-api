package posts

const (
	queryInsertComment = `
		INSERT INTO comments
		(
			post_id,
			user_id,
			comment_content,
			created_at,
			created_by,
			updated_at,
			updated_by
		) VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	queryGetCommentByPostID = `
		SELECT
			c.id,
			c.user_id,
			c.comment_content,
			u.username
		FROM comments c
		JOIN users u ON c.user_id = u.id
		WHERE c.post_id = ?
	`
)
