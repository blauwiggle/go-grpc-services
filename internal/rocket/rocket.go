package rocket

import "context"

// Rocket - should contain the definition of the Rocket type
type Rocket struct {
	Id      string
	Name    string
	Type    string
	Flights int
}

// Store - defines the interface we expect out database implementation to follow
type Store interface {
	GetRocketById(id string) (Rocket, error)
	InsertRocket(r Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

// Service - our rocket service, responsible for updating the rocket inventory
type Service struct {
	Store Store
}

// New - returns a new instance of the rocket service
func New(store Store) Service {
	return Service{
		Store: store,
	}
}

// GetRocketById - returns a rocket by id
func (s Service) GetRocketById(ctx context.Context, id string) (Rocket, error) {
	rkt, err := s.Store.GetRocketById(id)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// InsertRocket - inserts a new rocket into the database
func (s Service) InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	rkt, err := s.Store.InsertRocket(rkt)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// DeleteRocket - deletes a rocket from the database
func (s Service) DeleteRocket(ctx context.Context, id string) error {
	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}
	return nil
}
