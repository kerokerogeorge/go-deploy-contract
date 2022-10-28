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
	"github.com/joho/godotenv"
	storage "github.com/kerokerogeorge/go-deploy-smartcontract/contracts"
)

func WritingToSmartContract() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	client, err := ethclient.Dial(os.Getenv("URL"))
	if err != nil {
		log.Fatal(err)
	}

	// address := common.HexToAddress("0x06f447ce56C6cE7C40F20EFf118dB6Bb4fa60148")
	// instance, err := storage.NewStorage(address, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// value, err := instance.Retrieve(nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

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
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress("0x06f447ce56C6cE7C40F20EFf118dB6Bb4fa60148")
	instance, err := storage.NewStorage(address, client)
	if err != nil {
		log.Fatal(err)
	}
	value := big.NewInt(5)
	log.Println("Value:", value)
	tx, err := instance.Store(auth, value)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent:
	result, err := instance.Retrieve(nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("result", result)
}
