package service_customer

type CreateCustomerReq struct {
	CustomerName string  `json:"customer_name"`
	Amount       float64 `json:"amount"`
	TxId         string  `json:"tx_id"`
}

type GetCustomerData struct {
	TxId   string  `json:"tx_id"`
	Amount float64 `json:"amount"`
}
