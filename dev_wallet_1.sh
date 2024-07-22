#! /bin/bash

# keep pre assigned seed to receive mining reward when launching the dev_miner.sh script
cargo run --bin xelis_wallet --release -- --network dev --wallet-path ./wallets/dev --password test --rpc-password test --rpc-bind-address 127.0.0.1:8081 --rpc-username test --seed "etched copy copy yahoo pairing rated pause apart rafts motherly rekindle twice tequila soccer jogger rarest delayed bovine orchid tavern ferry aplomb pairing pivot twice"