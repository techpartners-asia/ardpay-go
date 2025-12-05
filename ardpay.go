package ardpay

import (
	"errors"
	"strconv"

	"resty.dev/v3"
)

type Ardpay interface {
	CreateQrPayment(input CreateQrInput) (*CreateQrPaymentResponse, error)     // QR төлбөр үүсгэх
	CheckQrPayment(input CheckQrPaymentInput) (*CheckQrPaymentResponse, error) // QR төлбөр шалгах
	CheckQr(input CheckQrInput) (*CheckQrResponse, error)                      // QR төлбөр төлөгдсөн эсэхийг шалгах
	CancelQrPayment(input CancelQrInput) (*CancelQrPaymentResponse, error)     // QR төлбөр цуцлах
	TanPayment(input TanInput) error                                           // ТАНтай худалдан авалт
}

type ardpay struct {
	Url        string
	MerchantID string
	PosNo      string
	APIKey     string
}

func New(url, merchantID, posNo, apiKey string) Ardpay {
	return &ardpay{
		Url:        url,
		MerchantID: merchantID,
		PosNo:      posNo,
		APIKey:     apiKey,
	}
}

type CreateQrInput struct {
	Amount      float64 `json:"tranAmnt"`
	Currency    string  `json:"tranCur"`
	Description string  `json:"tranDesc"`
	InvoiceID   string  `json:"invoiceId"`
	PaidLimit   float64 `json:"paidLimit"`
}

// QR төлбөр үүсгэх
func (a *ardpay) CreateQrPayment(input CreateQrInput) (*CreateQrPaymentResponse, error) {
	body := CreateQrPaymentRequest{
		PayeeCode:              a.MerchantID,
		PosNo:                  a.PosNo,
		TransactionAmount:      input.Amount,
		InvoiceID:              input.InvoiceID,
		TransactionDescription: input.Description,
		TransactionCurrency:    input.Currency,
		PaidLimit:              input.PaidLimit,
	}
	client := resty.New()
	defer client.Close()
	var response CreateQrPaymentResponse
	res, err := client.R().
		SetHeader("APIKEY", a.APIKey).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&response).
		Post(a.Url + "/resources/merch/v1.0/createqr")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, errors.New(res.String())
	}
	if response.ResponseCode != 0 {
		return nil, errors.New("Error code: " + strconv.Itoa(response.ResponseCode) + " Error description: " + response.ResponseDesc)
	}
	return &response, nil
}

type GetQrDataInput struct {
	Type         string `json:"type"`
	CustomerCode string `json:"customerCode"`
	BankCode     string `json:"bankCode"`
	BankVerfCode string `json:"bankVerfCode"`
	BankTxCode   string `json:"bankTxCode"`
	QrCode       string `json:"qrCode"`
}

func (a *ardpay) GetQrData(input GetQrDataInput) (*GetQrDataResponse, error) {
	body := GetQrDataRequest{
		PayeeCode:    a.MerchantID,
		PosNo:        a.PosNo,
		Type:         input.Type,
		CustomerCode: input.CustomerCode,
		BankCode:     input.BankCode,
		BankVerfCode: input.BankVerfCode,
		JsonData: JsonData{
			BankTxCode: input.BankTxCode,
			QrCode:     input.QrCode,
		},
	}
	client := resty.New()
	defer client.Close()
	var response GetQrDataResponse
	res, err := client.R().
		SetHeader("APIKEY", a.APIKey).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&response).
		Post(a.Url + "/resources/merch/v1.0/getqrdata")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, errors.New(res.String())
	}
	return &response, nil
}

type CheckQrPaymentInput struct {
	PaymentId string `json:"paymentId"`
	QrCode    string `json:"qrCode"`
}

// QR төлбөр шалгах
func (a *ardpay) CheckQrPayment(input CheckQrPaymentInput) (*CheckQrPaymentResponse, error) {
	body := CheckQrPaymentRequest{
		PayeeCode: a.MerchantID,
		PosNo:     a.PosNo,
		InvoiceID: input.PaymentId,
		QrCode:    input.QrCode,
	}
	client := resty.New()
	defer client.Close()
	var response CheckQrPaymentResponse
	res, err := client.R().
		SetHeader("APIKEY", a.APIKey).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&response).
		Post(a.Url + "/resources/merch/v1.0/checkQrPayment")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, errors.New(res.String())
	}
	if response.ResponseCode != 0 {
		return nil, errors.New("Error code: " + strconv.Itoa(response.ResponseCode) + " Error description: " + response.ResponseDesc)
	}
	return &response, nil
}

type CheckQrInput struct {
	QrCode []string `json:"qrCode"`
}

// QR төлбөр төлөгдсөн эсэхийг шалгах
func (a *ardpay) CheckQr(input CheckQrInput) (*CheckQrResponse, error) {
	body := CheckQrRequest{
		PayeeCode: a.MerchantID,
		PosNo:     a.PosNo,
		QrCode:    input.QrCode,
	}
	client := resty.New()
	defer client.Close()
	var response CheckQrResponse
	res, err := client.R().
		SetHeader("APIKEY", a.APIKey).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&response).
		Post(a.Url + "/resources/merch/v1.0/checkqr")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, errors.New(res.String())
	}
	if response.ResponseCode != 0 {
		return nil, errors.New("Error code: " + strconv.Itoa(response.ResponseCode) + " Error description: " + response.ResponseDesc)
	}
	return &response, nil
}

type CancelQrInput struct {
	QrCodes []string `json:"qrCodes"`
}

// QR төлбөр цуцлах
func (a *ardpay) CancelQrPayment(input CancelQrInput) (*CancelQrPaymentResponse, error) {
	body := CancelQrPaymentRequest{
		PayeeCode: a.MerchantID,
		PosNo:     a.PosNo,
		QrCode:    input.QrCodes,
	}
	client := resty.New()
	defer client.Close()
	var response CancelQrPaymentResponse
	res, err := client.R().
		SetHeader("APIKEY", a.APIKey).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&response).
		Post(a.Url + "/resources/merch/v1.0/cancelqr")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, errors.New(res.String())
	}
	if response.ResponseCode != 0 {
		return nil, errors.New("Error code: " + strconv.Itoa(response.ResponseCode) + " Error description: " + response.ResponseDesc)
	}
	return &response, nil
}

type TanInput struct {
	Amount      float64 `json:"tranAmnt"`
	Currency    string  `json:"tranCur"`
	Description string  `json:"tranDesc"`
	OrderNo     string  `json:"orderNo"`
	Tan         string  `json:"tan"`
	Msisdn      string  `json:"msisdn"`
}

// ТАНтай худалдан авалт
func (a *ardpay) TanPayment(input TanInput) error {
	body := TanPaymentRequest{
		PayeeCode:              a.MerchantID,
		PosNo:                  a.PosNo,
		OrderNo:                input.OrderNo,
		TransactionAmount:      input.Amount,
		TransactionCurrency:    input.Currency,
		TransactionDescription: input.Description,
		Tan:                    input.Tan,
		MSISDN:                 input.Msisdn,
	}
	client := resty.New()
	defer client.Close()
	res, err := client.R().
		SetHeader("APIKEY", a.APIKey).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(a.Url + "/resources/merch/v1.0/purchase")
	if err != nil {
		return err
	}
	if res.StatusCode() != 200 {
		return errors.New(res.String())
	}
	return nil
}
