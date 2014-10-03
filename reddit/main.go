package main

import (
    "fmt"
    "github.com/docopt/docopt-go"
    "github.com/jzelinskie/reddit"
)

func main() {
    usage := `reddit.

    Usage:
      reddit <username> <password>

    Options:
      -h --help     Show this screen.
      --version     Show version.`

    arguments, _ := docopt.Parse(usage, nil, true, "reddit 0", false)

    session, err := reddit.NewLoginSession(
        arguments["<username>"].(string),
        arguments["<password>"].(string),
        "contributors or whatever",
    )

    if err != nil {
        panic(err)
    }

    inte, err := SubredditContributors(session, "kotakuinaction")
    if err != nil {
        panic(err)
    }

    fmt.Println(ints)
}
