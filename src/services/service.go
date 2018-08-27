package services

import (
	"GomockExample/src/models"
	"fmt"
)

type Service struct {
	Persistence models.UserPersistence
}

func NewService(p models.UserPersistence) Service {
	return Service{
		Persistence: p,
	}
}

func (s *Service) Saver(u models.User) error {
	err := s.Persistence.Save(u)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
