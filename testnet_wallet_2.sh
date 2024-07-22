#! /bin/bash

cargo run --bin xelis_wallet --release -- --network testnet --log-level debug --daemon-address testnet-node.xelis.io --wallet-path ./wallets/testnet_2 --password test --rpc-password test --rpc-bind-address 127.0.0.1:8082 --rpc-username test