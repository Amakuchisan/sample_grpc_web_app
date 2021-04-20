import * as jspb from 'google-protobuf'



export class GetPicturesRequest extends jspb.Message {
  getNum(): number;
  setNum(value: number): GetPicturesRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetPicturesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetPicturesRequest): GetPicturesRequest.AsObject;
  static serializeBinaryToWriter(message: GetPicturesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetPicturesRequest;
  static deserializeBinaryFromReader(message: GetPicturesRequest, reader: jspb.BinaryReader): GetPicturesRequest;
}

export namespace GetPicturesRequest {
  export type AsObject = {
    num: number,
  }
}

export class GetPicturesReply extends jspb.Message {
  getPicturesList(): Array<Uint8Array | string>;
  setPicturesList(value: Array<Uint8Array | string>): GetPicturesReply;
  clearPicturesList(): GetPicturesReply;
  addPictures(value: Uint8Array | string, index?: number): GetPicturesReply;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetPicturesReply.AsObject;
  static toObject(includeInstance: boolean, msg: GetPicturesReply): GetPicturesReply.AsObject;
  static serializeBinaryToWriter(message: GetPicturesReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetPicturesReply;
  static deserializeBinaryFromReader(message: GetPicturesReply, reader: jspb.BinaryReader): GetPicturesReply;
}

export namespace GetPicturesReply {
  export type AsObject = {
    picturesList: Array<Uint8Array | string>,
  }
}

export class StreamGetPicturesReply extends jspb.Message {
  getPicture(): Uint8Array | string;
  getPicture_asU8(): Uint8Array;
  getPicture_asB64(): string;
  setPicture(value: Uint8Array | string): StreamGetPicturesReply;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StreamGetPicturesReply.AsObject;
  static toObject(includeInstance: boolean, msg: StreamGetPicturesReply): StreamGetPicturesReply.AsObject;
  static serializeBinaryToWriter(message: StreamGetPicturesReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StreamGetPicturesReply;
  static deserializeBinaryFromReader(message: StreamGetPicturesReply, reader: jspb.BinaryReader): StreamGetPicturesReply;
}

export namespace StreamGetPicturesReply {
  export type AsObject = {
    picture: Uint8Array | string,
  }
}

