package git_helper

import (
  "os/exec"
  "regexp"
  "errors"
)

type GitRepo struct {
  Owner, Name string
}

func (gh *GitHelper) GetRepo() GitRepo {
  // err is irrelevant. Parse git response will parse an empty string and throw back a fatal error.
  repostrings, err := gh.cmdShowOrigin()
  isLogged(err)

  owner,name, err := parseGitResponse(repostrings)
  isFatal(err)

  return GitRepo{Name: name, Owner: owner}
}

func parseGitResponse(repostrings string) (string, string, error) {
  var err error
  var owner string
  var repo string

  urlRegexp := regexp.MustCompile(`https:\/\/github\.com\/(.+)\/(.+).git`)

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


