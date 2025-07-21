package main

import (
	"testing"

	"github.com/tokuhirom/runn-intro/testutil"
)

func TestExprLang(t *testing.T) {
	testutil.RunChapterTests(t, "examples/expr-lang")
}
