package handler

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	storage "github.com/kerokerogeorge/go-deploy-smartcontract/contracts"
)

func QueryingMoodContract() {
	client, err := ethclient.Dial(os.Getenv("URL"))
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0xee32E4fD377bd797D8DAd5e42cb3ee3Aba681288")
	instance, err := storage.NewStorage(address, client)
	if err != nil {
		log.Fatal(err)
	}

	value, err := instance.GetMood(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(value)
}
