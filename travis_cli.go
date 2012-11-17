package main

import (
  "fmt"
  "flag"

  "./adapter"

  /*"io"*/
  /*"strings"*/
  /*"log"*/
)



func main() {
  owner := flag.String("owner","owner", "owner of the repository")
  repo := flag.String("repo","repository",  "name of the repository")
  flag.Parse()

  fmt.Println("Getting status for: ", *owner, "/", *repo)

  adapter := adapter.NewAdapter(*owner,*repo)
  builds := adapter.GetBuilds()
  build := builds[0]

  fmt.Println("Result ", build.HumanResult())
  fmt.Println("Branch ", build.Branch)
  fmt.Println("Commit ", build.HumanCommit())

}

func handleError(err error) {
  if err != nil {
    fmt.Println("ERROR")
  }
}









