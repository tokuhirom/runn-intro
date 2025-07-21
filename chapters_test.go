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

func TestRunners(t *testing.T) {
	t.Skip("Skip runners for now")
	testutil.RunChapterTests(t, "examples/runners")
}

func TestAdvanced(t *testing.T) {
	t.Skip("Skip advanced for now")
	testutil.RunChapterTests(t, "examples/advanced")
}

func TestTestHelpers(t *testing.T) {
	t.Skip("Skip test-helpers for now")
	testutil.RunChapterTests(t, "examples/test-helpers")
}

func TestExprLang(t *testing.T) {
	testutil.RunChapterTests(t, "examples/expr-lang")
}

func TestRunnBuiltins(t *testing.T) {
	testutil.RunChapterTests(t, "examples/runn-builtins")
}

func TestPractices(t *testing.T) {
	t.Skip("Skip practices for now")
	testutil.RunChapterTests(t, "examples/practices")
}

func TestReferences(t *testing.T) {
	t.Skip("Skip references for now")
	testutil.RunChapterTests(t, "examples/references")
}
