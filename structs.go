package utopiapaylib

type Invoice struct {
	ID                     string       `json:"id"`
	OrderID                string       `json:"orderID"` // random hex string by default
	Currency               CurrencyType `json:"currency"`
	Amount                 float64      `json:"amount"`
	Token                  string       `json:"token"`
	RedirectToURLOnSuccess string       `json:"webTo"`
}
