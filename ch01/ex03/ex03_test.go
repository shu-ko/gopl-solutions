package main

import (
  "testing"
)

var (
	args = []string{"abcdefg hijklmn opqrstu vwxyz"}
)

func BenchmarkPrintArgsUsingIteration(b *testing.B) {
  for i := 0; i < b.N; i++ {
    PrintArgsUsingIteration(args)
  }
}

func BenchmarkPrintArgsUsingJoin(b *testing.B) {
  for i := 0; i < b.N; i++ {
    PrintArgsUsingJoin(args)
  }
}
