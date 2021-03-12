package mysql_driver

import (
	"errors"
	"fmt"
	"log"
	"zmq/db/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	InvalidDBType = errors.New("invalid db type")
)

type database struct {
	db *gorm.DB
}

// Get database
func (s *database) GetDatabase(dbUrl map[string]string) (model.DatabaseAPI, error) {
	mysqlDB, err := open(dbUrl)
	s.db = mysqlDB

	if err != nil {
		log.Fatalln("failed to connect database")
	}

	switch dbUrl["Type"] {
	case "mysql":
		return nil, nil
	default:
		return nil, InvalidDBType
	}
}

func GetConnection(dbUrl map[string]string) model.DatabaseAPI {
	mysqlDB, err := open(dbUrl)

	if err != nil {
		log.Fatalln("failed to connect database")
	}

	return &database{db: mysqlDB}
}

func open(dburl map[string]string) (*gorm.DB, error) {
	// url would look like user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	// fmt.Println("==========open=======")
	// fmt.Println("=====================")
	// fmt.Println("Host: ", dburl["Host"])
	// fmt.Println("Name: ", dburl["Name"])
	// fmt.Println("User: ", dburl["User"])
	// fmt.Println("Password: ", dburl["Password"])

	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4,utf8&parseTime=True&loc=Local", dburl["User"], dburl["Password"], dburl["Host"], dburl["Name"])
	// fmt.Println(url)
	return gorm.Open("mysql", url)
}

func CreateDatabase(dbUrl map[string]string) error {
	switch dbUrl["Type"] {
	case "mysql":
		New(dbUrl)
		return nil
	default:
		return InvalidDBType
	}
}

func New(dbUrl map[string]string) {
	mysqlDB, err := open(dbUrl)

	if err != nil {
		panic("failed to connect database")
	}
	// create or update table
	mysqlDB.Set("gorm:table_options", "ENGINE=InnoDB COLLATE=utf8mb4_unicode_ci").AutoMigrate(
		&model.User{},
		&model.UserOption{},
		&model.WhileList{},
		&model.UserUI{},
		&model.DataRetention{},
	)

	// mysqlDB.Model(&model.UserOption{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	// mysqlDB.Model(&model.WhileList{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

}

func (s *database) HasTable(table interface{}) bool {
	return s.db.HasTable(table)
}

func (s *database) DropTable(table interface{}) {
	s.db.DropTable(table)
}

// Close handles any necessary cleanup
func (s *database) Close() error {
	return s.db.Close()
}
