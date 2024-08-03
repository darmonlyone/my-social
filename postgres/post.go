package postgres

import (
	"context"
	"database/sql"

	social "github.com/darmonlyone/my-social"
	"github.com/darmonlyone/my-social/postgres/boilentity"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type postRepo struct {
	db *sql.DB
}

func NewPostRepo(db *sql.DB) social.PostRepo {
	return &postRepo{
		db: db,
	}
}

func (r *postRepo) Find(ctx context.Context, id string) (*social.Post, error) {
	post, err := boilentity.Posts(qm.Where("id=?", id)).One(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return &social.Post{
		ID:        post.ID,
		CreatedBy: post.CreatedBy,
		Title:     post.Title,
		Content:   post.Content.String,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}, nil
}

func (r *postRepo) FindAll(ctx context.Context) ([]*social.Post, error) {
	posts, err := boilentity.Posts().All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	var result []*social.Post
	for _, post := range posts {
		result = append(result, &social.Post{
			ID:        post.ID,
			CreatedBy: post.CreatedBy,
			Title:     post.Title,
			Content:   post.Content.String,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		})
	}
	return result, nil
}

func (r *postRepo) Store(ctx context.Context, post *social.Post) error {
	boilPost := &boilentity.Post{
		ID:        post.ID,
		CreatedBy: post.CreatedBy,
		Title:     post.Title,
		Content:   null.StringFrom(post.Content),
	}
	return boilPost.Insert(ctx, r.db, boil.Infer())
}

func (r *postRepo) Update(ctx context.Context, id social.PostID, title *string, content *string) error {
	postEntity, err := boilentity.FindPost(ctx, r.db, id)
	if err != nil {
		return err
	}
	if title != nil {
		postEntity.Title = *title
	}
	if content != nil {
		postEntity.Content = null.StringFrom(*content)
	}
	_, err = postEntity.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (r *postRepo) Delete(ctx context.Context, id string) error {
	post, err := boilentity.Posts(qm.Where("id=?", id)).One(ctx, r.db)
	if err != nil {
		return err
	}
	_, err = post.Delete(ctx, r.db)
	return err
}
