package models

type Pessoa struct {
	Nome          string `json:"nome"`
	Idade         int    `json:"idade"`
	Cpf           string `json:"cpf"`
	Rg            string `json:"rg"`
	DataNasc      string `json:"data_nasc"`
	Sexo          string `json:"sexo"`
	Signo         string `json:"signo"`
	Mae           string `json:"mae"`
	Pai           string `json:"pai"`
	Email         string `json:"email"`
	Senha         string `json:"senha"`
	Cep           string `json:"cep"`
	Endereco      string `json:"endereco"`
	Numero        int    `json:"numero"`
	Bairro        string `json:"bairro"`
	Cidade        string `json:"cidade"`
	Estado        string `json:"estado"`
	TelefoneFixo  string `json:"telefone_fixo"`
	Celular       string `json:"celular"`
	Altura        string `json:"altura"`
	Peso          int    `json:"peso"`
	TipoSanguineo string `json:"tipo_sanguineo"`
	Cor           string `json:"cor"`
}
