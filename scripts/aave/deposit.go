package aave

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"go-aave/aave_lending_eth"
	"go-aave/cmd/connect"
	"go-aave/pool_address_provider"

	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func DepositAVAX(blockchain *ethclient.Client, loaner_address string, pk string) {
	// signer := connect.GetSigner(blockchain, pk)
	// // loaner_address := "0x5AC42D67bab747677FD5B5156258bB65fFB1e275"

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

	// Load WETHGateway Contract
	weth_gate_address := common.HexToAddress("0x8f57153F18b7273f9A814b93b31Cb3f9b035e7C2")
	WETHGateway, err := aave_lending_eth.NewAaveLendingEth(weth_gate_address, blockchain)
	if err != nil {
		log.Fatal(err)
	}

	// depositEth
	deposit_eth_tx, err := WETHGateway.DepositETH(deployer, pool_address, loaner, 0)
	if err != nil {
		log.Fatal(err)
	}
	println("Now Depositing")
	bind.WaitMined(context.Background(), blockchain, deposit_eth_tx)
	println("Deposit Complete")
	connect.GetAVAXBalance(blockchain, loaner.String())
	connect.Get_aWAVAXBalance(blockchain, loaner.String())

	signedTx, err := types.SignTx(deposit_eth_tx, types.NewEIP155Signer(chainID), private)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: https://testnet.snowtrace.io/tx/%s\n", signedTx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

}
