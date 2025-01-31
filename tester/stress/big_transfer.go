package stress

import (
	"tester/instance"
	"tester/printer"

	"github.com/google/uuid"
	xelisConfig "github.com/xelis-project/xelis-go-sdk/config"
	xelisWallet "github.com/xelis-project/xelis-go-sdk/wallet"
)

type BigTransferArgs struct {
	MaxTransfers int
	Destination  string
}

func BigTransfer(args BigTransferArgs) {
	var transfers []xelisWallet.TransferBuilder

	for i := 0; i < args.MaxTransfers; i++ {
		var extra interface{}
		var err error
		extra, err = uuid.NewRandom()
		if err != nil {
			printer.Fatal(err)
		}

		transfers = append(transfers, xelisWallet.TransferBuilder{
			Amount:      0,
			Destination: args.Destination,
			Asset:       xelisConfig.XELIS_ASSET,
			ExtraData:   &extra,
		})
	}

	stopLoad := printer.Load("Buidling")
	tx, err := instance.Wallet.BuildTransaction(xelisWallet.BuildTransactionParams{
		Transfers: transfers,
		Broadcast: false,
		TxAsHex:   true,
	})
	stopLoad()
	if err != nil {
		printer.Fatal(err)
	}

	stopLoad = printer.Load("Submitting")
	_, err = instance.Daemon.SubmitTransaction(*tx.TxAsHex)
	stopLoad()
	if err != nil {
		printer.Fatal(err)
	}

	printer.Success("New transaction sent %s to %s\n", tx.Hash, args.Destination)
}
