package grpc

import (
	pb "example.com/user_name/sample_grpc_web_app/services/server/pb/picture"
)

// Server は pb.PictureServerに対する実装
type Server struct {
	pb.UnimplementedPictureServer
}
