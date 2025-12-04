package ardpay

import (
	"errors"

	"resty.dev/v3"
)

type Ardpay interface {
	CreateQrPayment(amount float64, invoiceID string) (response CreateQrPaymentResponse, err error) // QR төлбөр үүсгэх
	CheckQrPayment(paymentId, qrCode string) (response CheckQrPaymentResponse, err error)           // QR төлбөр шалгах
	CancelQrPayment(qrCode string) (response CancelQrPaymentResponse, err error)                    // QR төлбөр цуцлах
	TanPayment(amount float64, desc, orderNo, tan, msisdn string) error                             // Тантай худалдан авалт
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
func (a *ardpay) CreateQrPayment(amount float64, invoiceID string) (response CreateQrPaymentResponse, err error) {
	body := CreateQrPaymentRequest{
		MerchantID: a.MerchantID,
		PosNo:      a.PosNo,
		Amount:     amount,
		InvoiceID:  invoiceID,
		PaidLimit:  1,
	}
	client := resty.New()
	defer client.Close()
	res, err := client.R().
		SetHeader("APIKEY", a.APIKey).
		SetHeader("Content-Type", "application/json").
		SetBody(body).        // default request content type is JSON
		SetResult(&response). // or SetResult(LoginResponse{}).
		Post(a.Url + "/resources/merch/v1.0/createqr")
	if err != nil {
		return response, err
	}
	if res.IsError() {
		return response, errors.New(res.String())
	}
	return response, nil
}

// QR төлбөр шалгах
func (a *ardpay) CheckQrPayment(paymentId, qrCode string) (response CheckQrPaymentResponse, err error) {
	body := CheckQrPaymentRequest{
		MerchantID: a.MerchantID,
		PosNo:      a.PosNo,
		InvoiceID:  paymentId,
		QrCode:     qrCode,
	}
	client := resty.New()
	defer client.Close()
	res, err := client.R().
		SetHeader("APIKEY", a.APIKey).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&response).
		Post(a.Url + "/resources/merch/v1.0/checkQrPayment")
	if err != nil {
		return response, err
	}
	if res.IsError() {
		return response, errors.New(res.String())
	}
	return response, nil
}

// QR төлбөр цуцлах
func (a *ardpay) CancelQrPayment(qrCode string) (response CancelQrPaymentResponse, err error) {
	body := CancelQrPaymentRequest{
		MerchantID: a.MerchantID,
		PosNo:      a.PosNo,
		QrCode:     qrCode,
	}
	client := resty.New()
	defer client.Close()
	res, err := client.R().
		SetHeader("APIKEY", a.APIKey).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&response).
		Post(a.Url + "/resources/merch/v1.0/cancelQrPayment")
	if err != nil {
		return response, err
	}
	if res.IsError() {
		return response, errors.New(res.String())
	}
	return response, nil
}

// Тантай худалдан авалт
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
