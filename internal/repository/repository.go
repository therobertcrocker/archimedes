package repository

// Repository represents the main repository for data entities
type Repository interface {
	Create(entity interface{}) error
	GetbyID(id string) (interface{}, error)
	Update(entity interface{}) error
	Delete(id string) error
}
