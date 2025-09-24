package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"slices"
	"strings"

	"github.com/wvdschel/llamaland/pkg/xdg"
)

var (
	configPath = flag.String("config", xdg.ConfigHome()+"/commandrunner/config.json", "Config file path")
)

func main() {
	flag.Parse()

	config, err := loadConfig()
	if err != nil {
		panic(err)
	}

	address := fmt.Sprintf("%s:%d", config.Host, config.Port)

	for path, cmd := range config.Commands {
		p := url.PathEscape(path)

		http.HandleFunc("/"+p, executeCommandHandler(cmd))
		http.HandleFunc("/"+p+".json", executeCommandHandler(cmd))
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<html><body><li>")

		for path := range config.Commands {
			fmt.Fprintf(w, "<a href=\"/%s\">%s</a>", url.PathEscape(path), path)
		}

		fmt.Fprintf(w, "</li></body></html>")
	})

	log.Printf("Starting server on %s", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

func executeCommandHandler(c CommandConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Run cmd and capture stdout, stderr and the exit status
		q := r.URL.Query()
		streaming := getBoolParam(q, "streaming")
		captureStdout := getBoolParam(q, "stdout")
		captureStderr := getBoolParam(q, "stderr")
		structured := strings.HasSuffix(strings.ToLower(r.URL.Path), ".json")

		if structured {
			w.Header().Set("Content-Type", "application/json")
		} else {
			w.Header().Set("Content-Type", "text/plain")
		}

		ctx, cancelFunc := context.WithTimeout(r.Context(), c.Timeout)
		c := exec.CommandContext(ctx, c.Command, c.Args...)
		c.Env = c.Environ()

		out, err := c.Output()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write(out)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func getBoolParam(q url.Values, key string) bool {
	if v, ok := q[key]; ok {
		val := strings.ToLower(strings.Join(v, ","))
		return slices.Contains([]string{"true", "", "1", "y", "yes"}, val)
	}
	return false
}
