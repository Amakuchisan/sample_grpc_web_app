package grpc

import (
	"context"

	pb "github.com/Amakuchisan/sample_grpc_web_app/services/db-manager/pb/picture"
	"github.com/Amakuchisan/tsuginiyomu/services/manager/manager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetPictures は新規単語の検索を行い, 結果を返す
func (s *Server) GetPictures(ctx context.Context, in *pb.GetPicturesRequest) (*pb.GetPicturesReply, error) {
	pictures, err := s.manager.GetPictures(ctx, in.Num)
	if err != nil {
		if err == manager.ErrInvalidArgument {
			return nil, status.Error(codes.InvalidArgument, "invalid argument")
		}
		return nil, err
	}
	return &pb.GetPicturesReply{Pictures: pictures}, nil
}
