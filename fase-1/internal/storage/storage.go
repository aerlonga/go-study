package storage

// Importa o pacote "models" do nosso próprio módulo.
// Caminho de import = nome do módulo (fase-1, definido no go.mod) + caminho da pasta.
import "fase-1/internal/models"

// interface = "contrato". Não implementa nada, só diz quais métodos
// algo precisa ter pra ser considerado um "Storage".
// Storage define o que qualquer implementação de persistência precisa oferecer.
// Handlers dependem dessa interface, não de um banco específico.
type Storage interface {
	// Qualquer tipo que tenha ESSES métodos, com essas assinaturas exatas,
	// automaticamente "satisfaz" a interface Storage (sem precisar declarar isso em lugar nenhum,
	// diferente de outras linguagens que usam "implements").
	CreateStudent(name, email string, age int) (int64, error)
	GetStudentByID(id int64) (models.Student, error)

	// []models.Student = slice (lista) de Student. Equivale a array dinâmico.
	GetStudents() ([]models.Student, error)

	UpdateStudent(id int64, name, email string, age int) error
	DeleteStudent(id int64) error
}

// Por que interface? Assim o handler (arquivo handler.go) recebe "Storage" genérico.
// Hoje é memória (memory.go), amanhã pode ser SQLite/Postgres, sem mudar o handler.
