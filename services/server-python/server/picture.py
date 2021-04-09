import grpc

import pb.picture.picture_pb2 as picture_pb2
import pb.picture.picture_pb2_grpc as picture_pb2_grpc

class Picture(picture_pb2_grpc.PictureServicer):
    def GetPictures(self, request, context) -> picture_pb2.GetPicturesReply():
        with grpc.insecure_channel('db-manager:50051') as channel:
            stub = picture_pb2_grpc.PictureStub(channel)
            response = stub.GetPictures(picture_pb2.GetPicturesRequest(num=request.num))
        return response
