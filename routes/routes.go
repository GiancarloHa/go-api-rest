package routes

import (
	"log"
	"net/http"

	"github.com/giancarloha/go-rest-api/controllers"
	"github.com/giancarloha/go-rest-api/middleware"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func HandleRequest() {
	mux := http.NewServeMux()
	mux.Handle("/", otelhttp.NewHandler(http.HandlerFunc(controllers.Home), "Home"))
	mux.Handle("/api/listmangas", otelhttp.NewHandler(http.HandlerFunc(controllers.TodosMangas), "ListMangas"))
	mux.Handle("/api/mangas", otelhttp.NewHandler(http.HandlerFunc(controllers.AddManga), "AddManga"))
	mux.Handle("/api/mangas/ano/{ano}/mes/{mes}", otelhttp.NewHandler(http.HandlerFunc(controllers.TodosMangasMes), "TodosMangasMes"))
	mux.Handle("/api/mangas/{id}", otelhttp.NewHandler(http.HandlerFunc(controllers.MangaOperations), "OperationsManga"))

	handlerWithMiddleware := middleware.ContentTypeMiddleware(mux)
	log.Fatal(http.ListenAndServe(":8000", handlerWithMiddleware))
}
