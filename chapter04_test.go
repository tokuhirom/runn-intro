package main

import (
	"testing"

	"github.com/tokuhirom/runn-intro/testutil"
)

func TestChapter04(t *testing.T) {
	testutil.RunChapterTests(t, "examples/chapter04")
}

// TODO: remove this, after fixing the test
func TestCompareBasic(t *testing.T) {
	testutil.RunTestForFiles(t, []string{
		"examples/chapter04/diff_example.yml",
	})
}
