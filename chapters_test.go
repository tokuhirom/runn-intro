package main

import (
	"testing"

	"github.com/tokuhirom/runn-intro/testutil"
)

func TestBasics(t *testing.T) {
	// Run all tests in basics
	testutil.RunChapterTests(t, "examples/basics")
}

func TestScenario(t *testing.T) {
	testutil.RunChapterTests(t, "examples/scenario")
}

func TestExprLang(t *testing.T) {
	testutil.RunChapterTests(t, "examples/expr-lang")
}

func TestRunnBuiltins(t *testing.T) {
	testutil.RunChapterTests(t, "examples/runn-builtins")
}
