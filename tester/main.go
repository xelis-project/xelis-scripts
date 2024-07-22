package main

import (
	"context"
	"flag"
	"math/rand"
	"os"

	xelisConfig "github.com/xelis-project/xelis-go-sdk/config"
	xelisDaemon "github.com/xelis-project/xelis-go-sdk/daemon"
	xelisWallet "github.com/xelis-project/xelis-go-sdk/wallet"
)

var (
	flagDaemonEndpoint *string
	flagWalletEndpoint *string
	flagUsername       *string
	flagPassword       *string
)

var (
	Daemon *xelisDaemon.RPC
	Wallet *xelisWallet.RPC
)

func setDaemonWalletFlags(flagCmd *flag.FlagSet) {
	flagDaemonEndpoint = flagCmd.String("de", xelisConfig.TESTNET_NODE_RPC, "daemon rpc endpoint")
	flagWalletEndpoint = flagCmd.String("we", xelisConfig.LOCAL_WALLET_RPC, "wallet rpc endpoint")
	flagUsername = flagCmd.String("wu", "test", "wallet rpc username")
	flagPassword = flagCmd.String("wp", "test", "wallet rpc password")
}

func connectDaemon() {
	var err error
	Daemon, err = xelisDaemon.NewRPC(context.Background(), *flagDaemonEndpoint)
	if err != nil {
		FatalError(err)
	}
}

func connectWallet() {
	var err error
	Wallet, err = xelisWallet.NewRPC(context.Background(), *flagWalletEndpoint, *flagUsername, *flagPassword)
	if err != nil {
		FatalError(err)
	}
}

func parseFlag(flagCmd *flag.FlagSet) {
	flagHelp := flagCmd.Bool("h", false, "help")

	flagCmd.Parse(os.Args[2:])
	help := *flagHelp

	if help {
		flagCmd.Usage()
		os.Exit(1)
	}
}

func main() {
	cmd := os.Args[1]

	switch cmd {
	case "spam_tx":
		flagCmd := flag.NewFlagSet("spam_tx", flag.ExitOnError)

		setDaemonWalletFlags(flagCmd)
		flagAmount := flagCmd.Uint64("a", 1, "amount in atomic value")
		flagDestination := flagCmd.String("d", "", "destination")
		flagTimeout := flagCmd.Float64("t", 1, "send transaction interval per second")
		flagRandAddr := flagCmd.Bool("r", false, "use random addr")
		parseFlag(flagCmd)
		connectDaemon()
		connectWallet()

		destination := *flagDestination
		amount := *flagAmount
		timeout := *flagTimeout
		randAddr := *flagRandAddr

		if !randAddr && destination == "" {
			Print("Missing destination address. Specify with -d.")
			return
		}

		SpamTx(SpamArgs{Amount: amount, Destination: destination, Timeout: timeout, RandAddr: randAddr})
	case "big_transfer":
		flagCmd := flag.NewFlagSet("big_transfer", flag.ExitOnError)

		setDaemonWalletFlags(flagCmd)
		flagDestination := flagCmd.String("d", "", "destination")
		flagMaxTransfers := flagCmd.Int("m", 255, "max total transfers")
		flagRandAddr := flagCmd.Bool("r", false, "use random addr")
		parseFlag(flagCmd)
		connectDaemon()
		connectWallet()

		destination := *flagDestination
		maxTransfers := *flagMaxTransfers
		randAddr := *flagRandAddr

		if randAddr {
			addrs, err := Daemon.GetAccounts(xelisDaemon.GetAccountsParams{})
			if err != nil {
				FatalError(err)
			}

			destination = addrs[rand.Intn(len(addrs))]
		}

		if destination == "" {
			Print("Missing destination address. Specify with -d.")
			return
		}

		BigTransfer(BigTransferArgs{Destination: destination, MaxTransfers: maxTransfers})
	default:
		Print("Specify a command.")
	}
}
