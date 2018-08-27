package models

type User struct {
	Name string
}

type UserPersistence interface {
	Save(u User) error
	Delete(u User) error
}
