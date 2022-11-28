package greetings

import "github.com/AyJayTee/greetings/internal/repository"

type Service struct {
	db repository.Repository
}

func NewService(r repository.Repository) *Service {
	return &Service{
		db: r,
	}
}

func (s *Service) Add(url string) (string, error) {
	return s.db.AddUrl(url)
}

func (s *Service) Fetch(id string) (string, error) {
	return s.db.FetchUrl(id)
}
