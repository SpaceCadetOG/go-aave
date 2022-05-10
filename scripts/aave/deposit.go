package aave

import (
	"fmt"
	"go-aave/aave"
	link "go-aave/chainlink"
	"go-aave/cmd/connect"
	"go-aave/scripts/utils"

	// "go-aave/scripts/utils"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// For now this will only deposit weth => soon we will scale to stables

/*
 opts *bind.TransactOpts = signer

 asset common.Address = token_address

 amount *big.Int = amount to send

 onBehalfOf common.Address = the address to recieved

 referralCode uint16 = 0

 (*types.Transaction, error) = nil
*/

func DepositAVAX(signer *bind.TransactOpts, amount *big.Int, loaner common.Address) {
	fmt.Println("Will deposit() AVAX to aave")
	// prep signer
	
	chain := connect.ConnectSigner()

	// 1. connect to the aave lending pool
	
	lending_address := common.HexToAddress("0x1775ECC8362dB6CaB0c7A9C0957cF656A5276c29") // fiji: 0x1775ECC8362dB6CaB0c7A9C0957cF656A5276c29 || cchain: 0x4F01AeD16D97E3aB5ab2B501154DC9bb0F1A5A2C
	aave_instance, err := aave.NewAave(lending_address, chain)
	if err != nil {
		log.Fatal(err)
	}
	_ = aave_instance


	link_agg_address := common.HexToAddress("") // fuji:  || cchain:
	link_instance, err := link.NewLink(link_agg_address, chain)
	if err != nil {
		log.Fatal(err)
	}
	_ = link_instance

	// 2. approve wethAddress to allow the lending pool to recieve tokens + amount
	loaner = common.HexToAddress(loaner.String())
	// utils.Approve_ERC20(AVAX) // or GetWeth

	// Print user balance
	// fmt.Printf("WETH Address Balance: %d", balance)
	utils.Get_AVAX_Balance(chain, loaner.String())
	utils.Approve_ERC20(signer, lending_address, amount) // spender address, tokens uint256
	// https://goethereumbook.org/transfer-tokens/

	// 3. call deposit() + send transaction => deposit(token_address || weth_address, amount, deployer, 0)
	// deposit_tx, err := aave_instance.Supply(signer, AVAX, amount, loaner, 0)
	// // Deposit(signer, weth_address, amount, loaner, 0)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println("Depositing + pending balance")
	// fmt.Printf("Depositing + pending balance: %d", deposit_tx)
	// print deposit reciept on etherscan
	// fmt.Printf("Deposited(etherscan) + token balance:")

	fmt.Println("Deposited + balance")

	// parms address, client


}

func DepositDAI(signer *bind.TransactOpts, amount *big.Int, loaner common.Address) {
	fmt.Println("Will deposit() DAI to aave")
	// prep signer
	chain := connect.ConnectSigner()
	// 1. connect to the aave lending pool
	lending_address := common.HexToAddress("0x7fdC1FdF79BE3309bf82f4abdAD9f111A6590C0f") // fiji: 0x7fdC1FdF79BE3309bf82f4abdAD9f111A6590C0f || cchain: 0x4F01AeD16D97E3aB5ab2B501154DC9bb0F1A5A2C
	aave_instance, err := aave.NewAave(lending_address, chain)
	if err != nil {
		log.Fatal(err)
	}
	_ = aave_instance

	link_agg_address := common.HexToAddress("") // fuji:  || cchain:
	link_instance, err := link.NewLink(link_agg_address, chain)
	if err != nil {
		log.Fatal(err)
	}
	_ = link_instance

	// 2. approve wethAddress to allow the lending pool to recieve tokens + amount
	loaner = common.HexToAddress(loaner.String())
	// utils.Approve_ERC20(AVAX) // or GetWeth

	// Print user balance
	// fmt.Printf("WETH Address Balance: %d", balance)
	connect.GetDAIBalance(chain, loaner.String())


	// 3. call deposit() + send transaction => deposit(token_address || weth_address, amount, deployer, 0)
	// deposit_tx, err := aave_instance.Supply(signer, AVAX, amount, loaner, 0)
	// // Deposit(signer, weth_address, amount, loaner, 0)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// Print(Depositing + pending balance)
	// fmt.Printf("Depositing + pending balance: %d", deposit_tx)
	// print deposit reciept on etherscan
	// fmt.Printf("Deposited(etherscan) + token balance:")

	// s
	// Print(Deposited + balance)

	// parms address, client

}

func DepositUSDC(signer *bind.TransactOpts, amount *big.Int, loaner common.Address) {
	fmt.Println("Will deposit() USDC to aave")
	// prep signer
	chain := connect.ConnectSigner()
	// 1. connect to the aave lending pool
	lending_address := common.HexToAddress("0x7fdC1FdF79BE3309bf82f4abdAD9f111A6590C0f") // fiji: 0x7fdC1FdF79BE3309bf82f4abdAD9f111A6590C0f || cchain: 0x4F01AeD16D97E3aB5ab2B501154DC9bb0F1A5A2C
	aave_instance, err := aave.NewAave(lending_address, chain)
	if err != nil {
		log.Fatal(err)
	}
	_ = aave_instance

	link_agg_address := common.HexToAddress("") // fuji:  || cchain:
	link_instance, err := link.NewLink(link_agg_address, chain)
	if err != nil {
		log.Fatal(err)
	}
	_ = link_instance

	// 2. approve wethAddress to allow the lending pool to recieve tokens + amount
	loaner = common.HexToAddress(loaner.String())
	// utils.Approve_ERC20(AVAX) // or GetWeth

	// Print user balance
	// fmt.Printf("WETH Address Balance: %d", balance)
	connect.GetUSDCBalance(chain, loaner.String())


	// 3. call deposit() + send transaction => deposit(token_address || weth_address, amount, deployer, 0)
	// deposit_tx, err := aave_instance.Supply(signer, AVAX, amount, loaner, 0)
	// // Deposit(signer, weth_address, amount, loaner, 0)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// Print(Depositing + pending balance)
	// fmt.Printf("Depositing + pending balance: %d", deposit_tx)
	// print deposit reciept on etherscan
	// fmt.Printf("Deposited(etherscan) + token balance:")

	// s
	// Print(Deposited + balance)

	// parms address, client



}

func DepositUSDT(signer *bind.TransactOpts, amount *big.Int, loaner common.Address) {
	fmt.Println("Will deposit() USDT to aave")
	// prep signer
	chain := connect.ConnectSigner()
	// 1. connect to the aave lending pool
	lending_address := common.HexToAddress("0x7fdC1FdF79BE3309bf82f4abdAD9f111A6590C0f") // fiji: 0x7fdC1FdF79BE3309bf82f4abdAD9f111A6590C0f || cchain: 0x4F01AeD16D97E3aB5ab2B501154DC9bb0F1A5A2C
	aave_instance, err := aave.NewAave(lending_address, chain)
	if err != nil {
		log.Fatal(err)
	}
	_ = aave_instance

	link_agg_address := common.HexToAddress("") // fuji:  || cchain:
	link_instance, err := link.NewLink(link_agg_address, chain)
	if err != nil {
		log.Fatal(err)
	}
	_ = link_instance

	// 2. approve wethAddress to allow the lending pool to recieve tokens + amount
	loaner = common.HexToAddress(loaner.String())
	// utils.Approve_ERC20(AVAX) // or GetWeth

	// Print user balance
	// fmt.Printf("WETH Address Balance: %d", balance)
	connect.GetUSDTBalance(chain, loaner.String())


	// 3. call deposit() + send transaction => deposit(token_address || weth_address, amount, deployer, 0)
	// deposit_tx, err := aave_instance.Supply(signer, AVAX, amount, loaner, 0)
	// // Deposit(signer, weth_address, amount, loaner, 0)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// Print(Depositing + pending balance)
	// fmt.Printf("Depositing + pending balance: %d", deposit_tx)
	// print deposit reciept on etherscan
	// fmt.Printf("Deposited(etherscan) + token balance:")

	// s
	// Print(Deposited + balance)

	// parms address, client



}

