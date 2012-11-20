package git_helper

import (
  "testing"
  "fmt"
)

func TestGetRepo(t *testing.T){
  repo := GetRepo()
  if repo.Name != "travis-cli" {
    fmt.Println(repo)
    t.Fail()
  }
}
