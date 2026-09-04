// package config: nome do pacote. Todo arquivo .go pertence a um pacote.
// Arquivos na mesma pasta = mesmo pacote, compartilham código entre si sem precisar importar.
package config

// struct = como uma "classe" só de dados (sem métodos obrigatórios).
// Config representa as configurações lidas do arquivo yaml em config/.
type Config struct {
	// campo Env, tipo string. A crase `yaml:"env"` é uma "tag":
	// diz pra biblioteca de yaml que esse campo corresponde à chave "env" no arquivo.
	Env string `yaml:"env"`

	// HTTPServer é outra struct (definida embaixo), aninhada dentro de Config.
	HTTPServer HTTPServer `yaml:"http_server"`
}

// Segunda struct, separada, representando o bloco "http_server:" do yaml.
type HTTPServer struct {
	Address string `yaml:"address"`
}

// Função exportada (começa com letra maiúscula = pública, visível fora do pacote).
// MustLoad devolve uma config fixa por enquanto.
// Depois pode ler de config/local.yaml usando algum pacote de yaml.
func MustLoad() *Config {
	// "*Config" no retorno da função acima = ponteiro pra Config.
	// Ponteiro = endereço de memória, evita copiar a struct inteira toda vez que ela circula pelo código.

	// return &Config{...} monta a struct e devolve o endereço dela (& = "endereço de").
	return &Config{
		Env: "local",
		HTTPServer: HTTPServer{
			Address: "localhost:8080",
		},
	}
}
