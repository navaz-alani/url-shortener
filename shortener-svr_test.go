package main

import (
	"context"
	"log"
	"testing"

	"github.com/navaz-alani/url-shortener/pb/go/pb"
	"google.golang.org/grpc"
)

func TestMain(t *testing.T) {
  conn, err := grpc.Dial("0.0.0.0:10000", grpc.WithInsecure())
  if err != nil {
    log.Fatalln(err)
  }
  defer conn.Close()

  client := pb.NewShortenerClient(conn)
  if sh, err := client.Shorten(context.TODO(), &pb.ShortenReq{ Url: "navaz.me" });
  err != nil {
    log.Println(err);
  } else {
    log.Println("got short: ", sh);
    if url, err := client.UnShorten(context.TODO(), &pb.Short{ Stub: sh.Stub });
    err != nil {
      log.Println(err);
    } else {
      log.Println("mapped to: ", url);
    }
  }
}
