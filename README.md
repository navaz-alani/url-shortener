# url-shortener

## About

Simple URL shortener app built with gRPC using Golang backend server +
Typescript (grpc-web). This is an opportunity for me to get started with
protocol buffers + gRPC as opposed to REST APIs that I am accustomed to. There
are 3 components, which tend to be common with grpc-web applications:

* `go-server` is the backend service
* `envoy` is the proxy
* `fe-client` is the frontend client
