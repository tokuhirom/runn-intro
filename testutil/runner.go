package testutil

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/k1LoW/runn"
)

// RunChapterTests runs all YAML tests in a chapter directory
func RunChapterTests(t *testing.T, chapterDir string, serverURL string) {
	t.Helper()
	
	// Find all YAML files
	files, err := filepath.Glob(filepath.Join(chapterDir, "*.yml"))
	if err != nil {
		t.Fatal(err)
	}
	
	// Also check subdirectories
	subFiles, err := filepath.Glob(filepath.Join(chapterDir, "*/*.yml"))
	if err != nil {
		t.Fatal(err)
	}
	files = append(files, subFiles...)
	
	// Run each file
	for _, file := range files {
		// Skip conceptual example files and database examples
		baseName := filepath.Base(file)
		if baseName == "intro-multi-protocol.yml" || baseName == "database-query.yml" {
			continue
		}
		
		t.Run(filepath.Base(file), func(t *testing.T) {
			// Override runners to use test server URL
			opts := []runn.Option{
				runn.T(t),
				runn.Runner("req", serverURL),
				runn.Runner("api", serverURL),
				runn.Runner("auth", serverURL),
				runn.Runner("blog-api", serverURL),
				runn.Runner("https://api.example.com", serverURL),
				runn.Runner("https://auth.example.com", serverURL),
				runn.Runner("https://blog-api.example.com", serverURL),
				runn.Runner("http://localhost:8080", serverURL),
			}
			
			o, err := runn.Load(file, opts...)
			if err != nil {
				t.Fatal(err)
			}
			
			if err := o.RunN(context.Background()); err != nil {
				t.Fatal(err)
			}
		})
	}
}