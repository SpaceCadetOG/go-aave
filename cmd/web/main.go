package main

import (
	// "context"
	// "crypto/ecdsa"

	"context"
	"crypto/ecdsa"
	"fmt"
	"go-aave/scripts/aave"
	

	"go-aave/pool_address_provider"

	"math/big"

	"log"

	// "github.com/ethereum/go-ethereum/accounts/abi/bind"
	// "github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

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


	const token_address = "0x18eE6714Bb1796b8172951D892Fb9f42a961C812"
	const a_token = "0xC50E6F9E8e6CAd53c42ddCB7A42d616d7420fd3e"

	// import account to send transaction
	private, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Fatal(err)
	}

	public := private.Public()
	publicECDSA, ok := public.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public ket to EDCSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicECDSA)
	fmt.Println(fromAddress.String())

	nonce, err := blockchain.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := blockchain.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chainID := big.NewInt(43113)

	// this will do the transaction set for you
	deployer, err := bind.NewKeyedTransactorWithChainID(private, chainID)
	if err != nil {
		log.Fatal(err)
	}
	deployer.Nonce = big.NewInt(int64(nonce))
	deployer.Value = big.NewInt(1000000000000000000) // in wei
	deployer.GasLimit = uint64(300000)             // in units
	deployer.GasPrice = gasPrice
	pool_provider_address := "0x1775ECC8362dB6CaB0c7A9C0957cF656A5276c29"
	pool_provider := common.HexToAddress(pool_provider_address)
	PoolProvider, err := pool_address_provider.NewPoolAddressProvider(pool_provider, blockchain)
	if err != nil {
		log.Fatal(err)
	}
	// Get Load Pool Contract
	pool_address, err := PoolProvider.GetPool(nil)
	if err != nil {
		log.Fatal(err)
	}

	// approve tokens first
	// utils.Approve_ERC20(blockchain, loaner_address, pk,  deployer, common.HexToAddress(token), amount)
	// Load WETHGateway Contract
	//	weth_gate_address := common.HexToAddress("0x8f57153F18b7273f9A814b93b31Cb3f9b035e7C2")
	// WETHGateway, err := aave_lending_eth.NewAaveLendingEth(weth_gate_address, blockchain)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	dai := common.HexToAddress("0xFc7215C9498Fc12b22Bc0ed335871Db4315f03d")
	fmt.Println("Here is the Pool Address:", pool_address.String())
	// go utils.Approve_ERC20(blockchain, deployer, pool_address, dai)



	// connect.GetDAIBalance(blockchain, loaner.String())
	// connect.GetUSDCBalance(blockchain, loaner.String())

	// Deposit Avax || Withdraw Avax
	aave.DepositDAI(blockchain, deployer, dai)
	// aave.DepositAVAX(blockchain, deployer)

	//	aave.WithdrawAllAvax(blockchain, deployer, dai)

	// aave.DepositDAI(blockchain, signer, amount, loaner) // signer *bind.TransactOpts, amount *big.Int, loaner common.Address
	// aave.DepositUSDC(signer, amount, loaner) // signer *bind.TransactOpts, amount *big.Int, loaner common.Address
	// aave.DepositUSDT(signer, amount, loaner) // signer *bind.TransactOpts, amount *big.Int, loaner common.Address
	// aave.Borrow()
	// aave.Repay()

	// utils.Approve_ERC20(blockchain, deployer, pool_address, WAVAX)
	// utils.Approve_ERC20(blockchain, deployer.From.String(), pk, common.HexToAddress("0x8f57153f18b7273f9a814b93b31cb3f9b035e7c2"), "0xC50E6F9E8e6CAd53c42ddCB7A42d616d7420fd3e", *amount)

}

/*
Store Contract Address 0x3AA7cEe3820f116A5eCaD62340F03fF50Ca5f973
Store Contract Deployed to 0x4c42fdf3e45c435ac667d849c1b27e5f3e8467bf792f2dade75c2810a99730b0
*/
