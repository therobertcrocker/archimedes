package entities

// Quest is a struct that represents a quest
type Quest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Level       int    `json:"level"`
	Rewards     Reward `json:"rewards"`

	// TODO: add rewards and objectives
}

// Reward is a struct that represents a quest reward
type Reward struct {
	Gold           int `json:"gold"`
	Experience     int `json:"experience"`
	TreasureRating int `json:"treasureRating"`

	//TODO: add items
}

func (q *Quest) GetID() string {
	return q.ID
}
