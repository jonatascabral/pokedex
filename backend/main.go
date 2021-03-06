package main

import (
    s "github.com/jonatascabral/pokedex/backend/api/server"
    "github.com/sirupsen/logrus"
)

var log *logrus.Logger

func main() {
    log = logrus.New()
    server := s.StartServer()
    server.LoadDependencies()
    server.LoadRoutes()
    server.Run()
}

