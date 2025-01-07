package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"theChallenge2/models"

	Transfer "github.com/starkbank/sdk-go/starkbank/transfer"
)

func GenerateTransfer(w http.ResponseWriter, r *http.Request) {

	var payload models.Payload
	json.NewDecoder(r.Body).Decode(&payload)

	if payload.Event.Log.Type == "credited" {

		transfers, err := Transfer.Create(
			[]Transfer.Transfer{
				{
					Amount:        int(payload.Event.Log.Invoice.Amount),
					Name:          "Stark Bank S.A.",
					TaxId:         "20.018.183/0001-80",
					BankCode:      "20018183", //Pix
					BranchCode:    "0001",
					AccountNumber: "6341320293482496",
					AccountType:   "payment",
				},
			}, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
			}
		}

		for _, transfer := range transfers {
			fmt.Println(transfer)
		}
	}

}
