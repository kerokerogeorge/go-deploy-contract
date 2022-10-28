package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	storage "github.com/kerokerogeorge/go-deploy-smartcontract/contracts"
)

func main() {
	client, err := ethclient.Dial("https://eth-goerli.g.alchemy.com/v2/b3-8tyI0X58hzWaRBudDw8fwZe83ugsJ")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("d33172bdb1848eb8b571e3043bf0d7e2bd37dc0c63cbfa6488be3021270e1fb3")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Println(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	// auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address, tx, instance, err := storage.DeployStorage(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(address.Hex())   // ==> 0x06f447ce56C6cE7C40F20EFf118dB6Bb4fa60148
	fmt.Println(tx.Hash().Hex()) // ==> 0x40ccf618ebacaf3f0b7587f62e1f434bc05303f9e7100ed85c58b914256eb12d

	_ = instance
}
