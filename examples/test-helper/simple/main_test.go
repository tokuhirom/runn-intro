package main

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/k1LoW/runn"
)

func TestAPI(t *testing.T) {
	// テストサーバーの起動
	server := NewServer()
	ts := httptest.NewServer(server.Handler())
	defer ts.Close()

	// runnの設定
	opts := []runn.Option{
		runn.T(t),
		runn.Runner("api", ts.URL),
	}

	// YAMLシナリオの実行
	o, err := runn.Load("testdata/api_test.yml", opts...)
	if err != nil {
		t.Fatal(err)
	}

	if err := o.RunN(context.Background()); err != nil {
		t.Fatal(err)
	}
}