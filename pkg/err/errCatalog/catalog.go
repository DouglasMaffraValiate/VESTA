package errCatalog

type ErrorCatalog struct {
	Code, Message string
}

var (
	//Password
	Password001 = ErrorCatalog{Code: "Password-001", Message: "Senha inválida. A senha não pode estar vazia ou conter texto inválido."}
	Password002 = ErrorCatalog{Code: "Password-002", Message: "Senha inválida. A senha deve ter pelo menos 16 caracteres."}
	Password003 = ErrorCatalog{Code: "Password-003", Message: "Senha inválida. A senha não pode ultrapassar 128 caracteres."}
	Password004 = ErrorCatalog{Code: "Password-004", Message: "Erro inesperado ao gerar hash da senha."}

	// Auth
	Auth001 = ErrorCatalog{Code: "Auth-001", Message: "Credenciais inválidas"}
	Auth002 = ErrorCatalog{Code: "Auth-002", Message: "Erro interno no processo de autenticação."}

	// Role
	Role001 = ErrorCatalog{Code: "Role-001", Message: "Role informada é inexistente no sistema."}

	// Email
	Email001 = ErrorCatalog{Code: "Email-001", Message: "Email invalido. O email esta vazio ou tem formato inválido."}
)
