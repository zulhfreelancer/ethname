package ethname

import (
	"context"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/rpc"
)

// GetEthClientName - will output "Parity" or "Geth"
func GetEthClientName(nodeURL string) string {
	client, _ := rpc.Dial(nodeURL)

	var longName string
	err := client.CallContext(context.Background(), &longName, "web3_clientVersion")
	if err != nil {
		log.Fatal(err)
	}

	shortName := (strings.Split((strings.Split(longName, "-"))[0], "/"))[0]
	return shortName
}
