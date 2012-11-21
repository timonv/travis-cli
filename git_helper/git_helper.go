package git_helper

type GitHelper struct {
  cmdBranch, cmdShowOrigin func() (string,error)
}

func NewGitHelper() GitHelper {
  return GitHelper{cmdBranch: cmdBranch, cmdShowOrigin: cmdShowOrigin}
}
