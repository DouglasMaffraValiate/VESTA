package password

import (
	"errors"

	"github.com/DouglasMaffraValiate/VESTA/pkg/err"
	"github.com/DouglasMaffraValiate/VESTA/pkg/err/errCatalog"
	"golang.org/x/crypto/bcrypt"
)

func (p *Password) ComparePassword(inputPassword string) (bool, *err.Error) {
	error := bcrypt.CompareHashAndPassword(
		[]byte(*p),
		[]byte(inputPassword),
	)

	if error != nil {
		if errors.Is(error, bcrypt.ErrMismatchedHashAndPassword) {
			return false, err.NewUnauthenticatedError(
				errCatalog.Auth001,
				"",
				error,
				nil,
			)
		}
		return false, err.NewInternalError(
			errCatalog.Auth002,
			"",
			error,
			nil,
		)
	}

	return true, nil
}
