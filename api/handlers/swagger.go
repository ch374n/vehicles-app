package handlers

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	//go:embed openapi-spec.yaml
	//go:embed index.html
	swaggerEditorAssets embed.FS
)

// Build an http handler that returns swagger UI assets.
func SwaggerUIHandler(r mux.Router) {
	swaggerAssetsContent := fs.FS(swaggerEditorAssets)
	r.Handle("/api-docs/*", http.StripPrefix("/api-docs", http.FileServer(http.FS(swaggerAssetsContent))))
	r.Handle("/api-docs", http.RedirectHandler("/api-docs/", http.StatusMovedPermanently))
}