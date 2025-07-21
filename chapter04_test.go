package main

import (
	"testing"

	"github.com/tokuhirom/runn-intro/testutil"
)

func TestRunnBuiltins(t *testing.T) {
	testutil.RunChapterTests(t, "examples/runn-builtins")
}
