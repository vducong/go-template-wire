package reusablecode

import (
	"go-template-wire/pkg/logger"
)

type Service struct {
	log  *logger.Logger
	repo *Repo
}

func newService(log *logger.Logger, repo *Repo) *Service {
	return &Service{log, repo}
}
