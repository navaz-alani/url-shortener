/**
 * @fileoverview gRPC-Web generated client stub for pb
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as pb_shortener_pb from '../pb/shortener_pb';


export class ShortenerClient {
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

  methodInfoShorten = new grpcWeb.AbstractClientBase.MethodInfo(
    pb_shortener_pb.Short,
    (request: pb_shortener_pb.ShortenReq) => {
      return request.serializeBinary();
    },
    pb_shortener_pb.Short.deserializeBinary
  );

  shorten(
    request: pb_shortener_pb.ShortenReq,
    metadata: grpcWeb.Metadata | null): Promise<pb_shortener_pb.Short>;

  shorten(
    request: pb_shortener_pb.ShortenReq,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pb_shortener_pb.Short) => void): grpcWeb.ClientReadableStream<pb_shortener_pb.Short>;

  shorten(
    request: pb_shortener_pb.ShortenReq,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pb_shortener_pb.Short) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pb.Shortener/Shorten',
        request,
        metadata || {},
        this.methodInfoShorten,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pb.Shortener/Shorten',
    request,
    metadata || {},
    this.methodInfoShorten);
  }

  methodInfoUnShorten = new grpcWeb.AbstractClientBase.MethodInfo(
    pb_shortener_pb.Url,
    (request: pb_shortener_pb.Short) => {
      return request.serializeBinary();
    },
    pb_shortener_pb.Url.deserializeBinary
  );

  unShorten(
    request: pb_shortener_pb.Short,
    metadata: grpcWeb.Metadata | null): Promise<pb_shortener_pb.Url>;

  unShorten(
    request: pb_shortener_pb.Short,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: pb_shortener_pb.Url) => void): grpcWeb.ClientReadableStream<pb_shortener_pb.Url>;

  unShorten(
    request: pb_shortener_pb.Short,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: pb_shortener_pb.Url) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/pb.Shortener/UnShorten',
        request,
        metadata || {},
        this.methodInfoUnShorten,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/pb.Shortener/UnShorten',
    request,
    metadata || {},
    this.methodInfoUnShorten);
  }

}

