package main

import (
	// "context"
	// "crypto/ecdsa"
	//"go-aave/scripts/aave"
	"go-aave/cmd/connect"
	"go-aave/scripts/aave"

	"log"
	"math/big"

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

const pk = "0x83dd616d3db5da014aeea76a32cce1ae57c4a5f89be0a9eb5f1acb6b519e1a03"
func main() {
	const fuji = "https://api.avax-test.network/ext/bc/C/rpc"	

	blockchain, err := ethclient.Dial(fuji) // ganache
	if err != nil {
		log.Fatal(err)
	}
	signer := connect.GetSigner(blockchain, pk)
	amount := big.NewInt(1)
	loaner_address := "0x1Dd092B79C0c340A0EBd9B90ec190D540091Ed24"
	loaner := common.HexToAddress(loaner_address)


	// // 2. Connect any account to make the transaction to transfer eth
	// // Load Private Keys
	// private, err := crypto.HexToECDSA("0520d203bb14d44929b0759412c14741b1ae38dfc566560629ddb1bcea906db7")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// public := private.Public()
	// publicECDSA, ok := public.(*ecdsa.PublicKey)
	// if !ok {
	// 	log.Fatal("error casting public ket to EDCSA")
	// }

	// fromAddress := crypto.PubkeyToAddress(*publicECDSA)

	// nonce, err := blockchain.PendingNonceAt(context.Background(), fromAddress)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// gasPrice, err := blockchain.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // this will do the transaction set for you
	// deployer := bind.NewKeyedTransactor(private)
	// deployer.Nonce = big.NewInt(int64(nonce))
	// deployer.Value = big.NewInt(0)     // in wei
	// deployer.GasLimit = uint64(300000) // in units
	// deployer.GasPrice = gasPrice


	// utils.Get_AVAX_Balance(blockchain, loaner.String())
	// connect.GetDAIBalance(blockchain, loaner.String())
	// connect.GetUSDCBalance(blockchain, loaner.String())
	// utils.GetWETH_Address()
	
	aave.DepositAVAX(signer, amount, loaner) // signer *bind.TransactOpts, amount *big.Int, loaner common.Address
	// aave.DepositDAI(signer, amount, loaner) // signer *bind.TransactOpts, amount *big.Int, loaner common.Address
	// aave.DepositUSDC(signer, amount, loaner) // signer *bind.TransactOpts, amount *big.Int, loaner common.Address
	// aave.DepositUSDT(signer, amount, loaner) // signer *bind.TransactOpts, amount *big.Int, loaner common.Address
	// aave.Borrow()
	// aave.Repay()
}

/*
Store Contract Address 0x3AA7cEe3820f116A5eCaD62340F03fF50Ca5f973
Store Contract Deployed to 0x4c42fdf3e45c435ac667d849c1b27e5f3e8467bf792f2dade75c2810a99730b0
*/
