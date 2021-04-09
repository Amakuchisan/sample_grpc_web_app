package repository

import (
	"github.com/Amakuchisan/sample_grpc_web_app/services/db-manager/domain"
	"github.com/jmoiron/sqlx"
)

// DB はデータベースのインターフェース
type DB interface {
	sqlx.Execer
	sqlx.ExecerContext
	sqlx.Queryer
	sqlx.QueryerContext
}

// Repository は domain.Repository に対するデータベースを使った実装
type Repository struct {
	picture *PictureRepository
}

// NewRepository は Repository を作成する
func NewRepository(db DB) *Repository {
	return &Repository{
		picture: newPictureRepository(db),
	}
}

// Picture は画像に対するリポジトリを返す
func (r *Repository) Picture() domain.PictureRepository {
	return r.picture
}
