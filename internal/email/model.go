package email

import (
	"net/mail"

	"github.com/DouglasMaffraValiate/VESTA/pck/err"
)

type Email string

func (e *Email) isValid() (bool, *err.Error) {
	emailStr := string(*e)
	err := err.NewValidationError(
		"EmailError001",
		"Email invalido. O email esta vazio ou tem formato inválido.",
		"",
		nil,
		map[string]string{"Email": emailStr},
	)

	if e == nil || len(*e) == 0 {
		return false, err
	}

	_, errr := mail.ParseAddress(emailStr)

	if errr != nil {
		err.Cause = errr
		return false, err
	}

	return true, nil
}
