package config

type (
	Event struct {
		WalletID int     `json:"walletID"`
		Amount   float64 `json:"amount"`
	}
)
