package grpc

import (
	"context"
	"io"
	"time"

	pb "example.com/user_name/sample_grpc_web_app/services/server/pb/picture"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// GetPictures は、filesディレクトリに含まれている画像を、指定された枚数、クライアントに返却する
func (s *Server) GetPictures(ctx context.Context, in *pb.GetPicturesRequest) (*pb.GetPicturesReply, error) {
	pictures, err := requestDBManager(in.GetNum())
	if err != nil {
		return &pb.GetPicturesReply{
			Pictures: nil,
		}, err
	}

	return &pb.GetPicturesReply{
		// Pictures: pictures,
		Pictures: pictures,
	}, nil
}

// StreamGetPictures is
func (s *Server) StreamGetPictures(in *pb.GetPicturesRequest, stream pb.Picture_StreamGetPicturesServer) error {
	return streamRequestDBManager(in.GetNum(), stream)
}

func requestDBManager(num uint32) ([][]byte, error) {
	address := "db-manager:50051"
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	client := pb.NewPictureClient(conn)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()
	getPicturesRequest := pb.GetPicturesRequest{
		Num: num,
	}
	reply, err := client.GetPictures(ctx, &getPicturesRequest)
	if err != nil {
		return nil, err
	}
	return reply.GetPictures(), nil
}

func streamRequestDBManager(num uint32, server pb.Picture_StreamGetPicturesServer) error {
	address := "db-manager:50051"
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	client := pb.NewPictureClient(conn)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()
	getPicturesRequest := pb.GetPicturesRequest{
		Num: num,
	}
	stream, err := client.StreamGetPictures(ctx, &getPicturesRequest)
	if err != nil {
		return errors.Wrap(err, "streamエラー")
	}
	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if err := server.Send(&pb.StreamGetPicturesReply{
			Picture: reply.GetPicture(),
		}); err != nil {
			return err
		}
	}
	return nil
}
