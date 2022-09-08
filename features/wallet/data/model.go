package data

import (
	"be11/apiclean/features/user"
	userModel "be11/apiclean/features/user/data"
	"be11/apiclean/features/wallet"

	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Jenis  string
	Nomor  string
	UserID uint
	User   userModel.User
}

func FromCore(data wallet.Core) Wallet {
	return Wallet{
		Jenis:  data.Jenis,
		Nomor:  data.Nomor,
		UserID: data.UserID,
	}
}

func (data *Wallet) toCore() wallet.Core {
	return wallet.Core{
		ID:     data.ID,
		Jenis:  data.Jenis,
		Nomor:  data.Nomor,
		UserID: data.UserID,
		User: user.Core{
			Name: data.User.Name,
		},
	}
}

func toCoreList(data []Wallet) []wallet.Core {
	var dataCore []wallet.Core
	for key, _ := range data {
		dataCore = append(dataCore, data[key].toCore())
	}
	return dataCore
}
