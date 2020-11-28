package main

import (
  "encoding/json"
  "fmt"
)

type Student struct {
  Name string
  Age  int
}

type Per struct {
  Hello   string
  Student Student
}

func main() {
  arr := make([]Per, 0)

  arr = append(arr, Per{
    Hello: "hhh",
    Student: Student{
      Name: "xm",
      Age:  12,
    },
  })

  b, err := json.Marshal(arr)
  if err != nil {
    fmt.Errorf("%s", err.Error())
  }

  fmt.Printf("%s", string(b))
}
