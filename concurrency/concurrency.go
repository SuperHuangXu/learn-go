package main

import (
  "log"
  "sync"
  "time"
)

type WebsiteChecker func(string) bool

type result struct {
  string
  bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
  results := make(map[string]bool)
  resultChannel := make(chan result)

  for _, url := range urls {
    go func(u string) {
      resultChannel <- result{u, wc(u)}
    }(url)
  }

  for i := 0; i < len(urls); i++ {
    result := <-resultChannel
    results[result.string] = result.bool
  }
  return results
}

func DoSomething(id int, wg *sync.WaitGroup) {
  defer wg.Done()
  log.Printf("before do job:(%d) \n", id)
  time.Sleep(3 * time.Second)
  log.Printf("after do job:(%d) \n", id)
}

func main() {
  var wg sync.WaitGroup
  start := time.Now()
  wg.Add(3)
  go DoSomething(1, &wg)
  go DoSomething(2, &wg)
  go DoSomething(3, &wg)
  wg.Wait()
  log.Printf("finish all jobs, time: %.2f\n", time.Since(start).Seconds())
}
