package service

import "github.com/FakharzadehH/CloudComputing-Fall-1402/internal/repository"

type Service struct {
	repos *repository.Repository
}

func New(repos *repository.Repository) *Service {
	return &Service{
		repos: repos,
	}
}
