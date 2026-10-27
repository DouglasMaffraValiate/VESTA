package err

type ErroResponseDTO struct {
	ID           string            `json:"id"`
	Instant      string            `json:"instant"`
	Codigo       string            `json:"codigo"`
	Tipo         string            `json:"tipo"`
	Mensagem     string            `json:"mensagem"`
	IDRequisicao string            `json:"id_requisicao,omitempty"`
	Atributos    map[string]string `json:"atributos,omitempty"`
}

func (e *Error) GetResponseDTO() *ErroResponseDTO {
	if e == nil {
		return nil
	}

	return &ErroResponseDTO{
		ID:           e.ID.String(),
		Instant:      e.Instant.Format("2006-01-02T15:04:05Z07:00"),
		Codigo:       e.Codigo,
		Tipo:         e.Tipo,
		Mensagem:     e.Mensagem,
		IDRequisicao: e.IDRequisicao,
		Atributos:    e.Atributos,
	}
}
