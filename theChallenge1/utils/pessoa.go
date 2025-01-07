package utils

import (
	"bytes"
	"challenge-mauricio-mendonca-starkbank/theChallenge1/models"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

func GetPessoa() ([]models.Pessoa, error) {
	url := "https://www.4devs.com.br/ferramentas_online.php"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("acao", "gerar_pessoa")
	_ = writer.WriteField("sexo", "I")
	_ = writer.WriteField("pontuacao", "5")
	_ = writer.WriteField("idade", "0")
	_ = writer.WriteField("cep_estado", "")
	_ = writer.WriteField("txt_qtde", "1")
	_ = writer.WriteField("cep_cidade", "")
	writer.Close()

	client := &http.Client{}
	req, _ := http.NewRequest(method, url, payload)

	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro na solicitação HTTP: %v", err)
	}
	defer res.Body.Close()

	bodyBytes, _ := io.ReadAll(res.Body)

	// Tentamos primeiro com um array de Pessoas
	var pessoas []models.Pessoa
	if err := json.Unmarshal(bodyBytes, &pessoas); err != nil {
		// Se falhar, tentamos com uma única Pessoa
		var pessoa models.Pessoa
		pessoas = append(pessoas, pessoa)
	}

	return pessoas, nil
}
