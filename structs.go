package ardpay

type (
	CreateQrPaymentRequest struct {
		PayeeCode              string  `json:"payeeCode"`
		PosNo                  string  `json:"posNo"`
		TransactionAmount      float64 `json:"tranAmnt"`
		TransactionCurrency    string  `json:"tranCur"`
		TransactionDescription string  `json:"tranDesc"`
		InvoiceID              string  `json:"invoiceId"`
		PaidLimit              float64 `json:"paidLimit"`
	}
	CreateQrPaymentResponse struct {
		ResponseCode  int    `json:"responseCode"`
		ResponseDesc  string `json:"responseDesc"`
		QrCode        string `json:"qrCode"`
		QrLink        string `json:"qrLink"`
		QpayAccountID string `json:"qpayAccountId"`
	}

	CheckQrPaymentRequest struct {
		PayeeCode string `json:"payeeCode"`
		PosNo     string `json:"posNo"`
		InvoiceID string `json:"paymentId"`
		QrCode    string `json:"qrCode"`
	}
	CheckQrPaymentResponse struct {
		ResponseCode int    `json:"responseCode"`
		ResponseDesc string `json:"responseDesc"`
		Type         string `json:"type"`
	}

	CancelQrPaymentRequest struct {
		PayeeCode string   `json:"payeeCode"`
		PosNo     string   `json:"posNo"`
		QrCode    []string `json:"qrCode"`
	}
	CancelQrPaymentResponse struct {
		ResponseCode int              `json:"responseCode"`
		ResponseDesc string           `json:"responseDesc"`
		Qrlist       []QrlistResponse `json:"qrlist"`
	}
	QrlistResponse struct {
		QrCode       string `json:"qrCode"`
		ResponseDesc string `json:"responseDesc"`
		ResponseCode int    `json:"responseCode"`
	}
	TanPaymentRequest struct {
		PayeeCode              string  `json:"payeeCode"` // Мерчантын дугаар
		PosNo                  string  `json:"posNo"`     // POS дугаар
		OrderNo                string  `json:"orderNo"`   // Захиалгын дугаар
		TransactionAmount      float64 `json:"tranAmnt"`  // Гүйлгээний дүн
		TransactionCurrency    string  `json:"tranCur"`   // Валют
		TransactionDescription string  `json:"tranDesc"`  // Гүйлгээний утга
		Tan                    string  `json:"tan"`       // Тан код
		MSISDN                 string  `json:"msisdn"`    // Утасны дугаар
	}
)
