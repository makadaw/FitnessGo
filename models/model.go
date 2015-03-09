package models

type AppModel interface {
	All() interface{}
	FindById(id interface{}) error
}
