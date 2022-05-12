package aave

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"go-aave/aave_lending_eth"
	"go-aave/cmd/connect"

	"go-aave/pool_address_provider"
	"go-aave/weth"

	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func WithdrawAvax(blockchain *ethclient.Client, loaner_address string, pk string, amount *big.Int) {
	fmt.Printf("Withdrawing %v funds from Aave", amount)
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
	deployer.GasLimit = uint64(600000)               // in units
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
	if err != nil {
		log.Fatal(err)
	}

		// approve tokens first
	// utils.Approve_ERC20(blockchain, loaner_address, pk,  deployer, common.HexToAddress(token), amount)
	// Load WETHGateway Contract
	weth_gate_address := common.HexToAddress("0x8f57153F18b7273f9A814b93b31Cb3f9b035e7C2")
	WETHGateway, err := aave_lending_eth.NewAaveLendingEth(weth_gate_address, blockchain)
	if err != nil {
		log.Fatal(err)
	}



	wavax, err := weth.NewWeth(common.HexToAddress(token_address), blockchain)
	if err != nil {
		log.Fatal(err)
	}
	wavax_tx_approve, err := wavax.Approve(deployer, pool_address, deployer.Value)
	if err != nil {
		log.Fatal(err)
	}


	println("Approving...")
	bind.WaitMined(context.Background(), blockchain, wavax_tx_approve)
	println("Approved!")


	signedTx_1, err := types.SignTx(wavax_tx_approve, types.NewEIP155Signer(chainID), private)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Allowance: %v!" ,wavax_tx_approve)

	fmt.Printf("Approve TX: https://testnet.snowtrace.io/tx/%s\n", signedTx_1.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870


	fmt.Println("Here is the Pool Address:", pool_address.String())
	println("Current Balances...")
	connect.GetAVAXBalance(blockchain, loaner.String())
	connect.Get_aWAVAXBalance(blockchain, loaner.String())

	
	// withdrawEth
	withdraw_eth_tx, err := WETHGateway.WithdrawETH(deployer, pool_address, deployer.Value, loaner)
	if err != nil {
		log.Fatal(err)
	}

	println("Now Withdrawing")
	bind.WaitMined(context.Background(), blockchain, withdraw_eth_tx)
	println("Withdrawal Complete")
	connect.GetAVAXBalance(blockchain, loaner.String())
	connect.Get_aWAVAXBalance(blockchain, loaner.String())

	signedTx_3, err := types.SignTx(withdraw_eth_tx, types.NewEIP155Signer(chainID), private)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: https://testnet.snowtrace.io/tx/%s\n", signedTx_3.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

	_ = wavax
}

func WithdrawAllAvax() {
	fmt.Println("Withdrawing All funds from Aave")

}
