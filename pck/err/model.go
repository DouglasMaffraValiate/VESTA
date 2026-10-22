package err

import (
	"errors"
	"net/http"
	"time"
	"uuid"
)

type ErroType struct {
	typeName       string
	httpStatusCode int
}

type Error struct {
	// Identificação do erro
	ID      uuid.UUID `json:"id"`
	Instant time.Time `json:"instant"`
	Codigo  string    `json:"codigo"`
	Tipo    string    `json:"tipo"`

	// Atributos gerais
	Mensagem       string `json:"mensagem"`
	HttpStatusCode int    `json:"http_status_code"`

	// Rastreabilidade distribuída
	IDRequisicao string            `json:"id_requisicao,omitempty"`
	Atributos    map[string]string `json:"atributos,omitempty"`

	Cause error
}

func (e *Error) Error() string {
	if e == nil {
		return "<nil>"
	}
	return e.Mensagem
}

func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Cause
}

func (e *Error) Is(alvo error) bool {
	var err *Error
	ok := errors.As(alvo, &err)
	return ok && e.Codigo != "" && e.Codigo == err.Codigo
}

func newError(tipo ErroType, code, mensagem, idRequisicao string, cause error, atributos map[string]string) *Error {
	err := &Error{
		ID:             uuid.New(),
		Instant:        time.Now().UTC(),
		Codigo:         code,
		Tipo:           tipo.typeName,
		Mensagem:       mensagem,
		HttpStatusCode: tipo.httpStatusCode,
		Cause:          cause,
		Atributos:      make(map[string]string),
	}

	if idRequisicao != "" && len(idRequisicao) > 1 {
		err.IDRequisicao = idRequisicao
	}

	if len(atributos) > 0 {
		err.Atributos = atributos
	}

	return err
}

func NewValidationError(code, mensagem, idRequisicao string, cause error, atributos map[string]string) *Error {
	return newError(ErroType{typeName: "validation", httpStatusCode: http.StatusBadRequest}, code, mensagem, idRequisicao, cause, atributos)
}

func NewNotFoundError(code, mensagem, idRequisicao string, cause error, atributos map[string]string) *Error {
	return newError(ErroType{typeName: "not_found", httpStatusCode: http.StatusNotFound}, code, mensagem, idRequisicao, cause, atributos)
}

func NewConflictError(code, mensagem, idRequisicao string, cause error, atributos map[string]string) *Error {
	return newError(ErroType{typeName: "conflict", httpStatusCode: http.StatusConflict}, code, mensagem, idRequisicao, cause, atributos)
}

func NewForbiddenError(code, mensagem, idRequisicao string, cause error, atributos map[string]string) *Error {
	return newError(ErroType{typeName: "forbidden", httpStatusCode: http.StatusForbidden}, code, mensagem, idRequisicao, cause, atributos)
}

func NewUnauthenticatedError(code, mensagem, idRequisicao string, cause error, atributos map[string]string) *Error {
	return newError(ErroType{typeName: "unauthenticated", httpStatusCode: http.StatusUnauthorized}, code, mensagem, idRequisicao, cause, atributos)
}

func NewInternalError(code, mensagem, idRequisicao string, cause error, atributos map[string]string) *Error {
	return newError(ErroType{typeName: "internal", httpStatusCode: http.StatusInternalServerError}, code, mensagem, idRequisicao, cause, atributos)
}

func NewUnavailableError(code, mensagem, idRequisicao string, cause error, atributos map[string]string) *Error {
	return newError(ErroType{typeName: "unavailable", httpStatusCode: http.StatusServiceUnavailable}, code, mensagem, idRequisicao, cause, atributos)
}

func NewTimeoutError(code, mensagem, idRequisicao string, cause error, atributos map[string]string) *Error {
	return newError(ErroType{typeName: "timeout", httpStatusCode: http.StatusRequestTimeout}, code, mensagem, idRequisicao, cause, atributos)
}

func NewRateLimitedError(code, mensagem, idRequisicao string, cause error, atributos map[string]string) *Error {
	return newError(ErroType{typeName: "rate_limited", httpStatusCode: http.StatusTooManyRequests}, code, mensagem, idRequisicao, cause, atributos)
}
