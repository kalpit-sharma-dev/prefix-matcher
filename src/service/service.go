package service

import (
	"errors"
	"prefix-matcher/src/repository"
)

type Service struct {
	repository repository.IRepository
}

func (s *Service) SearchWord(word string) (string, error) {
	count := s.repository.SearchWord(word)
	outputPrefix := ""
	for i := 0; i < count; i++ {
		outputPrefix = outputPrefix + string(word[i])
	}
	if len(outputPrefix) > 0 {
		return outputPrefix, nil
	}
	return outputPrefix, errors.New("no matching prefix found")
}
func NewService(repository repository.IRepository) *Service {
	return &Service{repository: repository}
}
