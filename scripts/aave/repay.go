package aave

import "fmt"

func Repay() {
	fmt.Println("Will interact with repay() on aave")
	// import account to send transaction
	// get accountBalance based of network and chainid

	// 1. call getLendingPool()
	// prep signer
	// 2. Get what borrowed
	// 		amount := GetBorrowUserData()
	// 3. send transaction => repay(token_address || weth_address, amount, 1, deployer)

	// print tx
}
