// Pacote "student": handlers HTTP relacionados a Student.
// Cada "recurso" (student, teacher, etc) ganharia sua própria pasta/pacote aqui dentro,
// igual controllers separados por recurso no Laravel.
package student

import (
	"encoding/json"
	"net/http"
	"strconv"

	"fase-1/internal/storage"
	"fase-1/internal/utils/response"
)

// struct privada (minúscula) pra representar o JSON de entrada (POST/PUT).
type studentRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// New cuida do POST /api/students — cria um student novo.
// Recebe a Storage por parâmetro (injeção de dependência), não importa qual implementação é.
func New(store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req studentRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			response.Error(w, http.StatusBadRequest, "corpo da requisição inválido")
			return
		}

		id, err := store.CreateStudent(req.Name, req.Email, req.Age)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		response.JSON(w, http.StatusCreated, map[string]int64{"id": id})
	}
}

// GetStudents cuida do GET /api/students — lista todos.
func GetStudents(store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		students, err := store.GetStudents()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, students)
	}
}

// GetStudentById cuida do GET /api/students/{id} — busca um específico.
func GetStudentById(store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// r.PathValue("id") lê o pedaço "{id}" definido no padrão da rota
		// (registrado em main.go como "GET /api/students/{id}").
		// Isso é recurso nativo do Go 1.22+, antes só framework tinha.
		idStr := r.PathValue("id")

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "id inválido")
			return
		}

		stud, err := store.GetStudentByID(id)
		if err != nil {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, stud)
	}
}

// UpdateStudent cuida do PUT /api/students — atualiza um existente.
// Aqui o id vem no corpo do JSON (por isso studentRequest ganhou campo ID abaixo),
// diferente do repo original que às vezes usa {id} na URL também — os dois jeitos existem,
// escolhi corpo pra variar e você ver as duas formas.
type updateRequest struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func UpdateStudent(store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req updateRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			response.Error(w, http.StatusBadRequest, "corpo da requisição inválido")
			return
		}

		if err := store.UpdateStudent(req.ID, req.Name, req.Email, req.Age); err != nil {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, map[string]string{"status": "atualizado"})
	}
}

// DeleteStudent cuida do DELETE /api/students/{id} — remove um existente.
func DeleteStudent(store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "id inválido")
			return
		}

		if err := store.DeleteStudent(id); err != nil {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, map[string]string{"status": "removido"})
	}
}
