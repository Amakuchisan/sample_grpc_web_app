package grpc

import (
	"context"
	"time"

	pb "example.com/user_name/sample_grpc_web_app/services/server/pb/picture"
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
