package main

import (
	"flag"
	"math/rand"
	"os"
	"tester/instance"
	"tester/printer"
	"tester/smart_contract"
	"tester/stress"

	xelisConfig "github.com/xelis-project/xelis-go-sdk/config"
	xelisDaemon "github.com/xelis-project/xelis-go-sdk/daemon"
)

func parseFlag(flagCmd *flag.FlagSet) {
	flagHelp := flagCmd.Bool("h", false, "help")

	flagCmd.Parse(os.Args[2:])
	help := *flagHelp

	if help {
		flagCmd.Usage()
		os.Exit(1)
	}
}

func initDaemonAndWallet(flagCmd *flag.FlagSet) {
	instance.SetDaemonWalletFlags(flagCmd)
	parseFlag(flagCmd)
	instance.ConnectDaemon()
	instance.ConnectWallet()
}

func main() {
	if len(os.Args) <= 1 {
		os.Args = append(os.Args, "")
	}
	cmd := os.Args[1]

	switch cmd {
	case "spam_tx":
		flagCmd := flag.NewFlagSet("spam_tx", flag.ExitOnError)
		flagAmount := flagCmd.Uint64("a", 1, "amount in atomic value")
		flagDestination := flagCmd.String("d", "", "destination")
		flagTimeout := flagCmd.Float64("t", 1, "send transaction interval per second")
		flagRandAddr := flagCmd.Bool("r", false, "use random addr")
		initDaemonAndWallet(flagCmd)

		destination := *flagDestination
		amount := *flagAmount
		timeout := *flagTimeout
		randAddr := *flagRandAddr

		if !randAddr && destination == "" {
			printer.Print("Missing destination address. Specify with -d.")
			return
		}

		stress.SpamTx(stress.SpamArgs{
			Amount:      amount,
			Destination: destination,
			Timeout:     timeout,
			RandAddr:    randAddr,
		})
	case "big_transfer":
		flagCmd := flag.NewFlagSet("big_transfer", flag.ExitOnError)
		flagDestination := flagCmd.String("d", "", "destination")
		flagMaxTransfers := flagCmd.Int("m", 255, "max total transfers")
		flagRandAddr := flagCmd.Bool("r", false, "use random addr")
		initDaemonAndWallet(flagCmd)

		destination := *flagDestination
		maxTransfers := *flagMaxTransfers
		randAddr := *flagRandAddr

		if randAddr {
			addrs, err := instance.Daemon.GetAccounts(xelisDaemon.GetAccountsParams{})
			if err != nil {
				printer.Fatal(err)
			}

			destination = addrs[rand.Intn(len(addrs))]
		}

		if destination == "" {
			printer.Print("Missing destination address. Specify with -d.")
			return
		}

		stress.BigTransfer(stress.BigTransferArgs{
			Destination:  destination,
			MaxTransfers: maxTransfers,
		})
	case "sc_install_helloworld":
		flagCmd := flag.NewFlagSet("sc_install", flag.ExitOnError)
		initDaemonAndWallet(flagCmd)
		smart_contract.InstallHelloWorld()
	case "sc_deposit":
		flagCmd := flag.NewFlagSet("sc_deposit", flag.ExitOnError)
		flagContract := flagCmd.String("c", "", "contract hash / txid")
		initDaemonAndWallet(flagCmd)
		contract := *flagContract
		smart_contract.DepositFunds(contract)
	case "sc_balance":
		flagCmd := flag.NewFlagSet("sc_balance", flag.ExitOnError)
		flagContract := flagCmd.String("c", "", "contract hash / txid")
		flagAsset := flagCmd.String("a", xelisConfig.XELIS_ASSET, "asset id")
		initDaemonAndWallet(flagCmd)
		contract := *flagContract
		asset := *flagAsset
		smart_contract.Balance(contract, asset)
	default:
		printer.Print("Specify a command.")
	}
}
