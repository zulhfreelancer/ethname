## Ethereum Node Name Checker (Golang)

A simple Go package that can tell if you are connected to Geth or Parity node

### Installation

```
// If you are not using Go 1.11 Modules
$ go get github.com/zulhfreelancer/ethname

// If you are using Go 1.11 Modules, run this before executing code below
$ go mod tidy
```

### Usage

```
import (
	"github.com/zulhfreelancer/ethname"
)

func yourFunction() {
    name := ethname.GetEthClientName("https://THE-ETHEREUM-NODE-IP-ADDRESS-AND-PORT")
	fmt.Println(name) // will output "Geth" or "Parity"
}
```
