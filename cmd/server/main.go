package main

import (
	"log"

	"github.com/blauwiggle/go-grpc-services/internal/db"
	"github.com/blauwiggle/go-grpc-services/internal/rocket"
)

func Run() error {
	// responsible for init and starting our gRPC server

	rocketStore, err := db.New()
	if err != nil {
		return err
	}

	_ = rocket.New(rocketStore)

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
