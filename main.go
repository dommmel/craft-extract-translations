package main

import "fmt"
import "regexp"
import "io/ioutil"
import "path/filepath"
import "os"
import "flag"


func AppendIfMissing(slice []string, s string) []string {
  for _, ele := range slice {
    if ele == s {
      return slice
    }
  }
  return append(slice, s)
}

func scanFile(path string, f os.FileInfo, err error) error {
  if !f.IsDir() {
    data, _ := ioutil.ReadFile(path)
    input := string(data)
   
    html_quotes := `(\{\{\s*|\(\s*|\{\%.*?|:\s*)["'](.*?)["'].*?\|\s*?(t|translate)\s*(\(|\||\}\}|\)|\%\}|,)`

    re := regexp.MustCompile(html_quotes)
    matches := re.FindAllStringSubmatch(input, -1)
    for i := range matches {
      // fmt.Printf("%v\n\n", matches[i])
      keys = AppendIfMissing(keys,matches[i][2])
    }
  }
  return nil
}

var keys  []string 
func main() {
  flag.Parse()
  var root string
  if flag.NArg() == 1 {
    root = flag.Arg(0)
  } else if flag.NArg() == 0 {
    root = "."
  } else {
    fmt.Printf("Wrong number of arguments.\n")
    return 
  }
  path, _ := filepath.Abs(root)
  filepath.Walk(path, scanFile)
  fmt.Printf("<?php\n\nreturn array(\n");
  for i := range keys {
    fmt.Printf("\t\"%v\" => \"\",\n", keys[i])
  }
  fmt.Printf(");\n");
}
