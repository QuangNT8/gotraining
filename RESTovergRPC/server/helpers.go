package server

import (
	"RESTovergRPC/backend"
	storage "RESTovergRPC/backend/postgres"
	"errors"
)

var (
	InvalidDBType = errors.New("invalid db type")
)

// Config backend include Type and db url
func get_backend(dbUrl map[string]string) (backend.Backend, error) {
	switch dbUrl["Type"] {
	case "postgres":
		return storage.New(dbUrl), nil
	default:
		return storage.New(dbUrl), nil
	}
}
