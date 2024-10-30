package rest

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"simple-crud-app/internal/domain"
	"strconv"
)

func (h *Handler) createGame(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logError("createGame", "read body error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var game domain.Game
	if err = json.Unmarshal(reqBytes, &game); err != nil {
		logError("createGame", "unmarshal request error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.gameService.Create(game)
	if err != nil {
		logError("createGame", "service error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) getGameByID(w http.ResponseWriter, r *http.Request) {
	id, err := extractIdFromRequest(r)
	if err != nil {
		logError("getGameByID", "extract id from request error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	game, err := h.gameService.GetByID(id)
	if err != nil {
		if errors.Is(err, domain.GameNotFound) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		logError("getGameByID", "service error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(game)
	if err != nil {
		logError("getGameByID", "marshal response error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func (h *Handler) getGames(w http.ResponseWriter, r *http.Request) {
	games, err := h.gameService.Get()
	if err != nil {
		logError("getGames", "service error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(games)
	if err != nil {
		logError("getGames", "marshal response error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func (h *Handler) updateGame(w http.ResponseWriter, r *http.Request) {
	id, err := extractIdFromRequest(r)
	if err != nil {
		logError("updateGame", "extract id from request error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logError("updateGame", "read body error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var inp domain.GameInput
	if err = json.Unmarshal(reqBytes, &inp); err != nil {
		logError("updateGame", "unmarshal request error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.gameService.Update(id, inp)
	if err != nil {
		logError("updateGame", "service error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) deleteGame(w http.ResponseWriter, r *http.Request) {
	id, err := extractIdFromRequest(r)
	if err != nil {
		logError("deleteGame", "extract id from request error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.gameService.Delete(id)
	if err != nil {
		logError("deleteGame", "service error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func extractIdFromRequest(r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return 0, err
	}

	if id <= 0 {
		return 0, errors.New("id has to be greater than 0")
	}

	return id, nil
}
