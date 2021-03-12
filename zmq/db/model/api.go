package model

type DatabaseAPI interface {
	HasTable(interface{}) bool

	DropTable(interface{})

	Close() error
}
