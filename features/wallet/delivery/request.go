package delivery

import (
	"be11/apiclean/features/wallet"
)

type WalletRequest struct {
	Jenis  string `json:"jenis" form:"jenis"`
	Nomor  string `json:"nomor" form:"nomor"`
	UserID int    `json:"user_id" form:"user_id"`
}

func toCore(data WalletRequest) wallet.Core {
	return wallet.Core{
		Jenis:  data.Jenis,
		Nomor:  data.Nomor,
		UserID: uint(data.UserID),
	}
}
