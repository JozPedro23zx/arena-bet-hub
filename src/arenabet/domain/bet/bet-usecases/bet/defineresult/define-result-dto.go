package defineresult

type BetInputDto struct {
	ID      string  `json:"id"`
	IDUser  string  `json:"iduser"`
	IDBet   string  `json:"idbet"`
	Payment float64 `json:"idpayment"`
	Victory bool    `json:"victory"`
	Status  bool    `json:"status"`
}
