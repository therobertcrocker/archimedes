package repository

type Entity interface {
	getID() string
} // This could be a more specific interface if you have common methods for all your entities

type Repository interface {
	Create(e Entity) error
	Retrieve(id string) (Entity, error)
	Update(e Entity) error
	Delete(id string) error
}

type QuestRepository interface {
	Repository
	FindAll() ([]Entity, error)
}
