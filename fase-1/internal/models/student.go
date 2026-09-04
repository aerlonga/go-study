// Pacote "models": guarda as entidades do domínio (o que é um "Student" no sistema).
package models

// Student é a entidade principal do domínio.
// Struct pública (S maiúsculo), assim outros pacotes conseguem usar esse tipo.
type Student struct {
	// Cada linha: NomeDoCampo Tipo `tag`
	// Tag json:"..." diz como esse campo aparece quando a struct vira JSON
	// (ex: response.JSON manda isso pro navegador/cliente).
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}
