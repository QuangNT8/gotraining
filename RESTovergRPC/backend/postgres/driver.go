package postgres

import (
	"RESTovergRPC/backend"
	api "RESTovergRPC/directory"
	"encoding/json"
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
	db.AutoMigrate(&Directory{}, &Entry{}, &UserType{})
	return &storer{db: db}
}

// CreateDirectory
func (s *storer) CreateDirectory(name string) (string, error) {
	fmt.Println("CreateDirectory", name)
	s.db.Create(&Directory{DirectoryName: name})
	return "OK", nil
}

type Test struct {
	directory_name string
	entry          string
}

func isJSON(s string) bool {
	fmt.Println("isJSON", s)

	var test Test
	json.Unmarshal([]byte(s), &test)
	fmt.Printf("Species: %s, Description: %s", test.directory_name, test.entry)

	var js map[string]interface{}
	fmt.Println("js map", js)
	return json.Unmarshal([]byte(s), &js) == nil
}

// AddEntry and return string "ok" or "fail"
func (s *storer) AddEntry(e *api.EntryRequest) (string, error) {

	// if e.Entry == nil {
	// 	return "Missing Entry Input", nil
	// }

	if e.DirectoryName == "" || e.Entry == nil {
		return "wrong format", nil
	}

	fmt.Printf("Entry Name: %s \r\n", e.Entry.Name)
	fmt.Printf("Entry LastName: %s \r\n", e.Entry.LastName)
	fmt.Printf("Entry PhNumber: %s \r\n", e.Entry.PhNumber)

	s.db.Create(&UserType{
		Name:        e.Entry.Name,
		LastName:    e.Entry.LastName,
		PhoneNumber: e.Entry.PhNumber})

	return "OK", nil
}

func (s *storer) GetUser(cmd string) (string, error) {
	fmt.Printf("command : %s", cmd)
	var userread UserType

	if cmd == "First" {
		s.db.First(&userread)
		fmt.Printf("User Data: %s \r\n", userread.PhoneNumber)
		fmt.Println(userread)
		return "OK", nil
	}
	return "Wrong Command", nil
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
