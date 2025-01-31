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
	TxCount      int
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

	var txs []xelisWallet.TransactionResponse

	stopLoad := printer.Load("Buidling txs (%d)", args.TxCount)
	for i := 0; i < args.TxCount; i++ {
		tx, err := instance.Wallet.BuildTransaction(xelisWallet.BuildTransactionParams{
			Transfers: transfers,
			Broadcast: false,
			TxAsHex:   true,
		})
		if err != nil {
			printer.Fatal(err)
			stopLoad()
			return
		}
		txs = append(txs, tx)
	}
	stopLoad()

	stopLoad = printer.Load("Submitting txs (%d)", args.TxCount)
	for _, tx := range txs {
		_, err := instance.Daemon.SubmitTransaction(*tx.TxAsHex)
		if err != nil {
			printer.Fatal(err)
			stopLoad()
			return
		}

		printer.Success("New tx sent %s\n", tx.Hash)
	}
	stopLoad()
}
