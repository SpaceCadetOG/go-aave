package aave

import (
	"go-aave/aave"
	"go-aave/cmd/connect"
	"log"

	"github.com/ethereum/go-ethereum/common"
)

func GetLendingPool()  {
	chain := connect.ConnectSigner()
	pool_address_provider := common.HexToAddress("0x1775ECC8362dB6CaB0c7A9C0957cF656A5276c29")
	getLendingPool_Instance, err := aave.NewAave(pool_address_provider, chain)
	if err != nil {
		log.Fatal(err)
	}
	_ = getLendingPool_Instance

}

func GetBorrowUserData() uint {
	//  totalCollateralETH, totalDebtETH, availableBorrowsETH  := getUserAccountData(account)
	//  return { availableBorrowsETH, totalDebtETH }
	return 157
}