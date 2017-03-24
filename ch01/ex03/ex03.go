package main

import (
  "fmt"
  "os"
  "strings"
)

func PrintArgsUsingIteration(args []string) {
  s, sep := "", ""
  for _, arg := range args {
    s += sep + arg
    sep = " "
  }
  fmt.Println(s)
}

func PrintArgsUsingJoin(args []string) {
  fmt.Println(strings.Join(args, " "))
}

func main () {
  PrintArgsUsingIteration(os.Args[1:])
  PrintArgsUsingJoin(os.Args[1:])
}
