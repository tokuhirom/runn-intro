package main

import (
	"testing"

	"github.com/tokuhirom/runn-intro/testutil"
)

func TestBasics(t *testing.T) {
	// Run all tests in basics
	testutil.RunChapterTests(t, "examples/basics")
}
