package postgres

import (
	"RESTovergRPC/backend"
	api "RESTovergRPC/directory"
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type storer struct {
	db *gorm.DB
}

func New(dburl map[string]string) backend.Backend {
	//url would look like "host=myhost port=myport user=gorm dbname=gorm password=mypassword"
	port, _ := strconv.Atoi(dburl["Port"])
	url := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		dburl["Host"], port, dburl["User"], dburl["Type"], dburl["Password"])
	fmt.Println(url)
	db, err := gorm.Open("postgres", url)

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Directory{}, &Entry{})
	return &storer{db: db}
}

// CreateDirectory
func (s *storer) CreateDirectory(name string) (string, error) {
	fmt.Println("CreateDirectory", name)
	s.db.Create(&Directory{DirectoryName: name})
	return "OK", nil
}

// AddEntry and return string "ok" or "fail"
func (s *storer) AddEntry(e *api.EntryRequest) (string, error) {
	s.db.Create(&Entry{DirectoryRefer: e.DirectoryName, Name: e.Entry.Name, LastName: e.Entry.LastName, PhNumber: e.Entry.PhNumber})
	return "OK", nil
}

// SearchEntry
func (s *storer) SearchEntry(query string) ([]*api.Entry, error) {
	var users []*api.Entry
	query_map := api.ExtractQuery(query)
	s.db.Where(query_map).Find(users)
	return users, nil
}

// Close handles any necessary cleanup
func (s *storer) Close() error {
	return s.db.Close()
}
