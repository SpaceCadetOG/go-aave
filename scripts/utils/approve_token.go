package utils

import (
	"context"
	"fmt"
	"go-aave/erc20"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
)

// require approve tokens
/*
owner:
0x219a57b14562bd77919e1ed12cb0b0b1650d8e2e

spender:
0xC4744c984975ab7d41e0dF4B37E048Ef8006115E // lendingPool

value:
1000000000000000000

nonce:
0

deadline:
115792089237316195423570985008687907853269984665640564039457584007913129639935

*/

func Approve_ERC20(blockchain *ethclient.Client,  deployer *bind.TransactOpts, spender common.Address, token_address common.Address) {

	erc20Contract, err := erc20.NewErc20(token_address, blockchain)
	if err != nil {
		log.Fatal(err)
	}
	tx, err := erc20Contract.Approve(deployer, spender, deployer.Value)
	if err != nil {
		log.Fatal(err)
	}

	
	println("Approving...")
	bind.WaitDeployed(context.Background(), blockchain, tx)
	fmt.Printf("Approved! tx sent: https://testnet.snowtrace.io/tx/%s\n", tx.Hash().Hex()) 

}       	