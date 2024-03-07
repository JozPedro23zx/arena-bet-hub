package betentities

type BetResult struct {
	ID       string
	IDUser   string
	IDBet    string
	Payment  float64
	result   bool
	canceled bool
}

func NewBetResult(id string, user string, bet string, payment float64, victory bool, status bool) *BetResult {
	return &BetResult{ID: id, IDUser: user, IDBet: bet, Payment: payment, result: victory, canceled: status}
}
