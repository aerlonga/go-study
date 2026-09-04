# fase-1 — exemplo de estrutura de projeto Go

Baseado no padrão `cmd/` + `internal/`, usado como referência pra montar `novo-fase-1`.

## Rodar

```
go run ./cmd/students-api
```

## Testar

```
curl -X POST localhost:8080/students -d '{"name":"Ana","email":"ana@x.com","age":20}'
curl "localhost:8080/students?id=1"
```

## Mapa de pastas

- `cmd/students-api/main.go` — ponto de entrada, monta tudo e sobe o servidor.
- `config/` — arquivos yaml de configuração.
- `internal/config/` — struct Go que representa a config.
- `internal/models/` — entidades do domínio (equivalente a Models do Laravel).
- `internal/storage/` — interface `Storage` + implementação em memória. Pra trocar por banco real, criar novo arquivo implementando a mesma interface.
- `internal/http/handlers/` — handlers HTTP (equivalente a Controllers).
- `internal/utils/response/` — helpers, ex: escrever resposta JSON padronizada.
