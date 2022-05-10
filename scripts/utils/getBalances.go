package utils

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetWETH() {
	fmt.Println("will convert ETH to WETH")

	// import account to send transaction
	// get accountBalance based of network and chainid

	// https://goethereumbook.org/en/smart-contract-read-erc20/

	// interact to the WETH contract on the network
	// *Iterate through the chainID later*
	// the deployer account or (account to send transaction)
	// Call the Deposit function
	// print the pending balance
	// Then print the balance after

}

// eventually will be a mapping that tracks all tokens

func Get_AVAX_Balance(chain *ethclient.Client, user string) {
	fmt.Println("will print AVAX balance")
	account := common.HexToAddress(user)
	balance, err := chain.BalanceAt(context.Background(), account, nil)
	formatBalance := new(big.Float)
	formatBalance.SetString(balance.String())
	balanceValue := new(big.Float).Quo(formatBalance, big.NewFloat(math.Pow10(18)))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("AVAX Address Balance:", balanceValue)
	_ = chain

}

// func PendingAvaxBalance(user string) {
// 	// Connect to Blockchain
// 	chain, err := ethclient.Dial(fuji)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	account := common.HexToAddress(user)
// 	pendingBalance, err := chain.PendingBalanceAt(context.Background(), account)
// 	formatPBalance := new(big.Float)
// 	formatPBalance.SetString(pendingBalance.String())
// 	balancePValue := new(big.Float).Quo(formatPBalance, big.NewFloat(math.Pow10(18)))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("connceted to", account, "|| Pending Balance", balancePValue)
// }

// func AvaxBalanceAtBlock(user string) {
// 	// Connect to Blockchain
// 	chain, err := ethclient.Dial(fuji)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	account := common.HexToAddress(user)
// 	// Get Balance @ a block
// 	blockNumber := big.NewInt(8764051)
// 	balanceAt, err := chain.BalanceAt(context.Background(), account, blockNumber)
// 	formatBalanceAt := new(big.Float)
// 	formatBalanceAt.SetString(balanceAt.String())
// 	balanceAtValue := new(big.Float).Quo(formatBalanceAt, big.NewFloat(math.Pow10(18)))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(balanceAt) // 25729324269165216042
// 	fmt.Println("connceted to", account, "|| Balance @ ", blockNumber, balanceAtValue)

// }

func Get_Avax_Address() common.Address {
	address := common.HexToAddress("")
	return address
}

// eventually will be a mapping that tracks all tokens
func Get_Token_Balance() (common.Address, *big.Int) {
	fmt.Println("will get token address")
	address := common.HexToAddress("0x37FAb20e8E95Abe04f7B7eA0BF9774654E3D17a7")
	fmt.Printf("will print token balance")
	balance := big.NewInt(1000000)
	return address, balance
}
