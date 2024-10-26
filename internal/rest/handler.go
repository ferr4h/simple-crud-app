package rest

import (
	"github.com/gorilla/mux"
	"net/http"
	"simple-crud-app/internal/domain"
)

type Game interface {
	Create(game domain.Game) error
	GetByID(id int64) (domain.Game, error)
	Get() ([]domain.Game, error)
	Update(id int64, inp domain.GameInput) error
	Delete(id int64) error
}

type Handler struct {
	gameService Game
}

func NewHandler(game Game) *Handler {
	return &Handler{
		gameService: game,
	}
}

func (h *Handler) InitRouter() *mux.Router {
	r := mux.NewRouter()

	games := r.PathPrefix("/games").Subrouter()
	{
		games.HandleFunc("", h.createGame).Methods(http.MethodPost)
		games.HandleFunc("", h.getGames).Methods(http.MethodGet)
		games.HandleFunc("/{id:[0-9]+}", h.getGameByID).Methods(http.MethodGet)
		games.HandleFunc("/{id:[0-9]+}", h.updateGame).Methods(http.MethodPut)
		games.HandleFunc("/{id:[0-9]+}", h.deleteGame).Methods(http.MethodDelete)
	}

	return r
}
