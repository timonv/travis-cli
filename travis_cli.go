package main

import (
  "fmt"
  "flag"

  "./adapter"
  "./git_helper"

  /*"io"*/
  /*"strings"*/
  /*"log"*/
)



func main() {
  owner := flag.String("owner","", "owner of the repository")
  repo := flag.String("repo","",  "name of the repository")
  branch := flag.String("branch","",  "name of the branch")
  flag.Parse()

  if *owner == "" || *repo == "" {
    repository := git_helper.GetRepo()
    *owner = repository.Owner
    *repo = repository.Name
  }

  if *branch == "" {
    *branch, _ = git_helper.CurrentBranch()
  }

  if *owner != "" || *repo != "" {
    fmt.Println("Getting status for: ", *owner, "/", *repo)
    fmt.Println("On branch: ", *branch)

    adapter := adapter.NewAdapter(*owner,*repo)
    builds := adapter.GetBuilds()
    if len(builds) > 0 {
      build := builds[0]
      fmt.Println("Result ", build.HumanResult())
      fmt.Println("Branch ", build.Branch)
      fmt.Println("Commit ", build.HumanCommit())
    } else {
      fmt.Println("Could not get build status")
    }

  }

}

func handleError(err error) {
  if err != nil {
    fmt.Println(err)
  }
}









