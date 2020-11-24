package iteration

import "testing"

func TestRepeat(t *testing.T) {
  res := Repeat("a", 5)
  expected := "aaaaa"
  if res != expected {
    t.Errorf("res: %s, expected: %s", res, expected)
  }
}

func BenchmarkRepeat(b *testing.B) {
  for i := 0; i < b.N; i++ {
    Repeat("a", 5)
  }
}
