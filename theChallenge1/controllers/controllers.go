package controllers

import (
	"challenge-mauricio-mendonca-starkbank/theChallenge1/utils"
	"fmt"
	"math/rand"
	"time"

	"github.com/starkbank/sdk-go/starkbank/invoice"
)

func postInvoice() {

	//função que gera a Invoice no stark Bank
	due := time.Now().Add(24 * time.Hour)
	pessoa, err := utils.GetPessoa() //Consulta dados de pessoas aleatorias
	if err != nil {
		fmt.Println("Falha ao gerar pessoas:", err)
		return
	}

	descriptions := make([]map[string]interface{}, 1)
	var data = map[string]interface{}{}
	data["key"] = "Teste"
	data["value"] = "Mauricio Mendonca"
	descriptions[0] = data

	discounts := make([]map[string]interface{}, 1)
	var dataDiscount = map[string]interface{}{}
	dataDiscount["percentage"] = 1
	dataDiscount["due"] = due
	discounts[0] = dataDiscount

	if len(pessoa) > 0 {
		invoices, err := invoice.Create( //criacao da invoice
			[]invoice.Invoice{
				{
					Amount:       rand.Intn(520000-20000+1) + 20000, //funcao para gerar um numero inteiro aleatorio entre 1 e 520000
					Name:         pessoa[0].Nome,
					TaxId:        pessoa[0].Cpf,
					Descriptions: descriptions,
					Discounts:    discounts,
					Due:          &due,
					Fine:         2.5,
					Interest:     1.3,
					Expiration:   int(time.Now().Add(24 * time.Hour).Unix()),
					Tags:         []string{"War supply", "Invoice #1234"},
				},
			}, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
			}
		}

		for _, invoice := range invoices {
			fmt.Printf("%+v", invoice)
		}
	}

}

func GenerateInvoiceAwaite() {
	// Inicializa o gerador de números aleatórios com a semente baseada no tempo atual
	rand.Seed(time.Now().UnixNano())

	// Define o tempo de início
	startTime := time.Now()

	// Loop que continua por até 24 horas
	for time.Since(startTime) < 24*time.Hour {
		//valor randomico entre 8 e 12
		numCount := rand.Intn(5) + 8
		fmt.Print(numCount)
		for i := 0; i < numCount; i++ {
			postInvoice()
		}
		//sleep por 3h
		time.Sleep(3 * time.Hour)
	}
}
