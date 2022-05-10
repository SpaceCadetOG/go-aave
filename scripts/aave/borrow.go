package aave

import (
	"fmt"
	// "go-aave/scripts/utils"
)


func Borrow() {
	fmt.Println("Will interact with borrow() on aave")
	// 1. call getLendingPool()
	// 2. avalBorrowsEth, TotalDebtEth
	// 3. Get UserBorrow Stats => call getBorrowUserData(lending_pool, user_address)
	// 4. Get Price of DAI or other token
	// 5. Get AmountToBorrowInDai	
	// utils.Approve_ERC20()
	// 6. send transaction => borrow(token_address || weth_address, amount_to_borrow, 1, 0, deployer)
	// 7. Get UserBorrow Stats to see update => call getBorrowUserData(lending_pool, user_address)


}
