package data

import (
	"be11/apiclean/features/wallet"
	"be11/apiclean/utils/helper"

	"gorm.io/gorm"
)

type walletData struct {
	db *gorm.DB
}

func New(db *gorm.DB) wallet.DataInterface {
	return &walletData{
		db: db,
	}
}

func (repo *walletData) Insert(data wallet.Core) (int, error) {
	dataModel := FromCore(data)
	tx := repo.db.Create(&dataModel)
	if tx.Error != nil {
		return 0, helper.CheckQueryError(tx.Error)
	}

	return int(tx.RowsAffected), nil
}
