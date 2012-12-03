package adapter

import (
	"fmt"
	"testing"
)

func TestAdapter(t *testing.T) {
	fmt.Println("Testing api")
	a := NewAdapter("owner", "repo")
	a.StubResponse([]byte(`[{"id":1744837,"repository_id":92936,"number":"1","state":"finished","result":1,"started_at":"2012-06-30T17:30:37Z","finished_at":"2012-06-30T17:56:21Z", "duration":7728,"commit":"4fc848c7bc623bc4d42b15c5bad6586a8dccfb8a","branch":"master","message":"Prepared for travis","event_type":"push"}]`))

	builds := a.GetBuilds()
	if len(builds) == 0 {
		fmt.Println("Builds: ", builds)
		t.Fail()
	}

	b := builds[0]
	if b.Branch != "master" {
		fmt.Println("Build: ", b)
		t.Fail()
	}
}
