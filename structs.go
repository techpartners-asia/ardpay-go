package ardpay

type (
	CreateQrPaymentRequest struct {
		MerchantID  string  `json:"payeeCode"`
		PosNo       string  `json:"posNo"`
		Amount      float64 `json:"tranAmnt"`
		Currency    string  `json:"tranCur"`
		Description string  `json:"tranDesc"`
		InvoiceID   string  `json:"invoiceId"`
		PaidLimit   float64 `json:"paidLimit"`
	}
	CreateQrPaymentResponse struct {
		ResponseCode  int    `json:"responseCode"`
		ResponseDesc  string `json:"responseDesc"`
		QrCode        string `json:"qrCode"`
		QrLink        string `json:"qrLink"`
		QpayAccountID string `json:"qpayAccountId"`
		PaymentID     string `json:"paymentId"`
	}

	CheckQrPaymentRequest struct {
		MerchantID string `json:"payeeCode"`
		PosNo      string `json:"posNo"`
		InvoiceID  string `json:"paymentId"`
		QrCode     string `json:"qrCode"`
	}
	CheckQrPaymentResponse struct {
		ResponseCode int    `json:"responseCode"`
		ResponseDesc string `json:"responseDesc"`
		Type         string `json:"type"`
	}

	CancelQrPaymentRequest struct {
		MerchantID string   `json:"payeeCode"`
		PosNo      string   `json:"posNo"`
		QrCode     []string `json:"qrCode"`
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
		MerchantID  string  `json:"payeeCode"` // Мерчантын дугаар
		PosNo       string  `json:"posNo"`     // POS дугаар
		OrderNo     string  `json:"orderNo"`   // Захиалгын дугаар
		Amount      float64 `json:"tranAmnt"`  // Гүйлгээний дүн
		Currency    string  `json:"tranCur"`   // Валют
		Description string  `json:"tranDesc"`  // Гүйлгээний утга
		Tan         string  `json:"tan"`       // Тан код
		MSISDN      string  `json:"msisdn"`    // Утасны дугаар
	}
)
