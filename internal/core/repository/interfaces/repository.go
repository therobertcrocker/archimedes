package interfaces

type Entity interface {
	GetID() string
}

type Repository interface {
	Create(e Entity) error
	Retrieve(id string) (Entity, error)
	Update(e Entity) error
	Delete(id string) error
}
