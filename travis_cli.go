package main

import (
  "fmt"
  "flag"
  /*"errors"*/

  "github.com/timonv/travis-cli/adapter"
  "github.com/timonv/travis-cli/git_helper"

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
      build,_ := getCorrectBuild(builds)
      fmt.Println("Result ", build.HumanResult())
      fmt.Println("Branch ", build.Branch)
      fmt.Println("Commit ", build.HumanCommit())
      fmt.Println("Finished at ", build.Finished_at)
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

func getCorrectBuild(builds []adapter.Build) (adapter.Build, error) {
  var correct adapter.Build
  var err error
  for _, build := range builds {
    if build.Finished_at != "" {
      correct = build
    }
  }

  return correct, err
}
