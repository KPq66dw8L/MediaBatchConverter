package main

import (
	"media_batch_converter/backend/internal/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Ctx)

	r.Route("/folders", func(r chi.Router) {
		// r.Get("/{id}", folders.GetFolder)
	})

	logrus.Info("[INFO] Server started on :8080")
	http.ListenAndServe(":8080", r)
}
