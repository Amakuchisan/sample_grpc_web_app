package grpc

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	pb "example.com/user_name/sample/services/server/pb/picture"
	"google.golang.org/grpc"
)

// GetPictures は、filesディレクトリに含まれている画像を、指定された枚数、クライアントに返却する
func (s *Server) GetPictures(ctx context.Context, in *pb.GetPicturesRequest) (*pb.GetPicturesReply, error) {
	// selectedFilePaths, err := getRandomFiles(in.GetNum()) // 要素のシャッフル
	// var pictures [][]byte
	// if err != nil {
	// 	return &pb.GetPicturesReply{
	// 		Pictures: pictures,
	// 	}, err
	// }
	// for _, fileName := range selectedFilePaths {
	// 	file, err := os.Open(fileName)
	// 	if err != nil {
	// 		return &pb.GetPicturesReply{
	// 			Pictures: pictures,
	// 		}, err
	// 	}

	// 	data, err := ioutil.ReadAll(file)
	// 	if err != nil {
	// 		return &pb.GetPicturesReply{
	// 			Pictures: pictures,
	// 		}, err
	// 	}
	// 	pictures = append(pictures, data)
	// }

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

func getRandomFiles(num uint32) ([]string, error) {
	dir := "/services/server/files/"
	file, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	fileNum := uint32(len(file))       // ファイルの個数
	indexes := make([]uint32, fileNum) //
	for i := uint32(0); i < fileNum; i++ {
		indexes[i] = i
	}
	rand.Shuffle(len(indexes), func(i, j int) {
		indexes[i], indexes[j] = indexes[j], indexes[i]
	})

	indexLen := num
	if num > fileNum {
		indexLen = fileNum
	}

	var newFileNames []string
	for i := uint32(0); i < indexLen; i++ {
		newFileNames = append(newFileNames, dir+file[indexes[i]].Name())
	}
	return newFileNames, nil
}

func requestDBManager(num uint32) ([][]byte, error) {
	fmt.Println("requestDB")
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
	fmt.Println("get reply")
	fmt.Println(reply)
	return reply.GetPictures(), nil
}
