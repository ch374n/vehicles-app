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
func SwaggerUIHandler(r *mux.Router) {
	swaggerAssetsContent := fs.FS(swaggerEditorAssets)
	swaggerUIHandler := http.FileServer(http.FS(swaggerAssetsContent))

	r.PathPrefix("/api-docs/").Handler(http.StripPrefix("/api-docs/", swaggerUIHandler))
	r.HandleFunc("/api-docs", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/api-docs/", http.StatusMovedPermanently)
	})
}
