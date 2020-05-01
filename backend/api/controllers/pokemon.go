package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/jonatascabral/pokedex/backend/pkg/services"
)

type PokemonController struct {
	service services.PokemonServiceInterface
}

func NewPokemonController(service services.PokemonServiceInterface) *PokemonController {
	return &PokemonController{
		service: service,
	}
}

func (c *PokemonController) GetAll(ctx *gin.Context) {

}
