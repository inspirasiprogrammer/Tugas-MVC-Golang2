package model

type Account struct {
	ID            int    `gorm:"primary_key" json:"-"`
	IdAccount     string `json:"id_account,omitempty"`
	Name          string `json:"name"`
	Password      string `json:"password,omitempty"`
	AccountNumber int    `json:"account_number,omitempty"`
	Saldo         int    `json:"sa1do"`
}

type Auth struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Transaction struct {
	ID                     int    `gorm:"primary_key" json:"-"`
	TransactionType        int    `json:"transaction_type,omitempty"`
	TransactionDescription string `json:"transaction_description"`
	Sender                 int    `json:"sender"`
	Amount                 int    `json:"amount"`
	Recipient              int    `json:"recipient"`
	Timestamp              int64  `json:"timestamp,omitempty"`
}
