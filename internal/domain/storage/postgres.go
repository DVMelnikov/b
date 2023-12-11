package storage

import "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"

type Postgres struct{}

func NewPostgres() *Postgres {
	return &Postgres{}
}

func (p *Postgres) GetAllFeed() ([]domain.Feed, error) {
	return []domain.Feed{
		{Id: 1, Title: "Title1", Description: "Description1"},
		{Id: 2, Title: "Title2", Description: "Description2"},
		{Id: 3, Title: "Title3", Description: "Description3"},
	}, nil
}

func (p *Postgres) GetFeed(id int) (domain.Feed, error) {
	if id == 1 {
		return domain.Feed{Id: 1, Title: "Title1", Description: "Description1", RegistationURL: "URL1"}, nil
	} else if id == 2 {
		return domain.Feed{Id: 2, Title: "Title2", Description: "Description2", RegistationURL: "URL2"}, nil
	} else if id == 3 {
		return domain.Feed{Id: 3, Title: "Title3", Description: "Description3", RegistationURL: "URL3"}, nil
	}
	return domain.Feed{}, nil
}
