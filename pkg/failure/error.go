package failure

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

func IsSQLRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func IsFSNotFound(err error) bool {
	return err != nil && strings.Contains(err.Error(), "code = NotFound")
}
