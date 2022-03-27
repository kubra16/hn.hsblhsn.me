package spa

import (
	"io/fs"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/caches"
)

type Asset struct {
	fs.FS
}

func RegisterRoutes(r *mux.Router, source fs.FS) {
	h := &Asset{
		FS: source,
	}
	r.PathPrefix("/").Handler(caches.Middleware(h, time.Hour*72))
}

func (a *Asset) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.FS(a)).ServeHTTP(w, r)
}