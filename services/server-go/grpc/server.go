package grpc

import (
	pb "example.com/user_name/sample/services/server/pb/picture"
)

// Server は pb.PictureServerに対する実装
type Server struct {
	pb.UnimplementedPictureServer
}
