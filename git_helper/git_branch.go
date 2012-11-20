package git_helper

import (
  "fmt"
  "os/exec"
  "strings"
  "errors"
)

func CurrentBranch() (string, error) {
  raw_branches, err := cmdBranch()
  current_branch, err := getCurrentBranch(raw_branches)
  handle(err)
  return current_branch, err
}

func cmdBranch() (string, error) {
  cmd := exec.Command("git", "branch")
  out, err := cmd.Output()

  handle(err)

  return string(out), err
}

func handle(e error) {
  if e != nil {
    fmt.Println(e)
  }
}

func getCurrentBranch(b string) (string, error) {
  var splitted []string
  var current string
  var err error

  splitted = strings.Split(b, "\n")

  for _, branch := range splitted {
    if string(branch) != "" && string(branch[0]) == "*" {
      current = branch[2:]
    }
  }

  if current == "" {
    err = errors.New("Can't find current branch")
  }
  return current, err
}
