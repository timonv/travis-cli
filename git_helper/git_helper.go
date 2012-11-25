package git_helper

import (
  "log"
)

type GitHelper struct {
  cmdBranch, cmdShowOrigin func() (string,error)
}

func NewGitHelper() GitHelper {
  return GitHelper{cmdBranch: cmdBranch, cmdShowOrigin: cmdShowOrigin}
}

func isFatal(e error) {
  if e != nil {
    log.Fatal(e)
  }
}

func isLogged(e error) {
  if e != nil {
    log.Println(e)
  }
}
