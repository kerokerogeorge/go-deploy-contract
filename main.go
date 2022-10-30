package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kerokerogeorge/go-deploy-smartcontract/handler"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	// deploy()
	// handler.LoadSmartContract()
	// handler.QueryingSmartContract()
	// handler.WritingToSmartContract()
	// handler.DeployTestContract()
	handler.WritingToMoodSmartContract()
	// handler.QueryingMoodContract()
}
