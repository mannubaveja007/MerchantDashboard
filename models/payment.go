package models

type Payment struct {
	TransactionID string  `json:"transaction_id"`
	Amount        float64 `json:"amount"`
	Status        string  `json:"status"`
}
type TransferRequest struct {
	Amount      float64 `json:"amount"`
	BankAccount string  `json:"bank_account"`
}

type ReceiveRequest struct {
	Amount float64 `json:"amount"`
}

type Transaction struct {
	FromAddress string  `json:"from_address"`
	ToAddress   string  `json:"to_address"`
	Amount      float64 `json:"amount"`
}
