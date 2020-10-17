package main

import (
	"context"
	"fmt"
	"log"

	"github.com/navaz-alani/url-shortener/pb/go/pb"
)

const StubLen = 6;

type shortenerSvr struct {
  	pb.UnimplementedShortenerServer
    stubs map[string]string
    urls  map[string]string
}

func newServer() *shortenerSvr {
	return &shortenerSvr{
    stubs: make(map[string]string),
    urls:  make(map[string]string),
  }
}


func (s *shortenerSvr) Shorten(c context.Context, u *pb.ShortenReq) (res *pb.Short, err error) {
  res= &pb.Short{}
  if _, ok := s.stubs[u.RequestedStub]; u.RequestedStub != "" && !ok {
    s.stubs[u.RequestedStub] = u.Url
    res.Stub = u.RequestedStub
    return res, nil
  }
  // check if url already has an assigned stub
  for url, stub := range s.urls {
    if (url == u.Url) {
      res.Stub = stub
      return res, nil
    }
  }
  // otherwise, generate new stub and assign to url
  stub := randStub(StubLen)
  s.stubs[stub] = u.Url
  s.urls[u.Url] = stub
  res.Stub = stub
  return res, nil
}

func (s *shortenerSvr) UnShorten(c context.Context, u *pb.Short) (res *pb.Url, err error) {
  res = &pb.Url{}
  if url, ok := s.stubs[u.Stub]; !ok {
    return res, fmt.Errorf("No such stub")
  } else {
    res.Url = url
    return res, nil
  }
}

