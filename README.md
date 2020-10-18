# [url-shortener](https://url.navaz.me)

## About

Simple URL shortener app built with gRPC using Golang backend server +
Typescript (grpc-web). This is an opportunity for me to get started with
protocol buffers + gRPC as opposed to REST APIs that I am accustomed to. There
are 3 components, which tend to be common with grpc-web applications:

* `go-server` is the backend service
* `envoy` is the proxy
* `fe-client` is the frontend client

Currently hosted at: https://url.navaz.me

## Build & Deploy

Using the command `docker-compose up -d`, all 3 services will initialise in the
background. The port 4000 is for the frontend and the port 4001 is for the
backend `grpc-proxy` container. There exists an nginx configuration file which
proxies requests to their respective services (`fe-client` and `envoy`
services).

