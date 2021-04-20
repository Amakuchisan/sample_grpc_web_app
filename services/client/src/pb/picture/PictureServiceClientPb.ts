/**
 * @fileoverview gRPC-Web generated client stub for picture
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as picture_pb from './picture_pb';


export class PictureClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoGetPictures = new grpcWeb.AbstractClientBase.MethodInfo(
    picture_pb.GetPicturesReply,
    (request: picture_pb.GetPicturesRequest) => {
      return request.serializeBinary();
    },
    picture_pb.GetPicturesReply.deserializeBinary
  );

  getPictures(
    request: picture_pb.GetPicturesRequest,
    metadata: grpcWeb.Metadata | null): Promise<picture_pb.GetPicturesReply>;

  getPictures(
    request: picture_pb.GetPicturesRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: picture_pb.GetPicturesReply) => void): grpcWeb.ClientReadableStream<picture_pb.GetPicturesReply>;

  getPictures(
    request: picture_pb.GetPicturesRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: picture_pb.GetPicturesReply) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/picture.Picture/GetPictures',
        request,
        metadata || {},
        this.methodInfoGetPictures,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/picture.Picture/GetPictures',
    request,
    metadata || {},
    this.methodInfoGetPictures);
  }

  methodInfoStreamGetPictures = new grpcWeb.AbstractClientBase.MethodInfo(
    picture_pb.StreamGetPicturesReply,
    (request: picture_pb.GetPicturesRequest) => {
      return request.serializeBinary();
    },
    picture_pb.StreamGetPicturesReply.deserializeBinary
  );

  streamGetPictures(
    request: picture_pb.GetPicturesRequest,
    metadata?: grpcWeb.Metadata) {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/picture.Picture/StreamGetPictures',
      request,
      metadata || {},
      this.methodInfoStreamGetPictures);
  }

}

