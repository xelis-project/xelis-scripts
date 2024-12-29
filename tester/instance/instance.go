package instance

import (
	"context"
	"flag"
	"tester/printer"

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

func SetDaemonWalletFlags(flagCmd *flag.FlagSet) {
	flagDaemonEndpoint = flagCmd.String("de", xelisConfig.LOCAL_NODE_RPC, "daemon rpc endpoint")
	flagWalletEndpoint = flagCmd.String("we", xelisConfig.LOCAL_WALLET_RPC, "wallet rpc endpoint")
	flagUsername = flagCmd.String("wu", "test", "wallet rpc username")
	flagPassword = flagCmd.String("wp", "test", "wallet rpc password")
}

func ConnectDaemon() {
	var err error
	Daemon, err = xelisDaemon.NewRPC(context.Background(), *flagDaemonEndpoint)
	if err != nil {
		printer.Fatal(err)
	}
}

func ConnectWallet() {
	var err error
	Wallet, err = xelisWallet.NewRPC(context.Background(), *flagWalletEndpoint, *flagUsername, *flagPassword)
	if err != nil {
		printer.Fatal(err)
	}
}
