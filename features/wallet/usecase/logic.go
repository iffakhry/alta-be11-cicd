package usecase

import "be11/apiclean/features/wallet"

type walletUsecase struct {
	walletData wallet.DataInterface
}

func New(data wallet.DataInterface) wallet.UsecaseInterface {
	return &walletUsecase{
		walletData: data,
	}
}

func (usecase *walletUsecase) Add(data wallet.Core) (int, error) {
	row, err := usecase.walletData.Insert(data)
	if err != nil {
		return 0, err
	}
	return row, nil
}
