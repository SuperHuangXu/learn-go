package main

import (
  "errors"
  "fmt"
)

type Student struct {
  Name string
  Age  int
}

func AddAge(stu Student) int {
  return stu.Age + 100
}

func main() {
  age := 11
  newAge, err := he(age)
  if err != nil {
    fmt.Println(err.Error())
  } else {
    fmt.Printf("newAge: %d", newAge)
  }
}

func he(num int) (res int, err error) {
  if num > 18 {
    err = errors.New("不可大于18")
    return
  }
  res = num + 100
  return
}
