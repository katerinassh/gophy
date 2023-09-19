package app

import (
	"crud-go/internal/adapters"
	"crud-go/internal/controllers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func createApp(app Application) *http.Server {
	itemRepository := adapters.NewPostgreSQLRepository(app.Pg)

	r := mux.NewRouter()
	r.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"status": "running"}`)
	})

	r.Handle("/items", controllers.GetItems(app.Context, itemRepository)).Methods("GET")
	r.Handle("/items/{id}", controllers.GetItemById()).Methods("GET")

	return &http.Server{
		Addr:    fmt.Sprintf(":%v", app.Port),
		Handler: r,
	}
}
