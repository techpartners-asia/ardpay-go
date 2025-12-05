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

	GetQrDataRequest struct {
		PayeeCode    string   `json:"payeeCode"`    // Мерчантын дугаар
		PosNo        string   `json:"posNo"`        // POS дугаар
		Type         string   `json:"type"`         // Төрөл /1 - qrPayment, 2 - confirm /
		CustomerCode string   `json:"customerCode"` // Банкны харилцагчийн дугаар
		BankCode     string   `json:"bankCode"`     // Банкны дугаар
		BankVerfCode string   `json:"bankVerfCode"` // Банкны харилцагчийн дугаар
		JsonData     JsonData `json:"jsonData"`
	}
	JsonData struct {
		BankTxCode string `json:"bankTxCode"` // Гүйлгээний банкны код
		QrCode     string `json:"qrCode"`     // QR код
	}
	GetQrDataResponse struct {
		PayeeCode   string        `json:"payeeCode"`
		PayeeType   string        `json:"payeeType"`
		TxnID       string        `json:"txnId"`
		ProdCode    string        `json:"prodCode"`
		QrSystem    string        `json:"qrSystem"`
		PaymentLine []PaymentLine `json:"paymentLine"`
		ChargeLine  []ChargeLine  `json:"chargeLine"`
	}
	PaymentLine struct {
		PaymentId     string `json:"paymentId"`
		PaymentDate   string `json:"paymentDate"`
		ObjectType    string `json:"objectType"`
		ObjectId      string `json:"objectId"`
		InvoiceId     string `json:"invoiceId"`
		BankCode      string `json:"bankCode"`
		BankName      string `json:"bankName"`
		BankTxnId     string `json:"bankTxnId"`
		BankTxnDate   string `json:"bankTxnDate"`
		AcntNo        string `json:"acntNo"`
		AcntType      string `json:"acntType"`
		AcntName      string `json:"acntName"`
		AcntCurCode   string `json:"acntCurCode"`
		Amnt          string `json:"amnt"`
		CurCode       string `json:"curCode"`
		ExchangeRate  string `json:"exchangeRate"`
		Desc          string `json:"desc"`
		CustFeeFlag   string `json:"custFeeFlag"`
		StatusCode    string `json:"statusCode"`
		StatusMsg     string `json:"statusMsg"`
		ColorCode     string `json:"colorCode"`
		PaymentName   string `json:"paymentName"`
		PaymentStatus string `json:"paymentStatus"`
		PaymentAmnt   string `json:"paymentAmnt"`
		PaymentRate   string `json:"paymentRate"`
		TxnAmnt       string `json:"txnAmnt"`
		TxnCurCode    string `json:"txnCurCode"`
		NbfiCode      string `json:"nbfiCode"`
		NbfiAcntCode  string `json:"nbfiAcntCode"`
		NbfiAcntType  string `json:"nbfiAcntType"`
		NbfiAcntName  string `json:"nbfiAcntName"`
		NbfiCurCode   string `json:"nbfiCurCode"`
	}
	ChargeLine struct {
		ChargeId     string `json:"chargeId"`
		PaymentId    string `json:"paymentId"`
		BankCode     string `json:"bankCode"`
		BankName     string `json:"bankName"`
		AcntName     string `json:"acntName"`
		AcntNo       string `json:"acntNo"`
		AcntCurCode  string `json:"acntCurCode"`
		Amnt         string `json:"amnt"`
		CurCode      string `json:"curCode"`
		Desc         string `json:"desc"`
		BankTxnId    string `json:"bankTxnId"`
		BankTxnDate  string `json:"bankTxnDate"`
		ExchangeRate string `json:"exchangeRate"`
		StatusCode   string `json:"statusCode"`
		StatusMsg    string `json:"statusMsg"`
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

	CheckQrRequest struct {
		PayeeCode string   `json:"payeeCode"`
		PosNo     string   `json:"posNo"`
		QrCode    []string `json:"qrCode"`
	}
	CheckQrResponse struct {
		ResponseCode int      `json:"responseCode"`
		ResponseDesc string   `json:"responseDesc"`
		QrList       []QrList `json:"qrList"`
	}
	QrList struct {
		QrCode       string `json:"qrCode"`
		PaymentId    int    `json:"paymentId"`
		TranId       string `json:"tranId"`
		TranDate     string `json:"tranDate"`
		TranAmnt     int    `json:"tranAmnt"`
		InvoiceId    string `json:"invoiceId"`
		ResponseCode int    `json:"responseCode"`
		ResponseDesc string `json:"responseDesc"`
	}

	CancelQrPaymentRequest struct {
		PayeeCode string   `json:"payeeCode"`
		PosNo     string   `json:"posNo"`
		QrCode    []string `json:"qrCode"`
	}
	CancelQrPaymentResponse struct {
		ResponseCode int      `json:"responseCode"`
		ResponseDesc string   `json:"responseDesc"`
		Qrlist       []QrList `json:"qrlist"`
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

type (
	CheckCallbackResponse struct {
		RetType     int    `json:"retType"`     // Алдааны дугаар /Амжилттай үед 0 гэж ирнэ/
		RetDesc     string `json:"retDesc"`     // Тайлбар
		PaymentId   int    `json:"paymentId"`   // Төлбөрийн дугаар
		BankTxnId   int    `json:"bankTxnId"`   // Банкны гүйлгээний дугаар
		BankTxnDate string `json:"bankTxnDate"` // Гүйлгээ хийсэн огноо
		Amount      int    `json:"amount"`      // Гүйлгээний дүн
		BillId      int    `json:"billId"`      // Биллийн дугаар
	}
)
