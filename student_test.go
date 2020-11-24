package main

import (
  "testing"
)

func TestStudent(t *testing.T) {
  assertCorrectMessage := func(t *testing.T, got, want string) {
    t.Helper()
    if got != want {
      t.Errorf("got '%q' want '%q'", got, want)
    }
  }

  t.Run("函数参数使用指针与否", func(t *testing.T) {
    stu := Student{"xm", 17}
    if stu.Age != 17 {
      t.Errorf("error")
    }
  })

  t.Run("success", func(t *testing.T) {
    stu := Student{"xh", 12}
    if stu.Age > 18 {
      t.Errorf("> 18")
    }
  })
  t.Run("name", func(t *testing.T) {
    stu := Student{
      "xm",
      18,
    }
    name := stu.Name
    want := "xm"
    assertCorrectMessage(t, name, want)
  })
}
