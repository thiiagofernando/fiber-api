package models

// Usuário representa um usuário no sistema
// @Description Modelo de usuário com campos GORM
// @Accept json
// @Produce json
type User struct {
	ID           uint   `json:"id"`           // @Description O identificador exclusivo do usuário
	UsuarioNome  string `json:"usuarioNome"`  // @Description Nome completo do usuário
	UsuarioLogin string `json:"usuarioLogin"` // @Description Login do usuário
	Senha        string `json:"senha"`        // @Description Senha do usuário
}
