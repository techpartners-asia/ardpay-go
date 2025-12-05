package ardpay

import (
	"errors"

	"resty.dev/v3"
)

type Ardpay interface {
	CreateQrPayment(amount float64, invoiceID string) (*CreateQrPaymentResponse, error) // QR төлбөр үүсгэх
	CheckQrPayment(paymentId, qrCode string) (*CheckQrPaymentResponse, error)           // QR төлбөр шалгах
	CancelQrPayment(qrCode string) (*CancelQrPaymentResponse, error)                    // QR төлбөр цуцлах
	TanPayment(amount float64, desc, orderNo, tan, msisdn string) error                 // ТАНтай худалдан авалт
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

// QR төлбөр үүсгэх
func (a *ardpay) CreateQrPayment(amount float64, invoiceID string) (*CreateQrPaymentResponse, error) {
	body := CreateQrPaymentRequest{
		MerchantID: a.MerchantID,
		PosNo:      a.PosNo,
		Amount:     amount,
		InvoiceID:  invoiceID,
		PaidLimit:  1,
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
	return &response, nil
}

// QR төлбөр шалгах
func (a *ardpay) CheckQrPayment(paymentId, qrCode string) (*CheckQrPaymentResponse, error) {
	body := CheckQrPaymentRequest{
		MerchantID: a.MerchantID,
		PosNo:      a.PosNo,
		InvoiceID:  paymentId,
		QrCode:     qrCode,
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
	return &response, nil
}

// QR төлбөр цуцлах
func (a *ardpay) CancelQrPayment(qrCode string) (*CancelQrPaymentResponse, error) {
	body := CancelQrPaymentRequest{
		MerchantID: a.MerchantID,
		PosNo:      a.PosNo,
		QrCode:     qrCode,
	}
	client := resty.New()
	defer client.Close()
	var response CancelQrPaymentResponse
	res, err := client.R().
		SetHeader("APIKEY", a.APIKey).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&response).
		Post(a.Url + "/resources/merch/v1.0/cancelQrPayment")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, errors.New(res.String())
	}
	return &response, nil
}

// ТАНтай худалдан авалт
func (a *ardpay) TanPayment(amount float64, desc, orderNo, tan, msisdn string) error {
	body := TanPaymentRequest{
		MerchantID:  a.MerchantID,
		PosNo:       a.PosNo,
		OrderNo:     orderNo,
		Amount:      amount,
		Currency:    "MNT",
		Description: desc,
		Tan:         tan,
		MSISDN:      msisdn,
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
