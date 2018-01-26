// Package borkbot provides a small server that will send borks to the Samantha and Christopher Slack
package borkbot

import (
	"errors"
)

// Service is the interface that provides the borkbot methods.
type Service interface {
	// FetchBork receives a request from the slack app and responds with a dog meme
	FetchBork(req fetchBorkRequest) (string, error)
}

type service struct {
	verficationToken string
}

func (s *service) FetchBork(req fetchBorkRequest) (string, error) {
	if s.verficationToken != req.Token {
		return "", errnotFromSlack
	}
	bork, err := s.borkGenerator()
	if err != nil {
		return "", err
	}
	return bork, nil
}

func (s *service) borkGenerator() (string, error) {
	borkURL := "https://barkpost.com/wp-content/uploads/2015/02/featmeme.jpg"
	return borkURL, nil
}

func (s *service) borkFetcher() {

}

// NewService creates a borkbot service
func NewService(token string) Service {
	return &service{
		verficationToken: token,
	}
}

var errnotFromSlack = errors.New("why you try to be a slack?")
