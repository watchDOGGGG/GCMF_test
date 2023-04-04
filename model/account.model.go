package model

type VerifyAccountModel struct {
	CustomerName  string `json:"customerName"`
	AccountNumber string `json:"accountNumber"`
	AccountName   string `json:"accountName"`
	BankName      string `json:"bankName"`
	Amount        string `json:"amount"`
	BankCode      string `json:"bankCode"`
	BankType      string `json:"banktype"`
	AccountType   string `json:"accountType"`
}
