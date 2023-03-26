package main


type App struct {
	db  DB
	apm APM

	server *server.Server
}
