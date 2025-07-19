package main

import (
	"testing"

	"github.com/tokuhirom/runn-intro/testutil"
)

func TestChapter01(t *testing.T) {
	// Start test server
	server := testutil.NewTestServer()
	defer server.Close()
	
	// Run all tests in chapter01
	testutil.RunChapterTests(t, "examples/chapter01", server.URL)
}