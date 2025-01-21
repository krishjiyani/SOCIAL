package main

import (
	"krishjiyani/SOCIAL/internal/store"
	"log"

	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	store  store.Storage
	db     dbConfig
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type config struct {
	addr string
	db   dbConfig
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// Set a timeout value on the request context (ctx), that will signal
	//through ctx.Done() that the request has time out and further
	//processing should be stopper.

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})

	return r
}

// r := chi.NewRouter()
// r.Use(middleware.Logger)
// r.Get("/", func(w http.ResponseWriter,r *http.Request))

// }

// mux := http.NewServeMux()

// mux.HandleFunc("GET/v1/health",app.healthCheckHandler)
//posts

//users

//auth

// return mux
// }

func (app *application) run(mux http.Handler) error {

	/* mount() *http.ServeMux {
		mux := http.NewServeMux()

		return mux
	}                      // run() error {

	*/
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("server has started at %s",
		app.config.addr)
	return srv.ListenAndServe()
}
