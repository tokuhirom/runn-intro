package main

import (
	"testing"

	"github.com/tokuhirom/runn-intro/testutil"
)

func TestScenario(t *testing.T) {
	testutil.RunChapterTests(t, "examples/scenario")
}
