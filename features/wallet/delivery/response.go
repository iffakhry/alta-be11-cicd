package delivery

import (
	"be11/apiclean/features/wallet"
)

type WalletResponse struct {
	ID     uint   `json:"id"`
	Jenis  string `json:"jenis"`
	Nomor  string `json:"nomor,omitempty"`
	UserID uint   `json:"user_id,omitempty"`
}

func fromCore(data wallet.Core) WalletResponse {
	return WalletResponse{
		ID:     data.ID,
		Jenis:  data.Jenis,
		Nomor:  data.Nomor,
		UserID: data.UserID,
	}
}

func fromCoreList(data []wallet.Core) []WalletResponse {
	var dataResponse []WalletResponse
	for _, v := range data {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
