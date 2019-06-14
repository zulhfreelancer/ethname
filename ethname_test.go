package ethname

import (
	"fmt"
	"testing"
)

func TestParity(t *testing.T) {
	// Below node is owned by MyCrypto
	// https://github.com/MyCryptoHQ/MyCrypto/blob/develop/common/libs/nodes/configs.ts

	name := GetEthClientName("https://api.mycryptoapi.com/eth")
	expect := "Parity"
	if name != expect {
		t.Errorf("Name was incorrect, got: %s, want: %s.", name, expect)
	}

	fmt.Println("Result from TestParity :", name)
}

func TestGeth(t *testing.T) {
	// Below node is owned by MyCrypto
	// https://github.com/MyCryptoHQ/MyCrypto/blob/develop/common/libs/nodes/configs.ts

	name := GetEthClientName("https://mainnet.infura.io/v3/c02fff6b5daa434d8422b8ece54c7286")
	expect := "Geth"
	if name != expect {
		t.Errorf("Name was incorrect, got: %s, want: %s.", name, expect)
	}

	fmt.Println("Result from TestGeth   :", name)
}

/* ---------------------------------------------------------------
+++++++++++++++++++++++ Example output: ++++++++++++++++++++++++++
$ go test
Result from TestParity : Parity
Result from TestGeth   : Geth
PASS
ok  	check-ethereum-node-type	2.049s
--------------------------------------------------------------- */
