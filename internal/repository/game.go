package repository

import (
	"database/sql"
	"fmt"
	"simple-crud-app/internal/domain"
	"strings"
)

type Games struct {
	db *sql.DB
}

func NewGames(db *sql.DB) *Games {
	return &Games{db}
}

func (b *Games) Create(game domain.Game) error {
	_, err := b.db.Exec("INSERT INTO games (title, developer, publisher, genre, publication_date, rating) values ($1, $2, $3, $4, $5, $6)",
		game.Title, game.Developer, game.Publisher, game.Genre, game.PublicationDate, game.Rating)

	return err
}

func (b *Games) GetByID(id int64) (domain.Game, error) {
	var game domain.Game
	err := b.db.QueryRow("SELECT id, title, developer, publisher, genre, publication_date, rating FROM games WHERE id=$1", id).
		Scan(&game.ID, &game.Title, &game.Developer, &game.Publisher, &game.Genre, &game.PublicationDate, &game.Rating)
	if err == sql.ErrNoRows {
		return game, domain.GameNotFound
	}

	return game, err
}

func (b *Games) Get() ([]domain.Game, error) {
	rows, err := b.db.Query("SELECT id, title, developer, publisher, genre, publication_date, rating FROM games")
	if err != nil {
		return nil, err
	}

	games := make([]domain.Game, 0)
	for rows.Next() {
		var game domain.Game
		if err := rows.Scan(&game.ID, &game.Title, &game.Developer, &game.Publisher, &game.Genre, &game.PublicationDate, &game.Rating); err != nil {
			return nil, err
		}

		games = append(games, game)
	}

	return games, rows.Err()
}

func (b *Games) Update(id int64, inp domain.GameInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if inp.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *inp.Title)
		argId++
	}

	if inp.Developer != nil {
		setValues = append(setValues, fmt.Sprintf("developer=$%d", argId))
		args = append(args, *inp.Developer)
		argId++
	}

	if inp.Publisher != nil {
		setValues = append(setValues, fmt.Sprintf("publisher=$%d", argId))
		args = append(args, *inp.Publisher)
		argId++
	}

	if inp.Genre != nil {
		setValues = append(setValues, fmt.Sprintf("genre=$%d", argId))
		args = append(args, *inp.Genre)
		argId++
	}

	if inp.PublicationDate != nil {
		setValues = append(setValues, fmt.Sprintf("publication_date=$%d", argId))
		args = append(args, *inp.PublicationDate)
		argId++
	}

	if inp.Rating != nil {
		setValues = append(setValues, fmt.Sprintf("rating=$%d", argId))
		args = append(args, *inp.Rating)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE games SET %s WHERE id=$%d", setQuery, argId)
	args = append(args, id)

	_, err := b.db.Exec(query, args...)
	return err
}

func (b *Games) Delete(id int64) error {
	_, err := b.db.Exec("DELETE FROM games WHERE id=$1", id)

	return err
}
