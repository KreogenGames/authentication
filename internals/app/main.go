package app

import (
	"context"
	"electro_student/auth/api"
	"electro_student/auth/api/middleware"
	"electro_student/auth/internals/app/db"
	"electro_student/auth/internals/app/handlers"
	"electro_student/auth/internals/app/processors"
	"electro_student/auth/internals/cfg"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Server struct {
	config cfg.Cfg
	ctx    context.Context
	srv    *http.Server
	db     *pgxpool.Pool
}

func NewServer(config cfg.Cfg, ctx context.Context) *Server {
	server := new(Server)
	server.ctx = ctx
	server.config = config
	return server
}

func (server *Server) Serve() {
	log.Println("Starting server")
	log.Println(server.config.GetDBString())
	var err error
	server.db, err = pgxpool.Connect(server.ctx, server.config.GetDBString())
	if err != nil {
		log.Fatalln(err) //Сделает os.Exit(1) и напишет ошибку
	}

	usersStorage := db.NewUsersStorage(server.db)
	rolesStorage := db.NewRolesStorage(server.db)
	gradesStorage := db.NewGradesStorage(server.db)

	usersProcessor := processors.NewUsersProcessor(usersStorage)
	rolesProcessor := processors.NewRolesProcessor(rolesStorage)
	gradesProcessor := processors.NewGradesProcessor(gradesStorage)

	userHandler := handlers.NewUsersHandler(usersProcessor)
	rolesHandler := handlers.NewRolesHandler(rolesProcessor)
	gradesHandler := handlers.NewGradesHandler(gradesProcessor)

	routes := api.CreateRoutes(userHandler, rolesHandler, gradesHandler)
	routes.Use(middleware.RequestLog)

	server.srv = &http.Server{
		Addr:    ":" + server.config.Port,
		Handler: routes,
	}

	log.Println("Server started")

	err = server.srv.ListenAndServe()

	if err != nil {
		log.Fatalln(err)
	}

	return
}

func (server *Server) Shutdown() {
	log.Printf("Server stopped")

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	server.db.Close()

	defer func() {
		cancel()
	}()

	var err error

	if err = server.srv.Shutdown(ctxShutdown); err != nil {
		log.Fatalf("Server shutdown failed: #%v", err)
	}

	log.Printf("Server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}
}
