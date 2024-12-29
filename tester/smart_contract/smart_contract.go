package smart_contract

import (
	"tester/instance"
	"tester/printer"

	xelisConfig "github.com/xelis-project/xelis-go-sdk/config"
	"github.com/xelis-project/xelis-go-sdk/daemon"
	xelisWallet "github.com/xelis-project/xelis-go-sdk/wallet"
)

//	entry main() {
//		println("Hello, World!");
//		return 0;
//	}
var hello_world_hex = "00000000000200090d48656c6c6f2c20576f726c64210004000000000000000000010000000c00000014540000010001001000010000"

func InstallHelloWorld() {
	tx, err := instance.Wallet.BuildTransaction(xelisWallet.BuildTransactionParams{
		Broadcast:      true,
		DeployContract: &hello_world_hex,
	})
	if err != nil {
		printer.Fatal(err)
	}

	printer.Success("%s\r", tx.Hash)
}

func DepositFunds(contract string) {
	printer.Print("Contract: %s\n", contract)

	tx, err := instance.Wallet.BuildTransaction(xelisWallet.BuildTransactionParams{
		InvokeContract: &xelisWallet.InvokeContractBuilder{
			Contract:   contract,
			MaxGas:     1000,
			ChunkId:    0,
			Parameters: []interface{}{},
			Deposits: map[string]xelisWallet.ContractDepositBuilder{
				xelisConfig.XELIS_ASSET: {
					Amount:  1,
					Private: false,
				},
			},
		},
		Broadcast: true,
	})
	if err != nil {
		printer.Fatal(err)
	}

	printer.Success("Tx: %s\n", tx.Hash)
}

func Balance(contract string, asset string) {
	printer.Print("Contract: %s\n", contract)
	printer.Print("Asset: %s\n", asset)

	data, err := instance.Daemon.GetContractBalance(daemon.GetContractBalanceParams{
		Contract: contract,
		Asset:    asset,
	})
	if err != nil {
		printer.Fatal(err)
	}
	printer.Success("%+v", data)
}
