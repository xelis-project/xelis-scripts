package main

import (
	"math/rand"
	"os"
	"os/signal"
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

	addrs, err := Daemon.GetAccounts(xelisDaemon.GetAccountsParams{Skip: 0, Maximum: 20})
	if err != nil {
		FatalError(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
out:
	for {
		select {
		case <-c:
			break out
		default:
			stopLoad := PrintLoad("Buidling")

			var extra interface{}
			var err error
			extra, err = uuid.NewRandom()
			if err != nil {
				FatalError(err)
			}

			destination := args.Destination
			if args.RandAddr {
				destination = addrs[rand.Intn(len(addrs))]
			}

			tx, err := Wallet.BuildTransaction(xelisWallet.BuildTransactionParams{
				Broadcast: true,
				Transfers: []xelisWallet.TransferOut{
					xelisWallet.TransferOut{
						Amount:      args.Amount,
						Asset:       xelisConfig.XELIS_ASSET,
						Destination: destination,
						ExtraData:   &extra,
					},
				},
			})
			stopLoad()
			if err != nil {
				PrintError(err)
			} else {
				PrintSuccess("New transaction sent %s to %s\n", tx.Hash, destination)
			}

			time.Sleep(sleep)
		}
	}
}
