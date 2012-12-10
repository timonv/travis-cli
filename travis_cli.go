package main

import (
	"flag"
	"log"
  "os"

  "github.com/timonv/travis-cli/adapter"
  "github.com/timonv/travis-cli/git_helper"
)

func main() {
	owner := flag.String("owner", "", "owner of the repository")
	repo := flag.String("repo", "", "name of the repository")
	branch := flag.String("branch", "", "name of the branch")
	flag.Parse()

	gh := git_helper.NewGitHelper()

	r := fixOwnerRepo(*owner, *repo, gh)
	b := fixBranch(*branch, gh)

	if r.Owner != "" || r.Name != "" {
		adapter := adapter.NewAdapter(r.Owner, r.Name)
		builds := adapter.GetBuilds()
		if len(builds) > 0 {
			build := getCorrectBuild(builds, b.Name)
			if build.Branch == "" {
				log.Fatal("Could not get build")
			}
      printBuild(build)
		} else {
			log.Fatal("Could not get build status")
		}
	}
}

func getCorrectBuild(builds []adapter.Build, branch string) adapter.Build {
	var correct adapter.Build
	for _, build := range builds {
		if build.Finished_at != "" && build.Branch == branch {
			correct = build
			break
		}
	}

	return correct
}

func fixOwnerRepo(o string, r string, gh git_helper.GitHelper) git_helper.GitRepo {
	var repo git_helper.GitRepo
	if o != "" && r != "" {
		repo = git_helper.GitRepo{Owner: o, Name: r}
	} else {
		repo = gh.GetRepo()
	}
	return repo
}

func fixBranch(b string, gh git_helper.GitHelper) git_helper.GitBranch {
	var branch git_helper.GitBranch
	if b != "" {
		branch = git_helper.GitBranch{Name: b}
	} else {
		branch = gh.CurrentBranch()
	}
	return branch
}

func printBuild(build adapter.Build) {
  red :="\x1b[31m" 
  green := "\x1b[32m" 
  reset := "\x1b[0m"
  result := build.HumanResult()
  if result == "Passed" {
    result = green + result + reset
  } else {
    result = red + result + reset
  }
  os.Stdout.Write([]byte(result))
  os.Stdout.Write([]byte("\t"))
  os.Stdout.Write([]byte(build.Branch))
  os.Stdout.Write([]byte("\t"))
  os.Stdout.Write([]byte(build.HumanCommit()))
  os.Stdout.Write([]byte("\t"))
  os.Stdout.Write([]byte(build.Finished_at))
  os.Stdout.Write([]byte("\t"))
  os.Stdout.Write([]byte(build.Finished_at))
  os.Stdout.Write([]byte("\n"))
}
