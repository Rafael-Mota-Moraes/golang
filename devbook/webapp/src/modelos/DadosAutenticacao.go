package modelos

// DadosAutenticacao contém o id e o token do usuário logado
type DadosAutenticacao struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
