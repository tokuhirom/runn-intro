package main

import (
	"os"
	"testing"

	"github.com/tokuhirom/runn-intro/testutil"
)

// TestSingleFile runs a single YAML file specified by TEST_FILE environment variable
// Usage: TEST_FILE=examples/runn-builtins/merge_example.yml go test -run TestSingleFile ./...
func TestSingleFile(t *testing.T) {
	testFile := os.Getenv("TEST_FILE")
	if testFile == "" {
		t.Skip("TEST_FILE environment variable not set")
	}

	// Check if file exists
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Fatalf("File not found: %s", testFile)
	}

	testutil.RunTestForFiles(t, []string{testFile})
}