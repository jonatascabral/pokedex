package server

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    ctrls "github.com/jonatascabral/pokedex/backend/api/controllers"
    "github.com/jonatascabral/pokedex/backend/api/middlewares"
    "github.com/jonatascabral/pokedex/backend/pkg/database"
    "github.com/jonatascabral/pokedex/backend/pkg/services"
    "github.com/jonatascabral/pokedex/backend/pkg/utils"
    "github.com/sarulabs/di"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
    "github.com/swaggo/swag"
)

const (
    PokemonControllerRef = "pokemon-controller"
    PokemonServiceRef    = "pokemon-service"
    DatabaseRef          = "database"
)

type Server struct {
    Container di.Container
    Engine    *gin.Engine
    Settings  *Settings
}

func StartServer(dotEnvFiles ...string) *Server {
    settings := LoadSettings(dotEnvFiles...)

    engine := gin.New()
    engine.Use(gin.Recovery())
    engine.Use(gin.LoggerWithConfig(gin.LoggerConfig{
        Output: utils.Log.Out,
    }))

    return &Server{
        Engine:   engine,
        Settings: settings,
    }
}

func (s *Server) LoadDependencies() {
    builder, err := di.NewBuilder()
    utils.PanicOnError(err)

    // Database
    builder.Add(di.Def{
        Name:  DatabaseRef,
        Scope: di.App,
        Build: func(c di.Container) (interface{}, error) {
            return database.ConnectDatabase(s.Settings.Database)
        },
    })

    // PokemonService
    builder.Add(di.Def{
        Name:  PokemonServiceRef,
        Scope: di.App,
        Build: func(c di.Container) (interface{}, error) {
            db := c.Get(DatabaseRef).(*gorm.DB)
            return services.NewPokemonService(db), nil
        },
    })

    // PokemonController
    builder.Add(di.Def{
        Name:  PokemonControllerRef,
        Scope: di.App,
        Build: func(c di.Container) (interface{}, error) {
            service := c.Get(PokemonServiceRef).(services.PokemonService)
            return ctrls.NewPokemonController(service), nil
        },
    })

    s.Container = builder.Build()
}

func (s *Server) LoadRoutes() {
    pokemonController := s.Container.Get(PokemonControllerRef).(*ctrls.PokemonController)
    api := s.Engine.Group("/api", middlewares.AuthHandler, middlewares.JsonHeaders)
    {
        pokemons := api.Group("/pokemons")
        {
            pokemons.GET("", pokemonController.GetAll)
        }
    }

    s.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (s *Server) Run() {
    s.ensureContainer()
    swag.Register(swag.Name, NewSwaggerDoc())
    s.Engine.Run(fmt.Sprintf(":%s", s.Settings.AppPort))
}

func (s *Server) RunMigration() {
    s.ensureContainer()
    db := s.Container.Get(DatabaseRef).(*gorm.DB)
    utils.Log.Info("Running migrations")

    migration := database.InitMigration(db)
    migration.Run()
}

func (s *Server) ensureContainer() {
    if s.Container == nil {
        s.LoadDependencies()
    }
}
