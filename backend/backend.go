package backend

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/tasylab/hn.hsblhsn.me/backend/graphql"
	"github.com/tasylab/hn.hsblhsn.me/backend/images"
)

// RegisterRoutes registers the routes for the backend.
func RegisterRoutes(
	router *mux.Router,
	gql *graphql.GQLHandler,
	imgProxy *images.ImageProxyHandler,
	socialPreviewHandler *images.SocialPreviewHandler,
) {
	router.Path("/graphql").
		Methods(http.MethodGet, http.MethodPost).
		Handler(gql)
	router.Path("/explorer").
		Methods(http.MethodGet).
		Handler(playground.Handler("GraphQL Explorer", "/graphql"))
	router.Path("/images/proxy.jpeg").
		Methods(http.MethodGet).
		Queries("size", "{size}", "src", "{src}").
		Handler(imgProxy)
	router.Path("/images/social_preview.jpeg").
		Methods(http.MethodGet).
		Queries("title", "{title}").
		Handler(socialPreviewHandler)
}
