package main

import (
  "encoding/json"
  "errors"
  "fmt"
  "net/http"
  "strings"
  "sync"
)

type Game struct {
  Name   string
  Cookie string
  Path   string
}

type Res struct {
  Code    int    `json:"code"`
  Message string `json:"message"`
}

func main() {
  const baseUrl = "http://api.turinglabs.net/api/v1/jd"
  list := []Game{
    {Name: "bean", Cookie: "4npkonnsy7xi3c2kvo2wsv3qwbuv4eanmle45aa@mlrdw3aw26j3wcmr7v4fyf2jn5hqbmm4lx4wczq@qdvzduhc42hhfvipyh4i36nhjpuznkt4rmjer7q@olmijoxgmjutytsgchyelpenupiu4bzmu37uxdy@e7lhibzb3zek23muja3wjev65czhr5n3f2r4l7y@65sstlhmnnljvvo2wru6opitdr43aivc46rnszi@mlrdw3aw26j3weqhxnlh2vb44ak7z77purfdypq",
      Path: "bean"},
    {Name: "farm", Cookie: "0774c1079eae45b8a93baea7019f4a6e@e189aa8efb9247c9b49f78e409bdefa1@7a8c3647972e41c4a10ecffb45ae53c9@56a7ad2b7ae542248fb60c3e02202a13@d808a122e97c48b6b89e6de40e10507a@eef238b81bfc4eb0aec485e01f118b5e@2f04a34b05724cfdbb271c0d71a5c462", Path: "farm"},
    {Name: "pet", Cookie: "MTE1NDAxNzcwMDAwMDAwMzYxNjM3MzU=@MTAxODcxOTI2NTAwMDAwMDAyMjE0OTMzMQ==@MTAxODc2NTEzNDAwMDAwMDAzMzU4NDIyMQ==@MTE1NDUwMTI0MDAwMDAwMDM4NDgxNDY1", Path: "pet"},
    {Name: "ddfactory", Cookie: "P04z54XCjVWnYaS5jYPD2P93HpCnGtMt2YW6GHpGg@P04z54XCjVWnYaS5m9cZ2f-3XRMlOuX3FNtqzs@P04z54XCjVWnYaS5j0KDWXw3HpJiAFVXAITNQ@P04z54XCjVWnYaS5m9cZ2b_235JkH9XNUFvNx8@P04z54XCjVWnYaS5m9cZ2X72H0YnM9_J6VBjbY@P04z54XCjVWnYaS5m9cZ2X92HRDkTB5StR3x6o@P04z54XCjVWnYaS5m9cZ2Wp2HpNnGdVzLAbyTQ", Path: "ddfactory"},
  }

  for i, game := range list {
    cookieArr := strings.Split(game.Cookie, "@")
    fmt.Printf("---%d - %s---\n", i, game.Name)
    var wg sync.WaitGroup
    for i2, code := range cookieArr {
      wg.Add(1)
      go func(i2 int, code string) {
        defer wg.Done()
        url := fmt.Sprintf("%s/%s/create/%s", baseUrl, game.Path, code)
        res, err := http.Get(url)
        if err != nil {
          panic(errors.New(code + "error"))
        }
        var r Res
        json.NewDecoder(res.Body).Decode(&r)
        fmt.Printf("%d - %s: %v\n", i2, url, r)
      }(i2, code)
    }
    wg.Wait()
  }
}
