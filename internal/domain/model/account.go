package model

type Account struct {
	Address    []byte
	NameTag    string
	Balance    float64
	Percentage float64
	TxCount    int
}
