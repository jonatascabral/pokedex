package services

import (
    "github.com/jinzhu/gorm"
    "github.com/jonatascabral/pokedex/backend/pkg/models"
)

type PokemonServiceInterface interface{
    GetAll() ([]*models.Pokemon, error)
    GetByName(string) (*models.Pokemon, error)
    GetById(int) (*models.Pokemon, error)
    GetByNumber(int) (*models.Pokemon, error)
}

type PokemonService struct {
    database *gorm.DB
}

func NewPokemonService(database *gorm.DB) PokemonService {
    return PokemonService{
        database: database,
    }
}

func (p PokemonService) GetAll() (pokemons []*models.Pokemon, err error) {
    return
}

func (p PokemonService) GetByName(name string) (pokemon *models.Pokemon, err error) {
    return
}

func (p PokemonService) GetById(id int) (pokemon *models.Pokemon, err error) {
    return
}

func (p PokemonService) GetByNumber(number int) (pokemon *models.Pokemon, err error) {
    return
}
