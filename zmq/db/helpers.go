package db

import (
	"errors"
	"zmq/db/model"
	"zmq/db/mysql_driver"
)

var (
	InvalidDBType = errors.New("invalid db type")
)

func GetDatabase(dbUrl map[string]string) (model.DatabaseAPI, error) {
	switch dbUrl["Type"] {
	case "mysql":
		return mysql_driver.GetConnection(dbUrl), nil
	default:
		return nil, InvalidDBType
	}
}

func CreateDatabase(dbUrl map[string]string) error {
	switch dbUrl["Type"] {
	case "mysql":
		mysql_driver.New(dbUrl)
		return nil
	default:
		return InvalidDBType
	}
}

func PurgeDB(dbUrl map[string]string) {
	db, _ := GetDatabase(dbUrl)
	defer db.Close()
	if db.HasTable(&model.UserUI{}) {
		db.DropTable(&model.UserUI{})
	}
	if db.HasTable(&model.WhileList{}) {
		db.DropTable(&model.WhileList{})
	}
	if db.HasTable(&model.UserOption{}) {
		db.DropTable(&model.UserOption{})
	}
	if db.HasTable(&model.DataRetention{}) {
		db.DropTable(&model.DataRetention{})
	}
	if db.HasTable(&model.User{}) {
		db.DropTable(&model.User{})
	}
	// clear cache
	// rdb := GetRedisConnection()
	// defer rdb.Close()
	// rdb.FlushAll()
	// clear all extracted pcap data
	// os.RemoveAll(config.PcapExtractDir())
	// recreate pcapextract dir
	// if _, err := os.Stat(config.PcapExtractDir()); os.IsNotExist(err) {
	// 	os.Mkdir(config.PcapExtractDir(), os.ModePerm)
	// }
}

// func GetRedisConnection() *redis.Client {
// 	// connect redis
// 	redisClient := redis.NewClient(&redis.Options{
// 		Addr:     config.RedisUrl(), // use default Addr
// 		Password: "",                // no password set
// 		DB:       0,                 // use default DB
// 	})
// 	_, err := redisClient.Ping().Result()

// 	if err != nil {
// 		panic(err)
// 	}
// 	return redisClient
// }
