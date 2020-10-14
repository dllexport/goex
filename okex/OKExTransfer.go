package okex

import (
	"fmt"
)

type OKExTransfer struct {
	*OKEx
}

type TransferResponse struct {
	TransferID string `json:"transfer_id"`
	Currency   string `json:"currency"`
	From       string `json:"from"`
	Amount     string `json:"amount"`
	To         string `json:"to"`
	Result     bool   `json:"result"`
}

func (ok *OKExTransfer) Transfer(currency string, amount float64, from string, to string) (*TransferResponse, error) {
	urlPath := "/api/account/v3/transfer"
	var response TransferResponse

	var param struct {
		Currency string `json:"currency"`
		Amount   string `json:"amount"`
		From     string `json:"from"`
		To       string `json:"to"`
	}

	param.Currency = currency
	param.Amount = fmt.Sprintf("%f", amount)
	param.From = from
	param.To = to
	reqBody, _, _ := ok.BuildRequestBody(param)
	err := ok.OKEx.DoRequest("POST", urlPath, reqBody, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
