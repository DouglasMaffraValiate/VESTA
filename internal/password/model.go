package password

import (
	"fmt"
	"unicode/utf8"

	"github.com/DouglasMaffraValiate/VESTA/pkg/err"
	"github.com/DouglasMaffraValiate/VESTA/pkg/err/errCatalog"
)

type Password string

type passwordRules struct {
	tamanhoMinimoEmCaracteres int
	tamanhoMaximoEmCaracteres int
}

var defaultPasswordRules = passwordRules{
	tamanhoMinimoEmCaracteres: 16,
	tamanhoMaximoEmCaracteres: 128,
}

func (p Password) isValid() (bool, *err.Error) {
	pStr := string(p)

	if !utf8.ValidString(pStr) || pStr == "" {
		return false, err.NewValidationError(
			errCatalog.Password001.Code,
			errCatalog.Password001.Message,
			"",
			nil,
			map[string]string{"Senha": "inválida ou em branco"},
		)
	}

	tamanho := utf8.RuneCountInString(pStr)

	if tamanho < defaultPasswordRules.tamanhoMinimoEmCaracteres {
		return false, err.NewValidationError(
			errCatalog.Password002.Code,
			errCatalog.Password002.Message,
			"",
			nil,
			map[string]string{
				"Tamanho atual": fmt.Sprintf("%d", tamanho),
			},
		)
	}

	if tamanho > defaultPasswordRules.tamanhoMaximoEmCaracteres {
		return false, err.NewValidationError(
			errCatalog.Password003.Code,
			errCatalog.Password003.Message,
			"",
			nil,
			map[string]string{
				"Tamanho atual": fmt.Sprintf("%d", tamanho),
			},
		)
	}

	return true, nil
}
