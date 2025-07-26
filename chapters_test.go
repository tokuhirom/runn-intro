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
	testutil.RunTestForFiles(t, []string{
		"examples/runners/http_basic_setup.yml",
		"examples/runners/http_body_formats.yml",
		"examples/runners/http_request_methods.yml",

		// 一旦、このテストは動いてない。後回し。
		//"examples/runners/grpc_basic.yml",

		"examples/runners/db_connections.yml",
		"examples/runners/db_basic_queries.yml",

		// CDPテストは不安定なので除外
		//"examples/runners/cdp_basic.concept.yml",

		"examples/runners/exec_basic.yml",
		"examples/runners/exec_file_operations.yml",
		//"examples/runners/ssh_basic.yml",
		//"examples/runners/ssh_health_check.yml",
	})

	//t.Skip("Skip runners for now")
	//testutil.RunChapterTests(t, "examples/runners")
}

func TestAdvanced(t *testing.T) {
	testutil.RunTestForFiles(t, []string{
		"examples/advanced/loop_basic.yml",
		"examples/advanced/loop_retry.yml",
		"examples/advanced/conditional_basic.yml",
		"examples/advanced/include_basic.yml",
		// common/auth.yml はinclude_basic.ymlからインクルードされるため単体テストは不要
	})
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
