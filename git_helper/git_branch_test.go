package git_helper

import (
  "testing"
  "fmt"
)

func TestCurrentBranch(t *testing.T) {
  branch, _ := CurrentBranch()
  if branch != "master" {
    fmt.Println("Branch mismatch ", string(branch))
    t.Fail()
  }
}

