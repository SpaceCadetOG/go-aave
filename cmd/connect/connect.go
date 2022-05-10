package connect

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"go-aave/token/erc20"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const fuji = "https://api.avax-test.network/ext/bc/C/rpc"

func ConnectSigner() *ethclient.Client {
	blockchain, err := ethclient.Dial(fuji) // ganache
	if err != nil {
		log.Fatal(err)
	}

	return blockchain
}

func GetSigner(blockchain *ethclient.Client, pk string) *bind.TransactOpts {

		// 2. Connect any account to make the transaction to transfer eth
	// Load Private Keys
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

	nonce, err := blockchain.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := blockchain.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// this will do the transaction set for you
	deployer := bind.NewKeyedTransactor(private)
	deployer.Nonce = big.NewInt(int64(nonce))
	deployer.Value = big.NewInt(0)     // in wei
	deployer.GasLimit = uint64(300000) // in units
	deployer.GasPrice = gasPrice

	return deployer
}

func GetBalance(blockchain *ethclient.Client, address string) {
		// Get Current Balance of address
		account := common.HexToAddress(address)

		balance, err := blockchain.BalanceAt(context.Background(), account, nil)
		formatBalance := new(big.Float)
		formatBalance.SetString(balance.String())
		balanceValue := new(big.Float).Quo(formatBalance, big.NewFloat(math.Pow10(18)))
		if err != nil {
			log.Fatal(err)
		}

	fmt.Println("AVAX Balance:", balanceValue)

}



func GetDAIBalance(blockchain *ethclient.Client, address string) {
	const Fuji_DAI = "0x63E537A69b3f5B03F4f46c5765c82861BD874b6e" 
	// Get Current Balance of address
	account := common.HexToAddress(address)
	daiAddress := common.HexToAddress(Fuji_DAI)
	dai_instance, err := erc20.NewErc20(daiAddress, blockchain)
	if err != nil {
		log.Fatal(err)
	}

	bal, err := dai_instance.BalanceOf(&bind.CallOpts{}, account)
	formatBalance := new(big.Float)
	formatBalance.SetString(bal.String())
	balanceValue := new(big.Float).Quo(formatBalance, big.NewFloat(math.Pow10(18)))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DAI Balance:", balanceValue)

}

func GetUSDCBalance(blockchain *ethclient.Client, address string) {
	const Fuji_USDC = "0x02444D214962eC73ab733bB00Ca98879efAAa73d" 
	// Get Current Balance of address
	account := common.HexToAddress(address)
	usdcAddress := common.HexToAddress(Fuji_USDC)
	usdc_instance, err := erc20.NewErc20(usdcAddress, blockchain)
	if err != nil {
		log.Fatal(err)
	}

	bal, err := usdc_instance.BalanceOf(&bind.CallOpts{}, account)
	formatBalance := new(big.Float)
	formatBalance.SetString(bal.String())
	balanceValue := new(big.Float).Quo(formatBalance, big.NewFloat(math.Pow10(6)))
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("USDC Balance:", balanceValue)

}

func GetUSDTBalance(blockchain *ethclient.Client, address string) {
	const Fuji_USDT = "0x18eE6714Bb1796b8172951D892Fb9f42a961C812" 
	// Get Current Balance of address
	account := common.HexToAddress(address)
	usdtAddress := common.HexToAddress(Fuji_USDT)
	usdt_instance, err := erc20.NewErc20(usdtAddress, blockchain)
	if err != nil {
		log.Fatal(err)
	}

	bal, err := usdt_instance.BalanceOf(&bind.CallOpts{}, account)
	formatBalance := new(big.Float)
	formatBalance.SetString(bal.String())
	balanceValue := new(big.Float).Quo(formatBalance, big.NewFloat(math.Pow10(6)))
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("USDT Balance:", balanceValue)

}


