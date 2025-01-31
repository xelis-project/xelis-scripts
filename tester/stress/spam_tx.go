package stress

import (
	"math/rand"
	"os"
	"os/signal"
	"tester/instance"
	"tester/printer"
	"time"

	"github.com/google/uuid"
	xelisConfig "github.com/xelis-project/xelis-go-sdk/config"
	xelisDaemon "github.com/xelis-project/xelis-go-sdk/daemon"
	xelisWallet "github.com/xelis-project/xelis-go-sdk/wallet"
)

type SpamArgs struct {
	Timeout     float64
	Amount      uint64
	Destination string
	RandAddr    bool
}

func SpamTx(args SpamArgs) {
	sleep := time.Duration(args.Timeout) * time.Second

	addrs, err := instance.Daemon.GetAccounts(xelisDaemon.GetAccountsParams{Skip: 0, Maximum: 20})
	if err != nil {
		printer.Fatal(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
out:
	for {
		select {
		case <-c:
			break out
		default:
			stopLoad := printer.Load("Buidling")

			var extra interface{}
			var err error
			extra, err = uuid.NewRandom()
			if err != nil {
				printer.Fatal(err)
			}

			destination := args.Destination
			if args.RandAddr {
				destination = addrs[rand.Intn(len(addrs))]
			}

			tx, err := instance.Wallet.BuildTransaction(xelisWallet.BuildTransactionParams{
				Broadcast: true,
				Transfers: []xelisWallet.TransferBuilder{
					{
						Amount:      args.Amount,
						Asset:       xelisConfig.XELIS_ASSET,
						Destination: destination,
						ExtraData:   &extra,
					},
				},
			})
			stopLoad()
			if err != nil {
				printer.Error(err)
			} else {
				printer.Success("New transaction sent %s to %s\n", tx.Hash, destination)
			}

			time.Sleep(sleep)
		}
	}
}
