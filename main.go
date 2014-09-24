package main

import (
    "fmt"
    "io/ioutil"
    "github.com/docopt/docopt-go"
    "encoding/json"
    "strconv"
)

type Comment struct {
    Text string `json:"text"`
    Points string `json:"points"`
    IntPoints int `json:"int_points"`
    User string `json:"user"`
    Time string `json:"time"`
}

type Post struct {
    Title string `json:"title"`
    User string `json:"user"`
    Time string `json:"time"`
    Points string `json:"points"`
    IntPoints int `json:"int_points"`
    Id string `json:"id"`
    Comments []Comment `json:"comments"`
}

type File struct {
    Posts []Post `json:"posts"`
}

func main() {
      usage := `jq2mongo.

Usage:
  jq2mongo <file>

Options:
  -h --help     Show this screen.
  --version     Show version.
`  
      arguments, _ := docopt.Parse(usage, nil, true, "jq2mongo 0", false)

      content, err := ioutil.ReadFile(arguments["<file>"].(string))
      if err != nil {
        panic(err)
      }

      var f File
      err = json.Unmarshal(content, &f)

      if err != nil {
        panic(err)
      }

      var o []byte

      for p := range(f.Posts) {
        post := f.Posts[p]

        ip, err := strconv.Atoi(post.Points)
        if err == nil {
            post.IntPoints = ip
        }

        for c := range(post.Comments) {
            comment := post.Comments[c]
            ipc, err := strconv.Atoi(comment.Points)
            if err == nil {
                comment.IntPoints = ipc
            }
            post.Comments[c] = comment
        }

        j, err := json.Marshal(post)
        if err != nil {
            panic(err)
        }

        j = append(j, '\n')
        o = append(o, j...)

      }

      fmt.Printf("%s", o)
}