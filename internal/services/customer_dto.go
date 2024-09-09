package service_customer

type CreateCustomerReq struct {
	CustomerName string `json:"customer_name"`
	TxId         string `json:"tx_id"`
}
