package git_helper

import (
  "os/exec"
  "regexp"
  "errors"
)

type GitRepo struct {
  Owner, Name string
}

func (gh *GitHelper) GetRepo() (GitRepo, error) {
  repostrings, err := gh.cmdShowOrigin()
  owner,name, err := parseGitResponse(repostrings)
  gr := GitRepo{Name: name, Owner: owner}

  return gr, err
}

func parseGitResponse(repostrings string) (string, string, error) {
  var err error
  var owner string
  var repo string

  urlRegexp, err := regexp.Compile(`https:\/\/github\.com\/(.+)\/(.+).git`)

  matches := urlRegexp.FindStringSubmatch(repostrings)
  if len(matches) >= 3 {
    owner = matches[1]
    repo = matches[2]
  } else {
    err = errors.New("Could not determine repo and owner")
  }

  return owner, repo, err
}

func cmdShowOrigin() (string, error) {
  cmd := exec.Command("git", "remote", "show", "origin")
  out, err := cmd.Output()


  return string(out), err
}


