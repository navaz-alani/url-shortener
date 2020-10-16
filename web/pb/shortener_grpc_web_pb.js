/**
 * @fileoverview gRPC-Web generated client stub for pb
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.pb = require('./shortener_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.pb.ShortenerClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.pb.ShortenerPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pb.ShortenReq,
 *   !proto.pb.Short>}
 */
const methodDescriptor_Shortener_Shorten = new grpc.web.MethodDescriptor(
  '/pb.Shortener/Shorten',
  grpc.web.MethodType.UNARY,
  proto.pb.ShortenReq,
  proto.pb.Short,
  /**
   * @param {!proto.pb.ShortenReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pb.Short.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.pb.ShortenReq,
 *   !proto.pb.Short>}
 */
const methodInfo_Shortener_Shorten = new grpc.web.AbstractClientBase.MethodInfo(
  proto.pb.Short,
  /**
   * @param {!proto.pb.ShortenReq} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pb.Short.deserializeBinary
);


/**
 * @param {!proto.pb.ShortenReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.pb.Short)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pb.Short>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pb.ShortenerClient.prototype.shorten =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pb.Shortener/Shorten',
      request,
      metadata || {},
      methodDescriptor_Shortener_Shorten,
      callback);
};


/**
 * @param {!proto.pb.ShortenReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pb.Short>}
 *     Promise that resolves to the response
 */
proto.pb.ShortenerPromiseClient.prototype.shorten =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pb.Shortener/Shorten',
      request,
      metadata || {},
      methodDescriptor_Shortener_Shorten);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pb.Short,
 *   !proto.pb.Url>}
 */
const methodDescriptor_Shortener_UnShorten = new grpc.web.MethodDescriptor(
  '/pb.Shortener/UnShorten',
  grpc.web.MethodType.UNARY,
  proto.pb.Short,
  proto.pb.Url,
  /**
   * @param {!proto.pb.Short} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pb.Url.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.pb.Short,
 *   !proto.pb.Url>}
 */
const methodInfo_Shortener_UnShorten = new grpc.web.AbstractClientBase.MethodInfo(
  proto.pb.Url,
  /**
   * @param {!proto.pb.Short} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pb.Url.deserializeBinary
);


/**
 * @param {!proto.pb.Short} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.pb.Url)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pb.Url>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pb.ShortenerClient.prototype.unShorten =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pb.Shortener/UnShorten',
      request,
      metadata || {},
      methodDescriptor_Shortener_UnShorten,
      callback);
};


/**
 * @param {!proto.pb.Short} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pb.Url>}
 *     Promise that resolves to the response
 */
proto.pb.ShortenerPromiseClient.prototype.unShorten =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pb.Shortener/UnShorten',
      request,
      metadata || {},
      methodDescriptor_Shortener_UnShorten);
};


module.exports = proto.pb;

