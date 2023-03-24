package main

import (
	"bios/config"
	"bios/controller"
	"bios/store"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	// Read config file and create store
	conf := parseConfig("conf.json")
	db := setupStore(conf)
	defer db.Close()

	// Create controller context
	ctx := controller.Context{DB: db, Conf: conf}

	// Register routes and start server
	router := setupRoutes(ctx)
	fmt.Println(http.ListenAndServe(conf.Addr, router))
}

// parseConfig parses and returns the config file
func parseConfig(path string) config.Config {
	conf, err := config.ParseConfig(path)
	if err != nil {
		panic(fmt.Errorf("error parsing config file: %v", err))
	}
	return conf
}

// setupStore initializes a new store object
func setupStore(conf config.Config) store.Store {
	db, err := store.GetConnection(conf.DB.Host, conf.DB.Port, conf.DB.Name, conf.DB.User, conf.DB.Pass, conf.DB.SSL)
	if err != nil {
		panic(err)
	}
	return db
}

// setupRoutes initializes the controller routes
// and returns the created router object
func setupRoutes(ctx controller.Context) chi.Router {
	router := chi.NewRouter()
	router.Use(
		middleware.Logger,    // Log requests
		middleware.Recoverer, // Recover from panic
		middleware.CleanPath, // Clean up urls with double slash
	)

	// Unauthorized routes
	router.Group(func(r chi.Router) {
		r.Get("/api/movies", ctx.GetMovies)
		r.Get("/api/showings", ctx.GetShowings)
		r.Get("/api/classifications", ctx.GetClassifications)
		r.Get("/api/genres", ctx.GetGenres)
		r.Get("/api/rooms", ctx.GetRooms)
		r.Post("/api/login", ctx.PostLogin)

		// Serve static files
		fs := http.FileServer(http.Dir("./files"))
		r.Handle("/files/*", http.StripPrefix("/files/", fs))
	})

	// Authorized routes
	router.Group(func(r chi.Router) {
		// Authentication middleware based on token
		// r.Use(auth.Context{Store: ctx.Store, Conf: ctx.Conf}.Authenticator)

		r.Post("/api/genres", ctx.EditGenres)
		r.Put("/api/genres", ctx.EditGenres)
		r.Delete("/api/genres", ctx.DeleteGenres)
		r.Post("/api/classifications", ctx.EditClassifications)
		r.Put("/api/classifications", ctx.EditClassifications)
		r.Delete("/api/classifications", ctx.DeleteClassifications)
		r.Post("/api/movies", ctx.EditMovies)
		r.Put("/api/movies", ctx.EditMovies)
		r.Delete("/api/movies", ctx.DeleteMovies)
		r.Post("/api/rooms", ctx.EditRooms)
		r.Put("/api/rooms", ctx.EditRooms)
		r.Delete("/api/rooms", ctx.DeleteRooms)
	})

	return router
}
