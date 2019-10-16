package stdos

import (
	"io/ioutil"
	"testing"

	"github.com/markbates/pkger/here"
	"github.com/markbates/pkger/pkging"
	"github.com/markbates/pkger/pkging/pkgtest"
)

func Test_Pkger(t *testing.T) {
	suite, err := pkgtest.NewSuite("stdos", func() (pkging.Pkger, error) {
		her, err := here.Current()
		if err != nil {
			return nil, err
		}
		dir, err := ioutil.TempDir("", "stdos")
		if err != nil {
			return nil, err
		}

		her.Dir = dir

		mypkging, err := New(her)
		if err != nil {
			return nil, err
		}

		return mypkging, nil
	})
	if err != nil {
		t.Fatal(err)
	}

	suite.Test(t)
}
