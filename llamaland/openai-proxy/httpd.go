// openai-compat/httpd.go
package openaicompat

import (
	"log"
	"net/http"
)

// stubHandler is a simple stub handler that returns a 200 OK response with an empty body
func stubHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// NewServer creates a new HTTP server with stubbed handlers for OpenAI API endpoints
func NewServer() *http.Server {
	mux := http.NewServeMux()

	// Add handlers for all OpenAI API endpoints
	mux.HandleFunc("/v1/models", stubHandler)
	mux.HandleFunc("/v1/completions", stubHandler)
	mux.HandleFunc("/v1/chat/completions", stubHandler)
	mux.HandleFunc("/v1/embeddings", stubHandler)
	mux.HandleFunc("/v1/files", stubHandler)
	mux.HandleFunc("/v1/fine-tunes", stubHandler)
	mux.HandleFunc("/v1/images/generations", stubHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	return server
}

func StartServer() {
	server := NewServer()
	log.Println("Starting OpenAI-compatible server on port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
