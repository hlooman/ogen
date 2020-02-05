# Ogen libraries

This repository contains a group of libraries used by the golang implementation (Ogen) of the Olympus protocol.

Some of this libraries are a direct copy of an external repository. 

* `amount`: An amount handling interface copied and modified from [btcsuite](https://github.com/btcsuite) amount interface.  [Original Source](https://github.com/btcsuite/btcutil/tree/master/amount.go)
* `base58`: An implementation of base58 encode.
* `bech32`: An implementation of bech32 encode.
* `bip39`: An implementation of bip39 on golang.
* `chainhash`: Hashing functions utility for Ogen.
* `hdwallets`: A HD wallets implementation using bls key pairs.
* `serializer`: A library to serialize common go objects based on [btcsuite](https://github.com/btcsuite) `wire` package. [Original Source](https://github.com/btcsuite/btcd/tree/master/wire) 