package db

import (
	"fmt"
	"os"

	"github.com/blauwiggle/go-grpc-services/internal/rocket"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func New() (Store, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbSslmode := os.Getenv("DB_SSL_MODE")

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", dbUsername, dbPassword, dbName, dbHost, dbPort, dbSslmode)

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return Store{}, err
	}

	return Store{
		db: db,
	}, nil
}

func (s Store) GetRocketById(id string) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}

func (s Store) InsertRocket(rkt rocket.Rocket) (rocket.Rocket, error) {
	return rocket.Rocket{}, nil
}

func (s Store) DeleteRocket(id string) error {
	return nil
}
