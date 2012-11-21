package git_helper

import (
  "testing"
  "fmt"
)

func TestGitHelperBranch (t *testing.T) {
  helper := NewGitHelper()
  helper.cmdBranch = func() (string,error) {
    return "* master", nil
  }

  current, err := helper.CurrentBranch()
  if err != nil {
    fmt.Println(err)
    t.Fail()
  }
  if current.Name != "master" {
    fmt.Println("Branch mismatch ", string(current.Name))
    t.Fail()
  }
}

func TestGitHelperRepo (t *testing.T) {
  helper := NewGitHelper()
  helper.cmdShowOrigin = func() (string,error) {
    return `https://github.com/timonv/travis-cli.git`, nil
  }

  current, err := helper.GetRepo()
  if err != nil {
    fmt.Println(err)
    t.Fail()
  }
  if current.Name != "travis-cli" && current.Owner != "timonv" {
    fmt.Println("Branch mismatch ", current)
    t.Fail()
  }
}




