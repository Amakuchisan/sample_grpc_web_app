package manager

import (
	"context"

	"github.com/Amakuchisan/sample_grpc_web_app/services/db-manager/domain"
	"github.com/Amakuchisan/sample_grpc_web_app/services/db-manager/repository"
)

// GetPictures はDBから単語の取得を行う
func (m *Manager) GetPictures(ctx context.Context, num uint32) ([][]byte, error) {
	repo := repository.NewRepository(m.db)
	pictures, err := domain.GetPicture(num)(ctx, repo)
	if err != nil {
		return nil, err
	}
	return structToInterface(pictures), nil
}

func structToInterface(pictures []domain.Picture) [][]byte {
	images := [][]byte{}
	for _, picture := range pictures {
		images = append(images, picture.Image)
	}
	return images
}
