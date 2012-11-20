package git_helper

import (
  "fmt"
  /*"strings"*/
  "os/exec"
  "regexp"
  "errors"
)

type GitRepo struct {
  Owner, Name string
}

func GetRepo() GitRepo {
  repostrings, _ := cmdShowOrigin()
  owner,name, err := parseGitResponse(repostrings)
  handle(err)

  fmt.Println(owner,name)
  repo := GitRepo{Owner: owner, Name: name}
  return repo
}

func cmdShowOrigin() (string, error) {
  cmd := exec.Command("git", "remote", "show", "origin")
  out, err := cmd.Output()

  handle(err)

  return string(out), err
}

func parseGitResponse(repostrings string) (string, string, error) {
  var err error
  var owner string
  var repo string

  urlRegexp, err := regexp.Compile(`https:\/\/github\.com\/(.+)\/(.+).git`)
  handle(err)

  matches := urlRegexp.FindStringSubmatch(repostrings)
  fmt.Println(matches)
  if len(matches) >= 3 {
    owner = matches[1]
    repo = matches[2]
  } else {
    err = errors.New("Could not determine repo and owner")
  }

  return owner, repo, err
}
