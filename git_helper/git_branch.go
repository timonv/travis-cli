package git_helper

import (
  /*"fmt"*/
  "os/exec"
  "strings"
  "errors"
)

type GitBranch struct {
  Name string
}

func (gh *GitHelper) CurrentBranch() (GitBranch, error) {
  branch_name, err := gh.getCurrentBranch()
  current_branch := GitBranch{Name: branch_name}
  return current_branch, err
}

func (gh *GitHelper) getCurrentBranch() (string, error) {
  var current string
  raw_output, err := gh.cmdBranch()
  splitted := strings.Split(raw_output, "\n")

  for _, branch := range splitted {
    if string(branch) != "" && string(branch[0]) == "*" {
      current = branch[2:]
      break
    }
  }

  if current == "" {
    err = errors.New("Can't find current branch")
  }
  return current, err
}

func cmdBranch() (string, error) {
  cmd := exec.Command("git", "branch")
  out, err := cmd.Output()
  return string(out), err
}
