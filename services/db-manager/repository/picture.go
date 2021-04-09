package repository

import (
	"context"
	"database/sql"

	"github.com/Amakuchisan/sample_grpc_web_app/services/db-manager/domain"
	"github.com/jmoiron/sqlx"
)

// PictureRepository は domain.PictureRepository に対するデータベースを使った実装
type PictureRepository struct {
	db DB
}

func newPictureRepository(db DB) *PictureRepository {
	return &PictureRepository{db}
}

// Find はリポジトリから指定された個数、ランダムに画像を取得する
func (r *PictureRepository) Find(ctx context.Context, num uint32) ([]domain.Picture, error) {
	pictures := []domain.Picture{}
	err := sqlx.SelectContext(
		ctx,
		r.db,
		&pictures,
		`
		SELECT image FROM picture ORDER BY RAND() LIMIT ?
		`,
		num,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return pictures, nil
}
