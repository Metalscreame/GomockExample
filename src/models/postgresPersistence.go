package models

type PostgresService struct {
}

func (s *PostgresService) Save(u User) error {
	//...saving to the database
	return nil
}

func (s *PostgresService) Delete(u User) error {
	//...deleting from the database
	return nil
}
