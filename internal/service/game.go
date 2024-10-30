package service

import (
	"simple-crud-app/internal/domain"
)

type GamesRepository interface {
	Create(game domain.Game) error
	GetByID(id int64) (domain.Game, error)
	Get() ([]domain.Game, error)
	Update(id int64, inp domain.GameInput) error
	Delete(id int64) error
}

type Games struct {
	repository GamesRepository
}

func NewGames(repository GamesRepository) *Games {
	return &Games{
		repository: repository,
	}
}

func (b *Games) Create(game domain.Game) error {
	return b.repository.Create(game)
}

func (b *Games) GetByID(id int64) (domain.Game, error) {
	return b.repository.GetByID(id)
}

func (b *Games) Get() ([]domain.Game, error) {
	return b.repository.Get()
}

func (b *Games) Update(id int64, inp domain.GameInput) error {
	return b.repository.Update(id, inp)
}

func (b *Games) Delete(id int64) error {
	return b.repository.Delete(id)
}
