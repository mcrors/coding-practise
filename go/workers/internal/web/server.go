// Package web contains all things http and web related, such as registering routes
package web

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"text/template"
)

//go:embed templates static
var embeddedFiles embed.FS

type server struct {
	tmpls     map[string]*template.Template
	fragTmpls map[string]*template.Template
	dev       bool
}

func RegisterRoutes(mux *http.ServeMux, dev bool) error {
	s := &server{dev: dev}
	staticFS, err := fs.Sub(embeddedFiles, "static")
	if err != nil {
		return fmt.Errorf("static sub-fs: %w", err)
	}
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServerFS(staticFS)))

	mux.HandleFunc("GET /", s.indexHandler)
	return nil
}
