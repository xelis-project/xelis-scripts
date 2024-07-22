package main

import (
	"github.com/google/uuid"
	xelisConfig "github.com/xelis-project/xelis-go-sdk/config"
	xelisWallet "github.com/xelis-project/xelis-go-sdk/wallet"
)

type BigTransferArgs struct {
	MaxTransfers int
	Destination  string
}

func BigTransfer(args BigTransferArgs) {
	var transfers []xelisWallet.TransferOut

	for i := 0; i < args.MaxTransfers; i++ {
		var extra interface{}
		var err error
		extra, err = uuid.NewRandom()
		if err != nil {
			FatalError(err)
		}

		transfers = append(transfers, xelisWallet.TransferOut{
			Amount:      0,
			Destination: args.Destination,
			Asset:       xelisConfig.XELIS_ASSET,
			ExtraData:   &extra,
		})
	}

	stopLoad := PrintLoad("Buidling")
	tx, err := Wallet.BuildTransaction(xelisWallet.BuildTransactionParams{
		Transfers: transfers,
		Broadcast: false,
		TxAsHex:   true,
	})
	stopLoad()
	if err != nil {
		FatalError(err)
	}

	stopLoad = PrintLoad("Submitting")
	_, err = Daemon.SubmitTransaction(tx.TxAsHex)
	stopLoad()
	if err != nil {
		FatalError(err)
	}

	PrintSuccess("New transaction sent %s to %s\n", tx.Hash, args.Destination)
}
