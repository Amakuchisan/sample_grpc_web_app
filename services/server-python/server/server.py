from concurrent import futures
import logging

import grpc
import pb.picture.picture_pb2_grpc as picture_pb2_grpc
from server.picture import Picture

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    picture_pb2_grpc.add_PictureServicer_to_server(Picture(), server)
    server.add_insecure_port('[::]:50050')
    server.start()
    server.wait_for_termination()
