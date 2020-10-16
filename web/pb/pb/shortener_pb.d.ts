import * as jspb from 'google-protobuf'



export class Url extends jspb.Message {
  getUrl(): string;
  setUrl(value: string): Url;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Url.AsObject;
  static toObject(includeInstance: boolean, msg: Url): Url.AsObject;
  static serializeBinaryToWriter(message: Url, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Url;
  static deserializeBinaryFromReader(message: Url, reader: jspb.BinaryReader): Url;
}

export namespace Url {
  export type AsObject = {
    url: string,
  }
}

export class ShortenReq extends jspb.Message {
  getUrl(): string;
  setUrl(value: string): ShortenReq;

  getRequestedstub(): string;
  setRequestedstub(value: string): ShortenReq;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ShortenReq.AsObject;
  static toObject(includeInstance: boolean, msg: ShortenReq): ShortenReq.AsObject;
  static serializeBinaryToWriter(message: ShortenReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ShortenReq;
  static deserializeBinaryFromReader(message: ShortenReq, reader: jspb.BinaryReader): ShortenReq;
}

export namespace ShortenReq {
  export type AsObject = {
    url: string,
    requestedstub: string,
  }
}

export class Short extends jspb.Message {
  getStub(): string;
  setStub(value: string): Short;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Short.AsObject;
  static toObject(includeInstance: boolean, msg: Short): Short.AsObject;
  static serializeBinaryToWriter(message: Short, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Short;
  static deserializeBinaryFromReader(message: Short, reader: jspb.BinaryReader): Short;
}

export namespace Short {
  export type AsObject = {
    stub: string,
  }
}

