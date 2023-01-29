package config

type (
	Event struct {
		WalletID int     `json:"wallet_id"`
		Amount   float64 `json:"amount"`
	}
)
