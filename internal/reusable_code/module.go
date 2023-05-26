package reusablecode

import (
	"go-template-wire/pkg/databases"
	"go-template-wire/pkg/logger"
)

type Module struct {
	Repo    *Repo
	Service *Service
}

func NewModule(log *logger.Logger, db databases.MySQLDB) *Module {
	repo := newRepo(db)
	s := newService(log, repo)
	return &Module{repo, s}
}
