package grpc

import (
	"context"

	"github.com/Amakuchisan/sample_grpc_web_app/services/db-manager/manager"
	pb "github.com/Amakuchisan/sample_grpc_web_app/services/db-manager/pb/picture"
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

// StreamGetPictures は新規単語の検索を行い, 結果を返す
func (s *Server) StreamGetPictures(in *pb.GetPicturesRequest, stream pb.Picture_StreamGetPicturesServer) error {
	ctx := context.Background()
	pictures, err := s.manager.GetPictures(ctx, in.Num)
	if err != nil {
		if err == manager.ErrInvalidArgument {
			return status.Error(codes.InvalidArgument, "invalid argument")
		}
		return err
	}
	for _, picture := range pictures {
		if err := stream.Send(&pb.StreamGetPicturesReply{
			Picture: picture,
		}); err != nil {
			return err
		}
	}
	return nil
}
