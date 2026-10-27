package email

import (
	"net/mail"

	"github.com/DouglasMaffraValiate/VESTA/pkg/err"
	"github.com/DouglasMaffraValiate/VESTA/pkg/err/errCatalog"
)

type Email string

func (e *Email) isValid() (bool, *err.Error) {

	err := err.NewValidationError(
		errCatalog.Email001,
		"",
		nil,
		nil,
	)

	if e == nil || len(*e) == 0 {
		err.Atributos = map[string]string{"Email": ""}
		return false, err
	}

	emailStr := string(*e)
	_, errr := mail.ParseAddress(emailStr)

	if errr != nil {
		err.Atributos = map[string]string{"Email": emailStr}
		err.Cause = errr
		return false, err
	}

	return true, nil
}
