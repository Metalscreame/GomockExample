package models

//go:generate mockgen -destination=../mocks/mock_UserPersistence.go -package=mocks GomockExample/src/models UserPersistence

type User struct {
	Name string
}

type UserPersistence interface {
	Save(u User) error
	Delete(u User) error
}
