package testutil

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/k1LoW/runn"
)

// RunChapterTests runs all YAML tests in a chapter directory
func RunChapterTests(t *testing.T, chapterDir string, serverURL string) {
	t.Helper()

	// Find all YAML files (再帰的に探索)
	var files []string
	err := filepath.WalkDir(chapterDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".yml") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

	// Run each file
	for _, file := range files {
		// Skip conceptual example files and database examples
		baseName := filepath.Base(file)
		if strings.HasSuffix(baseName, ".concept.yml") {
			t.Logf("Skip %s, due to conceptual example files can't run.", baseName)
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
			result := o.Result()

			// Outの結果をバッファに書き出し、.outファイルに保存
			var buf bytes.Buffer
			err = result.Out(&buf, false)
			if err != nil {
				t.Fatal(err)
			}
			outFile := strings.Replace(file, ".yml", ".out", 1)
			if err := os.WriteFile(outFile, buf.Bytes(), 0644); err != nil {
				t.Fatalf("failed to write out file: %v", err)
			}
		})
	}
}
