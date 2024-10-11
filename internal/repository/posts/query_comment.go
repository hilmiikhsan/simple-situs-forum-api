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
)
