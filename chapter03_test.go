package main

import (
	"testing"

	"github.com/tokuhirom/runn-intro/testutil"
)

func TestChapter03(t *testing.T) {
	server := testutil.NewTestServer()
	defer server.Close()

	testutil.RunChapterTests(t, "examples/chapter03", server.URL)
}