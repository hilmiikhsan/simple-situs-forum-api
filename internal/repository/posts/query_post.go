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

	queryGetAllPost = `
		SELECT
			p.id,
			p.user_id,
			u.username,
			p.post_title,
			p.post_content,
			p.post_hashtags
		FROM posts p
		JOIN users u ON p.user_id = u.id
		ORDER BY p.updated_at DESC
		LIMIT ? OFFSET ?
	`

	queryGetPostByID = `
		SELECT
			p.id,
			p.user_id,
			u.username,
			p.post_title,
			p.post_content,
			p.post_hashtags,
			uv.is_liked
		FROM posts p
		JOIN user_activities uv ON uv.post_id = p.id
		WHERE p.id = ?
	`
)
