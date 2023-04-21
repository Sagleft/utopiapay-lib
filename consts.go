package utopiapaylib

const apiHost = "https://utopiapay.io"

type CurrencyType string

const (
	CurrencyCrypton CurrencyType = "CRP"
	CurrencyUUSD    CurrencyType = "UUSD"
)

// scalarSize is the size of the scalar input to X25519
const scalarSize = 32
