package assets

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed bundles
var bundlesFS embed.FS

// bundles имя вложенной папки для встраивания
const bundles = "bundles"

// NewAssetFileServer файловый сервер статики с использованием `встраивания` в приложение
func NewAssetFileServer() (http.Handler, error) {
	fileSystemProvider, err := fs.Sub(bundlesFS, bundles)
	if err != nil {
		return nil, err
	}

	return http.FileServer(http.FS(fileSystemProvider)), nil
}
