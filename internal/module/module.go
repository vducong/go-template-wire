package module

import (
	reusablecode "go-template-wire/internal/reusable_code"
	"go-template-wire/pkg/user"

	"github.com/google/wire"
)

var ModuleSet = wire.NewSet(
	reusablecode.NewModule, user.New,
)
