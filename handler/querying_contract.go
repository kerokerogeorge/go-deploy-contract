package handler

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	storage "github.com/kerokerogeorge/go-deploy-smartcontract/contracts"
)

func QueryingSmartContract() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	client, err := ethclient.Dial(os.Getenv("URL"))
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x06f447ce56C6cE7C40F20EFf118dB6Bb4fa60148")
	instance, err := storage.NewStorage(address, client)
	if err != nil {
		log.Fatal(err)
	}

	value, err := instance.Retrieve(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(value)
}
