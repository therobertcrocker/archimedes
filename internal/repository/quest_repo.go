package repository

import (
	"errors"

	repository "github.com/therobertcrocker/archimedes/internal/repository/interfaces"
)

type Quest struct {
	ID          string
	Name        string
	Description string
	Level       int
	Experience  int
}

type QuestRepository struct {
	quests map[string]Quest
}

func (q Quest) getID() string {
	return q.ID
}

func NewQuestRepository() *QuestRepository {
	return &QuestRepository{
		quests: make(map[string]Quest),
	}
}

func (r *QuestRepository) Create(e repository.Entity) error {
	quest, ok := e.(Quest)
	if !ok {
		return errors.New("entity is not a quest")
	}

}
