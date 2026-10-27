package password

import (
	"github.com/DouglasMaffraValiate/VESTA/pkg/err"
	"github.com/DouglasMaffraValiate/VESTA/pkg/err/errCatalog"
	"golang.org/x/crypto/bcrypt"
)

func (p *Password) hashPassword() *err.Error {
	hash, er := bcrypt.GenerateFromPassword([]byte(*p), bcrypt.DefaultCost)

	if er != nil {
		return err.NewInternalError(
			errCatalog.Password004,
			"",
			er,
			nil,
		)
	}

	*p = Password(hash)
	return nil
}
