package helper

import (
	"errors"

	"gorm.io/gorm"
)

func CheckQueryError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("error record not found")
	} else if errors.Is(err, gorm.ErrPreloadNotAllowed) {
		return errors.New("error preload not allowed")
	}
	return err
}
