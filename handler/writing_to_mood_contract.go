package handler

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	storage "github.com/kerokerogeorge/go-deploy-smartcontract/contracts"
)

func WritingToMoodSmartContract() {
	client, err := ethclient.Dial(os.Getenv("URL"))
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
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
	// コントラクトアドレス
	address := common.HexToAddress("0xee32E4fD377bd797D8DAd5e42cb3ee3Aba681288")

	gasTipCap, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		log.Println(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = 300000
	auth.GasPrice = gasPrice
	auth.GasTipCap = gasTipCap

	instance, err := storage.NewStorage(address, client)
	if err != nil {
		log.Fatal(err)
	}
	tx, err := instance.SetMood(auth, "sign from transaction")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("tx sent: %s", tx.Hash().Hex())
}
