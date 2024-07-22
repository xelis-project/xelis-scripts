#! /bin/bash

cargo run --bin xelis_wallet --release -- --network testnet --log-level info --daemon-address testnet-node.xelis.io --wallet-path ./wallets/testnet --password test --rpc-password test --rpc-bind-address 127.0.0.1:8081 --rpc-username test