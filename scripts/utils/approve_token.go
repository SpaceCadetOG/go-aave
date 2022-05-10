package utils

import (
	"fmt"
	"go-aave/cmd/connect"
	"go-aave/token/erc20"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func Approve_ERC20(signer *bind.TransactOpts, spender common.Address, amount *big.Int)  { 
	fmt.Println("Approving ERC20")
	// prep signer
	chain := connect.ConnectSigner()
	
		// 1. connect to the aave lending pool
	 	// fiji: 0x7fdC1FdF79BE3309bf82f4abdAD9f111A6590C0f || cchain: 0x4F01AeD16D97E3aB5ab2B501154DC9bb0F1A5A2C
	

		token, err := erc20.NewErc20(spender, chain)
		if err != nil {
			log.Fatal(err)
		}
		token_approval, err := token.Approve(signer, spender, amount)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Token Approval", token_approval)

		_ = token
}