package posts

import (
	"context"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/hilmiikhsan/situs-forum/internal/model/posts"
)

func (r *repository) CreatePost(ctx context.Context, model posts.PostModel) error {
	_, err := r.db.ExecContext(ctx, queryInsertPosts,
		model.UserID,
		model.PostTitle,
		model.PostContent,
		model.PostHashtags,
		model.CreatedAt,
		model.CreatedBy,
		model.UpdatedAt,
		model.UpdatedBy,
	)
	if err != nil {
		log.Error().Err(err).Msg("failed to insert post")
		return err
	}

	return nil
}

func (r *repository) GetAllPost(ctx context.Context, limit, offset int) (posts.GetAllPostResponse, error) {
	responses := posts.GetAllPostResponse{}

	rows, err := r.db.QueryContext(ctx, queryGetAllPost, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("failed to query post")
		return responses, err
	}
	defer rows.Close()

	data := make([]posts.Post, 0)

	for rows.Next() {
		var (
			model    posts.PostModel
			username string
		)

		err = rows.Scan(
			&model.ID,
			&model.UserID,
			&username,
			&model.PostTitle,
			&model.PostContent,
			&model.PostHashtags,
		)
		if err != nil {
			log.Error().Err(err).Msg("failed to scan post")
			return responses, err
		}

		data = append(data, posts.Post{
			ID:           model.ID,
			UserID:       model.UserID,
			Username:     username,
			PostTitle:    model.PostTitle,
			PostContent:  model.PostContent,
			PostHashtags: strings.Split(model.PostHashtags, ","),
		})
	}

	responses.Data = data
	responses.Pagination = posts.Pagination{
		Limit:  limit,
		Offset: offset,
	}

	return responses, nil
}

func (r *repository) GetPostByID(ctx context.Context, id int64) (*posts.Post, error) {
	var (
		model    posts.PostModel
		username string
		isLiked  bool
	)

	row := r.db.QueryRowContext(ctx, queryGetPostByID, id)

	err := row.Scan(
		&model.ID,
		&model.UserID,
		&username,
		&model.PostTitle,
		&model.PostContent,
		&model.PostHashtags,
		&isLiked,
	)
	if err != nil {
		log.Error().Err(err).Msg("failed to scan post")
		return nil, err
	}

	return &posts.Post{
		ID:           model.ID,
		UserID:       model.UserID,
		Username:     username,
		PostTitle:    model.PostTitle,
		PostContent:  model.PostContent,
		PostHashtags: strings.Split(model.PostHashtags, ","),
		IsLiked:      isLiked,
	}, nil
}
