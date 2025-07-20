package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tokuhirom/runn-intro/testutil"
)

func main() {
	// Create the same test server as used in tests
	server := testutil.NewTestBlogServer()
	defer server.Close()

	// Extract the port from test server URL
	fmt.Printf("Test server (with go-httpbin) running at: %s\n", server.URL)
	fmt.Println("Press Ctrl+C to stop...")

	// Create a real server on port 8080
	realServer := &http.Server{
		Addr:    ":8080",
		Handler: server.Config.Handler,
	}

	log.Fatal(realServer.ListenAndServe())
}
