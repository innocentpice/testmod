package wallets

type Wallet struct {
	Money float32
}

func NewWallet(number float32) Wallet{
	wallet := Wallet{
		Money: number,
	}
	return wallet
}

func (w *Wallet) WithDraw(number float32) float32{
	if w.Money < number {
		return 0.0
	}
	w.Money -= number;
	return number
}

func (w *Wallet) Deposit(number float32) float32{
	w.Money += number;
	return number
}

