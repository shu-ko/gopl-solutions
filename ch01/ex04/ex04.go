package main

import (
  "bufio"
  "fmt"
  "os"
)

func main () {
  counts := make(map[string]int)
  containsFilenames := make(map[string][]string)
  files := os.Args[1:]

  for _, arg := range files {
    f, err := os.Open(arg)
    if err != nil {
      fmt.Fprintf(os.Stderr, "%v\n", err)
      continue
    }
    countLinesAndStoreFilename(f, counts, containsFilenames, arg)
    f.Close()
  }

  for line, n := range counts {
    if n > 1 {
      fmt.Printf("%d\t%s\t", n, line)

      fmt.Printf("file name: ")
      for _, arg := range containsFilenames[line] {
        fmt.Printf("%s\t", arg)
      }
      fmt.Printf("\n")
    }
  }
}

func countLinesAndStoreFilename(f *os.File, counts map[string]int, containsFilenames map[string][]string, filename string) {
  input := bufio.NewScanner(f)
  for input.Scan() {
    line := input.Text()
    counts[line]++

    // avoid overlap element
    isContain := false
    for _, arg := range containsFilenames[line] {
      if filename == arg {
        isContain = true
        break
      }
    }
    if isContain == false {
        containsFilenames[line] = append(containsFilenames[line], filename)
    }
  }
}
