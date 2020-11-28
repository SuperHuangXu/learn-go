package string

import (
  "testing"
)

func assertEquals(t *testing.T, want string, got string) {
  t.Helper()
  if got != want {
    t.Errorf("want: %s, but got: %s", want, got)
  }
}

func TestJoin(t *testing.T) {
  t.Run("sep 为空", func(t *testing.T) {
    strArr := []string{"h", "e", "l", "l", "o"}
    got := StrJoin(strArr, "")
    want := "hello"
    assertEquals(t, want, got)
  })
  t.Run("sep 为 ,", func(t *testing.T) {
    strArr := []string{"h", "e", "l", "l", "o"}
    got := StrJoin(strArr, ",")
    want := "h,e,l,l,o"
    assertEquals(t, want, got)
  })
}
