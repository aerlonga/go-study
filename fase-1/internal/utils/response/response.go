package response

import (
	"encoding/json" // converte structs Go <-> JSON
	"net/http"      // pacote padrão do Go pra servidor HTTP, requests, responses etc
)

// JSON escreve qualquer struct como resposta JSON, evitando repetir isso em cada handler.
//
// Parâmetros:
//   w http.ResponseWriter -> é o "objeto resposta" que o Go te dá em todo handler HTTP.
//                             Escrever nele = mandar dado pro navegador/cliente que fez a requisição.
//   status int            -> código HTTP (200, 201, 404 etc)
//   data any               -> "any" aceita QUALQUER tipo (equivalente a "interface{}").
//                             Assim essa função serve pra mandar qualquer struct como JSON.
func JSON(w http.ResponseWriter, status int, data any) {
	// define o cabeçalho HTTP dizendo "o corpo dessa resposta é JSON"
	w.Header().Set("Content-Type", "application/json")

	// define o código de status HTTP (precisa vir ANTES de escrever o corpo)
	w.WriteHeader(status)

	// json.NewEncoder(w) cria um "codificador" que escreve direto no w (no fluxo de resposta).
	// .Encode(data) transforma "data" (a struct) em JSON e já manda pro cliente.
	json.NewEncoder(w).Encode(data)
}

// Error é um atalho pra devolver erro em formato JSON padronizado: {"error": "mensagem"}
func Error(w http.ResponseWriter, status int, message string) {
	// map[string]string{...} = dicionário literal, aqui só com uma chave "error"
	JSON(w, status, map[string]string{"error": message})
}
