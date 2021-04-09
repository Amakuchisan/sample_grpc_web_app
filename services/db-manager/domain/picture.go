package domain

import (
	"context"
)

// PictureID は単語にユニークに割り当てられる ID
type PictureID uint32

// Picture は単語を表す
type Picture struct {
	Image []byte    `db:"image"`
}

// PictureRepository はユーザーのリポジトリ
type PictureRepository interface {
	Find(ctx context.Context, num uint32) ([]Picture, error)
}

// GetPicture は、画像を指定された枚数取得する
func GetPicture(num uint32) func(ctx context.Context, r Repository) ([]Picture, error) {
	return func(ctx context.Context, r Repository) ([]Picture, error) {
		pictures, err := r.Picture().Find(ctx, num)
		return pictures, err
	}
}
