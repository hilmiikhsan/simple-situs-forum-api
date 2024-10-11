package posts

const (
	queryInsertPosts = `
		INSERT INTO posts
		(
			user_id,
			post_title,
			post_content,
			post_hashtags,
			create_at,
			created_by,
			updated_at,
			updated_by
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
)
