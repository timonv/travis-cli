package main

import (
  "fmt"
  "flag"

  /*"github.com/timonv/travis-cli/adapter"*/
  /*gith "github.com/timonv/travis-cli/git_helper"*/
  "./adapter"
  gith "./git_helper"
)



func main() {
  owner := flag.String("owner","", "owner of the repository")
  repo := flag.String("repo","",  "name of the repository")
  branch := flag.String("branch","",  "name of the branch")
  flag.Parse()

  gh := gith.NewGitHelper()

  if *owner == "" || *repo == "" {
    repository := gh.GetRepo()
    *owner = repository.Owner
    *repo = repository.Name
  }

  if *branch == "" {
    gbranch := gh.CurrentBranch()
    *branch = gbranch.Name
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

func getCorrectBuild(builds []adapter.Build) (adapter.Build, error) {
  var correct adapter.Build
  var err error
  for _, build := range builds {
    if build.Finished_at != ""  {
      correct = build
    }
  }

  return correct, err
}
