package main

import (
	"github.com/octavore/naga/service"
	admin "github.com/octavore/press/admin"
	"github.com/octavore/press/api"
	"github.com/octavore/press/server"
)

type App struct {
	Server *server.Module
	API    *api.Module
	Admin  *admin.Module
}

func (p *App) Init(c *service.Config) {}

func main() {
	service.Run(&App{})
}
