package main

import (
	"testing"

	"github.com/tokuhirom/runn-intro/testutil"
)

func TestChapter02(t *testing.T) {
	server := testutil.NewTestBlogServer()
	defer server.Close()

	testutil.RunChapterTests(t, "examples/chapter02", server.URL)
}
