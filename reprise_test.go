package reprise

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestWriteSimple(t *testing.T) {
	testConfig := &Banner{
		Name:         "Name",
		Description:  "Description",
		Version:      "Development build",
		GoVersion:    "Unknown",
		WebsiteURL:   "https://spoonboy.io",
		VcsURL:       "https://github.com/spoonboy-io/myrepo",
		VcsName:      "Github",
		EmailAddress: "email@emailaddress.com",
	}

	want := `
**************************************************
* Name - Description                             *
* ------------------                             *
* Version: Development build                     *
* Go Build Version: Unknown                      *
* Website: https://spoonboy.io                   *
* Github:  https://github.com/spoonboy-io/myrepo *
* Email:   email@emailaddress.com                *
**************************************************

`
	got := captureFmt(func() {
		WriteSimple(testConfig)
	})

	assertEqual(t, got, want)
}

func captureFmt(f func()) string {
	revertStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = revertStdout

	return string(out)
}

func assertEqual(t *testing.T, got, want string) {
	if got != want {
		// if failing uncomment the below line to get a non-pretty diff of 'got' vs 'want'
		// fmt.Printf("want: %q, got: %q", want, got)
		t.Errorf("Fail wanted %s got %s", want, got)
	}
}
