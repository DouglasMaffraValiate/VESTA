package iff

func IF[T any](condicao bool, seVerdadeiro, seFalso T) T {
	if condicao {
		return seVerdadeiro
	}
	return seFalso
}
