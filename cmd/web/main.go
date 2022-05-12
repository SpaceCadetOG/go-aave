package main

import (
	// "context"
	// "crypto/ecdsa"


	"go-aave/scripts/aave"
	"math/big"

	"log"

	// "github.com/ethereum/go-ethereum/accounts/abi/bind"
	// "github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
)

/*
// borrow from aave using go

step 0 => get Wrapped or Check balance
	Check token Balance for both wrapped eth
	we need to get WETH or WAVAX()

steps 1 Deposit
	we need to get the account that will be sending the transaction
	We must connect to the aave pool
	We need to Deposit

steps 2 Borrow
	we need to get the account that will be sending the transaction
	we need to get borrow Data(How much we can borrow)
	we need to price of DAI or Other
	We must borrow DAI
	We need to Deposit

steps 3 Repay
	we need to get the account that will be sending the transaction
	we need to get borrow Data(How Much We Owe)
	we need approve tokens
	We must connect to the aave pool
	We need to Deposit

Global Steps
	we need to get the account that will be sending the transaction
	Approve Tokens
	getLendingPool()
*/
// type Deposits struct {
// 	token common.Address
// 	amount        *big.Int
// 	loaner        common.Address
// }

const pk = "7aa1eea0195e3755bd6949ade6686fd8a3e8adcefa7c6797c0efbf473e88322e"
func main() {
	const fuji = "wss://speedy-nodes-nyc.moralis.io/8f7573e4ff0324155dcbec2e/avalanche/testnet/ws"	

	blockchain, err := ethclient.Dial(fuji) // ganache
	if err != nil {
		log.Fatal(err)
	}

	// signer := connect.GetSigner(blockchain, "7aa1eea0195e3755bd6949ade6686fd8a3e8adcefa7c6797c0efbf473e88322e")
	loaner_address := "0x5AC42D67bab747677FD5B5156258bB65fFB1e275"
	loaner := common.HexToAddress(loaner_address)


	amount := big.NewInt(1000000000000000000)


	// // get pool address
	// 	// Load Pool Contract
	// 	Pool, err := aave_pool.NewAavePool(pool_address, blockchain)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// depositEth

	
	// connect.GetDAIBalance(blockchain, loaner.String())
	// connect.GetUSDCBalance(blockchain, loaner.String())

	// utils.GetWETH_Address()
	
	// Deposit Avax || Withdraw Avax
	// aave.DepositAVAX(blockchain, loaner.String(), pk)
		aave.WithdrawAvax(blockchain, loaner.String(), pk, amount)


	// aave.DepositDAI(blockchain, signer, amount, loaner) // signer *bind.TransactOpts, amount *big.Int, loaner common.Address
	// aave.DepositUSDC(signer, amount, loaner) // signer *bind.TransactOpts, amount *big.Int, loaner common.Address
	// aave.DepositUSDT(signer, amount, loaner) // signer *bind.TransactOpts, amount *big.Int, loaner common.Address
	// aave.Borrow()
	// aave.Repay()
}

/*
Store Contract Address 0x3AA7cEe3820f116A5eCaD62340F03fF50Ca5f973
Store Contract Deployed to 0x4c42fdf3e45c435ac667d849c1b27e5f3e8467bf792f2dade75c2810a99730b0
*/
