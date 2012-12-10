package adapter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
Builds
SAMPLE
{
  "id":1744837,"repository_id":92936,"number":"1","state":"finished","result":1,"started_at":"2012-06-30T17:30:37Z","finished_at":"2012-06-30T17:56:21Z",
  "duration":7728,"commit":"4fc848c7bc623bc4d42b15c5bad6586a8dccfb8a","branch":"master","message":"Prepared for travis","event_type":"push"
}
*/
type Build struct {
	Branch, Commit, Messages, Finished_at string
	Repository_id, Result                 float64
}

func (b Build) HumanResult() string {
	var status string

	if b.Result == float64(0) {
		status = "Passed"
	} else {
		status = "Failed"
	}
	return status
}

func (b Build) HumanCommit() string {
	return b.Commit[0:6]
}

// Adapter
type adapter struct {
	Owner, Repo     string
	StubbedResponse []byte
}

func NewAdapter(owner, repo string) adapter {
	a := adapter{Owner: owner, Repo: repo}
	return a
}

func (a *adapter) StubResponse(r []byte) {
	a.StubbedResponse = r
}

func (a adapter) GetBuilds() []Build {
	url := a.getBuildsURL()
	raw_json, err := a.makeRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	var builds []Build
	errm := json.Unmarshal(raw_json, &builds)
	if errm != nil {
		fmt.Println(errm)
	}

	return builds
}

func (a adapter) makeRequest(url string) ([]byte, error) {
	var err error
	var raw_json []byte
	if a.StubbedResponse == nil {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		raw_json, err = ioutil.ReadAll(resp.Body)
	} else {
		raw_json, err = a.StubbedResponse, nil
	}
	return raw_json, err
}

func (a adapter) getBuildsURL() string {
	pre_url := "http://api.travis-ci.org/%s/%s/builds.json"
	return fmt.Sprintf(pre_url, a.Owner, a.Repo)
}
