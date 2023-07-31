package implementations

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/therobertcrocker/archimedes/internal/core/repository/interfaces"
	"github.com/therobertcrocker/archimedes/internal/entities"
)

// QuestRepository is a struct that holds a map of quests.
type QuestRepository struct {
	quests map[string]*entities.Quest
}

// NewQuestRepository initializes a new QuestRepository and returns a pointer to it.
func NewQuestRepository() *QuestRepository {
	log.Info("Initializing new QuestRepository")
	return &QuestRepository{
		quests: make(map[string]*entities.Quest),
	}
}

// Create adds a new quest to the repository.
func (r *QuestRepository) Create(e interfaces.Entity) error {
	// Type assert the entity to a Quest.
	quest, ok := e.(*entities.Quest)
	if !ok {
		log.Error("Attempted to add non-quest entity to QuestRepository")
		return errors.New("entity is not a quest")
	}
	// Check if the quest already exists in the repository.
	if _, ok := r.quests[quest.GetID()]; ok {
		log.WithField("id", quest.GetID()).Error("Attempted to add duplicate quest to QuestRepository")
		return errors.New("quest already exists")
	}
	// Add the quest to the repository.
	r.quests[quest.GetID()] = quest
	log.WithField("id", quest.GetID()).Info("Added quest to QuestRepository")
	return nil
}

// Retrieve gets a quest from the repository by its ID.
func (r *QuestRepository) Retrieve(id string) (interfaces.Entity, error) {
	// Check if the quest exists in the repository.
	if quest, ok := r.quests[id]; ok {
		log.WithField("id", id).Info("Retrieved quest from QuestRepository")
		return quest, nil
	}
	log.WithField("id", id).Error("Failed to retrieve quest from QuestRepository")
	return nil, errors.New("quest not found")
}

// Update modifies an existing quest in the repository.
func (r *QuestRepository) Update(e interfaces.Entity) error {
	// Type assert the entity to a Quest.
	quest, ok := e.(*entities.Quest)
	if !ok {
		log.Error("Attempted to update non-quest entity in QuestRepository")
		return errors.New("entity is not a quest")
	}
	// Check if the quest exists in the repository.
	if _, ok := r.quests[quest.GetID()]; !ok {
		log.WithField("id", quest.GetID()).Error("Attempted to update non-existent quest in QuestRepository")
		return errors.New("quest does not exist")
	}
	// Update the quest in the repository.
	r.quests[quest.GetID()] = quest
	log.WithField("id", quest.GetID()).Info("Updated quest in QuestRepository")
	return nil
}

// Delete removes a quest from the repository by its ID.
func (r *QuestRepository) Delete(id string) error {
	// Check if the quest exists in the repository.
	if _, ok := r.quests[id]; !ok {
		log.WithField("id", id).Error("Attempted to delete non-existent quest from QuestRepository")
		return errors.New("quest does not exist")
	}
	// Delete the quest from the repository.
	delete(r.quests, id)
	log.WithField("id", id).Info("Deleted quest from QuestRepository")
	return nil
}
