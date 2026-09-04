package storage

import (
	"fmt"

	"fase-1/internal/models"
)

// MemoryStorage é uma implementação simples em memória, só pra exemplo/testes.
// Uma implementação real (ex: SQLite) ficaria em outro arquivo aqui dentro,
// tipo sqlite.go, implementando a mesma interface Storage.
type MemoryStorage struct {
	students map[int64]models.Student
	nextID   int64
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		students: make(map[int64]models.Student),
		nextID:   1,
	}
}

func (s *MemoryStorage) CreateStudent(name, email string, age int) (int64, error) {
	id := s.nextID
	s.students[id] = models.Student{ID: id, Name: name, Email: email, Age: age}
	s.nextID++
	return id, nil
}

func (s *MemoryStorage) GetStudentByID(id int64) (models.Student, error) {
	student, ok := s.students[id]
	if !ok {
		return models.Student{}, fmt.Errorf("student com id %d não encontrado", id)
	}
	return student, nil
}

// GetStudents devolve todos os students cadastrados, em formato de lista (slice).
func (s *MemoryStorage) GetStudents() ([]models.Student, error) {
	// make([]models.Student, 0, len(s.students)):
	// cria slice vazio já reservando espaço pra "len(s.students)" itens,
	// só otimização (evita realocar memória a cada append).
	result := make([]models.Student, 0, len(s.students))

	// "for range" percorre o map. "_" descarta a chave (id), só queremos o valor.
	for _, student := range s.students {
		result = append(result, student) // append adiciona item no fim do slice
	}

	return result, nil
}

// UpdateStudent sobrescreve os dados de um student já existente.
func (s *MemoryStorage) UpdateStudent(id int64, name, email string, age int) error {
	// confere se existe antes de sobrescrever (senão criaria um novo id "do nada")
	if _, ok := s.students[id]; !ok {
		return fmt.Errorf("student com id %d não encontrado", id)
	}

	s.students[id] = models.Student{ID: id, Name: name, Email: email, Age: age}
	return nil
}

// DeleteStudent remove um student do map.
func (s *MemoryStorage) DeleteStudent(id int64) error {
	if _, ok := s.students[id]; !ok {
		return fmt.Errorf("student com id %d não encontrado", id)
	}

	// delete() é função embutida do Go pra remover chave de um map.
	delete(s.students, id)
	return nil
}
