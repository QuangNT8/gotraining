package server

import (
	"RESTovergRPC/backend"
	api "RESTovergRPC/directory"
	"fmt"

	"golang.org/x/net/context"
)

// Directory implements the DirectoryServer
type Directory struct {
	backend backend.Backend
}

func NewDirectoryServer(dbUrl map[string]string) (*Directory, error) {

	db, err := get_backend(dbUrl)

	if err != nil {
		return nil, err
	}
	return &Directory{backend: db}, err
}

// CreateDirectory create a directory to stores entries
func (d *Directory) CreateDirectory(ctx context.Context, req *api.DirectoryRequest) (*api.SuccessResponse, error) {

	success, err := d.backend.CreateDirectory(req.DirectoryName)

	fmt.Println(req.DirectoryName)

	return &api.SuccessResponse{Success: success}, err
}

// AddEntry creates a new entry
func (d *Directory) AddEntry(ctx context.Context, req *api.EntryRequest) (*api.SuccessResponse, error) {

	success, err := d.backend.AddEntry(req)
	fmt.Println(req)
	return &api.SuccessResponse{Success: success}, err
}

// SearchEntity finds existing entities matching a query
func (d *Directory) SearchEntry(ctx context.Context, req *api.SearchEntryRequest) (*api.SearchEntriesResponse, error) {

	result, err := d.backend.SearchEntry(req.Query)

	if err != nil {
		return nil, err
	}
	resp := &api.SearchEntriesResponse{Entries: result}

	return resp, nil
}

func (d *Directory) Echo(ctx context.Context, req *api.EchoRequest) (*api.EchoResponse, error) {

	success := "ok"
	fmt.Println(req)
	return &api.EchoResponse{EchoSuccess: success}, nil
}

func (d *Directory) GetUser(ctx context.Context, req *api.GetUserRequest) (*api.GetUserResponse, error) {

	success, err := d.backend.GetUser(req.Command)

	fmt.Println(req.Command)

	return &api.GetUserResponse{Success: success}, err
}

// Cleanup
func (d *Directory) Close() error {
	return d.backend.Close()
}
