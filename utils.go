package main

import (
  "math/rand"
)

var randAlphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStub(length int) string {
  stub := make([]rune, length)
  for i := range stub {
    stub[i] = randAlphabet[rand.Intn(len(randAlphabet))]
  }
  return string(stub)
}
