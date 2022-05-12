package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"go-aave/cmd/connect"
	"go-aave/erc20"
	"go-aave/pool_address_provider"

	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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

func Approve_ERC20(blockchain *ethclient.Client, loaner_address string, pk string, signer *bind.TransactOpts, spender common.Address, amount *big.Int)  { 
	fmt.Println("Approving ERC20...")
	const a_token = "0xC50E6F9E8e6CAd53c42ddCB7A42d616d7420fd3e"
	// prep signer

		// 1. connect to the aave lending pool
	 	// fiji: 0x7fdC1FdF79BE3309bf82f4abdAD9f111A6590C0f || cchain: 0x4F01AeD16D97E3aB5ab2B501154DC9bb0F1A5A2C


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
	deployer.GasLimit = uint64(300000)               // in units
	deployer.GasPrice = gasPrice
	loaner := common.HexToAddress(loaner_address)
	pool_provider_address := "0x1775ECC8362dB6CaB0c7A9C0957cF656A5276c29"
	pool_provider := common.HexToAddress(pool_provider_address)
	PoolProvider, err := pool_address_provider.NewPoolAddressProvider(pool_provider, blockchain)
	if err != nil {
		log.Fatal(err)
	}

	// Get Load Pool Contract
	pool_address, err := PoolProvider.GetPool(nil)
	fmt.Println("Pool Address:", pool_address.String())
	if err != nil {
		log.Fatal(err)
	}
	println("Current Balances...")
	connect.GetAVAXBalance(blockchain, loaner.String())
	connect.Get_aWAVAXBalance(blockchain, loaner.String())




	token, err := erc20.NewErc20(common.HexToAddress(a_token), blockchain)
	if err != nil {
		log.Fatal(err)
	}
	token_approval, err := token.Approve(deployer, pool_address, deployer.Value)
	if err != nil {
		log.Fatal(err)
	}

	println("Approving...")
	bind.WaitMined(context.Background(), blockchain, token_approval)
	println("Approved!")

	token_allowance, err := token.Allowance(nil, loaner, pool_address)
	if err != nil {
		log.Fatal(err)
	}

	println("Approving...")
	bind.WaitMined(context.Background(), blockchain, token_approval)
	println("Approved!")


	signedTx_1, err := types.SignTx(token_approval, types.NewEIP155Signer(chainID), private)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Allowance: %v!", token_allowance)

	fmt.Printf("Approve TX: https://testnet.snowtrace.io/tx/%s\n", signedTx_1.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
	_ = token

}