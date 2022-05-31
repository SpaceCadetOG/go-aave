package aave

import (
	"context"


	"fmt"
	"go-aave/aave_lending_eth"
	"go-aave/aave_pool"
	"go-aave/cmd/connect"
	"go-aave/pool_address_provider"

	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"


	"github.com/ethereum/go-ethereum/ethclient"
)

// blockchain *ethclient.Client, deployer *bind.TransactOpts, token_address common.Address

func DepositAVAX(blockchain *ethclient.Client, deployer *bind.TransactOpts) {
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
	println("Avax Current Balances...")
	connect.GetAVAXBalance(blockchain, deployer.From.String())
	println("aWAvax Current Balances...")
	connect.Get_aWAVAXBalance(blockchain, deployer.From.String())

	// Load WETHGateway Contract
	weth_gate_address := common.HexToAddress("0x8f57153F18b7273f9A814b93b31Cb3f9b035e7C2")
	WETHGateway, err := aave_lending_eth.NewAaveLendingEth(weth_gate_address, blockchain)
	if err != nil {
		log.Fatal(err)
	}

	// depositEth
	deposit_eth_tx, err := WETHGateway.DepositETH(deployer, pool_address, deployer.From, 0)
	if err != nil {
		log.Fatal(err)
	}
	println("Now Depositing AVAX")
	bind.WaitMined(context.Background(), blockchain, deposit_eth_tx)
	println("AVAX Deposit Complete")
	connect.GetAVAXBalance(blockchain, deployer.From.String())
	connect.Get_aWAVAXBalance(blockchain, deployer.From.String())

	// signedTx, err := types.SignTx(deposit_eth_tx, types.NewEIP155Signer(chainID), private)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Printf("tx sent: https://testnet.snowtrace.io/tx/%s\n", deposit_eth_tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

}

func DepositDAI(blockchain *ethclient.Client, deployer *bind.TransactOpts, token_address common.Address) {
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

	fmt.Println("Pool Address:", pool_address.String())

	println("Dai Current Balances...")
	connect.GetDAIBalance(blockchain, deployer.From.String())
	println("aDai Current Balances...")
	connect.Get_aDAIBalance(blockchain, deployer.From.String())

	pool, err := aave_pool.NewAavePool(pool_address, blockchain)
	if err != nil {
		log.Fatal(err)
	}

	// depositEth
	deposit_dai_tx, err := pool.Supply(deployer, token_address, deployer.Value, deployer.From, 0)
	//supply_dai_tx, err := pool.SupplyWithPermit(deployer, token_address, deployer.Value, deployer.From, 0, big.NewInt(1), 8, data, hash)
	if err != nil {
		log.Fatal(err)
	}
	println("Now Depositing Dai")
	go bind.WaitMined(context.Background(), blockchain, deposit_dai_tx)
	println("Dai Deposit Complete")

	fmt.Printf("tx sent: https://testnet.snowtrace.io/tx/%s\n", deposit_dai_tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
	connect.GetDAIBalance(blockchain, deployer.From.String())
	connect.Get_aDAIBalance(blockchain, deployer.From.String())
}
