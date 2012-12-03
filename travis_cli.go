package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/timonv/travis-cli/adapter"
	gith "github.com/timonv/travis-cli/git_helper"
	/*"./adapter"*/
	/*gith "./git_helper"*/
)

func main() {
	owner := flag.String("owner", "", "owner of the repository")
	repo := flag.String("repo", "", "name of the repository")
	branch := flag.String("branch", "", "name of the branch")
	flag.Parse()

	gh := gith.NewGitHelper()

	r := fixOwnerRepo(*owner, *repo, gh)
	b := fixBranch(*branch, gh)

	if r.Owner != "" || r.Name != "" {
		fmt.Println("Getting status for: ", r.Owner, "/", r.Name)
		fmt.Println("On branch: ", b.Name)

		adapter := adapter.NewAdapter(r.Owner, r.Name)
		builds := adapter.GetBuilds()
		if len(builds) > 0 {
			build := getCorrectBuild(builds, b.Name)
			if build.Branch == "" {
				log.Fatal("Could not get build")
			}
			fmt.Println("Result ", build.HumanResult())
			fmt.Println("Branch ", build.Branch)
			fmt.Println("Commit ", build.HumanCommit())
			fmt.Println("Finished at ", build.Finished_at)
		} else {
			fmt.Println("Could not get build status")
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

func fixOwnerRepo(o string, r string, gh gith.GitHelper) gith.GitRepo {
	var repo gith.GitRepo
	if o != "" && r != "" {
		repo = gith.GitRepo{Owner: o, Name: r}
	} else {
		repo = gh.GetRepo()
	}
	return repo
}

func fixBranch(b string, gh gith.GitHelper) gith.GitBranch {
	var branch gith.GitBranch
	if b != "" {
		branch = gith.GitBranch{Name: b}
	} else {
		branch = gh.CurrentBranch()
	}
	return branch
}
