package models

import "time"

type Payload struct {
	Event Event `json:"event"`
}

type Event struct {
	Created      time.Time `json:"created"`
	ID           string    `json:"id"`
	Log          Log       `json:"log"`
	Subscription string    `json:"subscription"`
	WorkspaceID  string    `json:"workspaceId"`
}

type Log struct {
	Authentication string    `json:"authentication"`
	Created        time.Time `json:"created"`
	Errors         []string  `json:"errors"`
	ID             string    `json:"id"`
	Invoice        Invoice   `json:"invoice"`
	Type           string    `json:"type"`
}

type Invoice struct {
	Amount             int64         `json:"amount"`
	Brcode             string        `json:"brcode"`
	Created            time.Time     `json:"created"`
	Descriptions       []KeyValue    `json:"descriptions"`
	DiscountAmount     int64         `json:"discountAmount"`
	Discounts          []Discount    `json:"discounts"`
	DisplayDescription string        `json:"displayDescription"`
	Due                time.Time     `json:"due"`
	Expiration         int64         `json:"expiration"`
	Fee                int64         `json:"fee"`
	Fine               float64       `json:"fine"`
	FineAmount         int64         `json:"fineAmount"`
	ID                 string        `json:"id"`
	Interest           float64       `json:"interest"`
	InterestAmount     int64         `json:"interestAmount"`
	Link               string        `json:"link"`
	Name               string        `json:"name"`
	NominalAmount      int64         `json:"nominalAmount"`
	PDF                string        `json:"pdf"`
	Rules              []interface{} `json:"rules"`  // Ajuste se os dados de "rules" tiverem um tipo específico
	Splits             []interface{} `json:"splits"` // Ajuste se os dados de "splits" tiverem um tipo específico
	Status             string        `json:"status"`
	Tags               []string      `json:"tags"`
	TaxID              string        `json:"taxId"`
	TransactionIDs     []string      `json:"transactionIds"`
	Updated            time.Time     `json:"updated"`
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Discount struct {
	Due        time.Time `json:"due"`
	Percentage int       `json:"percentage"`
}
