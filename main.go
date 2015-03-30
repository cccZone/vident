package main

import (
       "fmt"
       "io/ioutil"
)

func readFile(fileName string) string {
     b, err := ioutil.ReadFile(fileName)
     if err != nil {
        panic(err)
     }
     
     return string(b)
}

func main() {
     file_contents := readFile("test.vi")
     fmt.Printf("%s\n", file_contents)
}
