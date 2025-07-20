package testutil

import (
	"bytes"
	"context"
	"github.com/mccutchen/go-httpbin/v2/httpbin"
	"net/http/httptest"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/k1LoW/runn"
)

// RunChapterTests runs all YAML tests in a chapter directory
func RunChapterTests(t *testing.T, chapterDir string) {
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

	// go-httpbinサーバーを起動（必要な場合に使うため）
	httpbinObj := httpbin.New()
	httpbinServer := httptest.NewServer(httpbinObj.Handler())
	defer httpbinServer.Close()
	httpbinServerURL := httpbinServer.URL

	// blogサーバーを起動（必要な場合に使うため）
	blogServer := NewTestBlogServer()
	defer blogServer.Close()
	blogServerURL := blogServer.URL

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
			}

			// go-httpbin runnerが必要な場合はここでURLをセット
			keys, err := GetRunnerKeys(file)
			if err != nil {
				t.Fatal(err)
			}

			for _, key := range keys {
				if key == "httpbin" {
					// keys に httpbin が含まれていたら httpbin を起動し、serverURL を指定
					opts = append(opts, runn.Runner("httpbin", httpbinServerURL))
				}
				if key == "blog" {
					// keys に blog が含まれていたら blog を起動し、serverURL を指定
					opts = append(opts, runn.Runner("blog", blogServerURL))
				}
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

			// ANSIエスケープシーケンスを除去
			plain := buf.Bytes()
			plain = stripANSI(plain)

			outFile := strings.Replace(file, ".yml", ".out", 1)
			if err := os.WriteFile(outFile, plain, 0644); err != nil {
				t.Fatalf("failed to write out file: %v", err)
			}
		})
	}
}

// ANSIエスケープシーケンス除去用関数
func stripANSI(b []byte) []byte {
	re := regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)
	return re.ReplaceAll(b, []byte(""))
}
