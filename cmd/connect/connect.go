package connect

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"go-aave/erc20"

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

func GetAVAXBalance(blockchain *ethclient.Client, address string) {
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
//
func Get_WAVAXBalance(blockchain *ethclient.Client, address string) (common.Address) {
	const token = "0x18eE6714Bb1796b8172951D892Fb9f42a961C812" 
	// Get Current Balance of address
	account := common.HexToAddress(address)
	tokenAddress := common.HexToAddress(token)
	token_instance, err := erc20.NewErc20(tokenAddress, blockchain)
	if err != nil {
		log.Fatal(err)
	}

	bal, err := token_instance.BalanceOf(&bind.CallOpts{}, account)
	formatBalance := new(big.Float)
	formatBalance.SetString(bal.String())
	balanceValue := new(big.Float).Quo(formatBalance, big.NewFloat(math.Pow10(18)))
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("WAVAX Balance:", balanceValue)
	return tokenAddress
}
//
func Get_WETHBalance(blockchain *ethclient.Client, address string) {
	const token = "0x28A8E6e41F84e62284970E4bc0867cEe2AAd0DA4" 
	// Get Current Balance of address
	account := common.HexToAddress(address)
	tokenAddress := common.HexToAddress(token)
	token_instance, err := erc20.NewErc20(tokenAddress, blockchain)
	if err != nil {
		log.Fatal(err)
	}

	bal, err := token_instance.BalanceOf(&bind.CallOpts{}, account)
	formatBalance := new(big.Float)
	formatBalance.SetString(bal.String())
	balanceValue := new(big.Float).Quo(formatBalance, big.NewFloat(math.Pow10(18)))
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("WETH Balance:", balanceValue)

}
//
func Get_WBTCBalance(blockchain *ethclient.Client, address string) {
	const token = "0x09C85Ef96e93f0ae892561052B48AE9DB29F2458" 
	// Get Current Balance of address
	account := common.HexToAddress(address)
	tokenAddress := common.HexToAddress(token)
	token_instance, err := erc20.NewErc20(tokenAddress, blockchain)
	if err != nil {
		log.Fatal(err)
	}

	bal, err := token_instance.BalanceOf(&bind.CallOpts{}, account)
	formatBalance := new(big.Float)
	formatBalance.SetString(bal.String())
	balanceValue := new(big.Float).Quo(formatBalance, big.NewFloat(math.Pow10(18)))
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("WBTC Balance:", balanceValue)

}

func GetDAIBalance(blockchain *ethclient.Client, address string) {
	const Fuji_DAI = "0xFc7215C9498Fc12b22Bc0ed335871Db4315f03d3" 
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
	const Fuji_USDC = "0x3E937B4881CBd500d05EeDAB7BA203f2b7B3f74f" 
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
	const Fuji_USDT = "0xD90db1ca5A6e9873BCD9B0279AE038272b656728" 
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

//
func Get_LINKBalance(blockchain *ethclient.Client, address string) {
	const token = "0x73b4C0C45bfB90FC44D9013FA213eF2C2d908D0A" 
	// Get Current Balance of address
	account := common.HexToAddress(address)
	tokenAddress := common.HexToAddress(token)
	token_instance, err := erc20.NewErc20(tokenAddress, blockchain)
	if err != nil {
		log.Fatal(err)
	}

	bal, err := token_instance.BalanceOf(&bind.CallOpts{}, account)
	formatBalance := new(big.Float)
	formatBalance.SetString(bal.String())
	balanceValue := new(big.Float).Quo(formatBalance, big.NewFloat(math.Pow10(18)))
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("LINK Balance:", balanceValue)

}


func Get_aUSDTBalance(blockchain *ethclient.Client, address string) {
	const token = "0x3a7e85a86F952CB61485e2D20BDDb6e15204744f" 
	// Get Current Balance of address
	account := common.HexToAddress(address)
	tokenAddress := common.HexToAddress(token)
	token_instance, err := erc20.NewErc20(tokenAddress, blockchain)
	if err != nil {
		log.Fatal(err)
	}

	bal, err := token_instance.BalanceOf(&bind.CallOpts{}, account)
	formatBalance := new(big.Float)
	formatBalance.SetString(bal.String())
	balanceValue := new(big.Float).Quo(formatBalance, big.NewFloat(math.Pow10(6)))
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("aUSDT Balance:", balanceValue)

}

func Get_aUSDCBalance(blockchain *ethclient.Client, address string) {
	const token = "0xA79570641bC9cbc6522aA80E2de03bF9F7fd123a" 
	// Get Current Balance of address
	account := common.HexToAddress(address)
	tokenAddress := common.HexToAddress(token)
	token_instance, err := erc20.NewErc20(tokenAddress, blockchain)
	if err != nil {
		log.Fatal(err)
	}

	bal, err := token_instance.BalanceOf(&bind.CallOpts{}, account)
	formatBalance := new(big.Float)
	formatBalance.SetString(bal.String())
	balanceValue := new(big.Float).Quo(formatBalance, big.NewFloat(math.Pow10(6)))
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("aUSDC Balance:", balanceValue)

}

func Get_aDAIBalance(blockchain *ethclient.Client, address string) {
	const token = "0xC42f40B7E22bcca66B3EE22F3ACb86d24C997CC2" 
	// Get Current Balance of address
	account := common.HexToAddress(address)
	tokenAddress := common.HexToAddress(token)
	token_instance, err := erc20.NewErc20(tokenAddress, blockchain)
	if err != nil {
		log.Fatal(err)
	}

	bal, err := token_instance.BalanceOf(&bind.CallOpts{}, account)
	formatBalance := new(big.Float)
	formatBalance.SetString(bal.String())
	balanceValue := new(big.Float).Quo(formatBalance, big.NewFloat(math.Pow10(18)))
	if err != nil {
		log.Fatal(err)
	}
	
	
	fmt.Println("aDai Balance:", balanceValue)

}

func Get_aWAVAXBalance(blockchain *ethclient.Client, address string) {
	const token = "0xC50E6F9E8e6CAd53c42ddCB7A42d616d7420fd3e" 
	// Get Current Balance of address
	account := common.HexToAddress(address)
	tokenAddress := common.HexToAddress(token)
	token_instance, err := erc20.NewErc20(tokenAddress, blockchain)
	if err != nil {
		log.Fatal(err)
	}

	bal, err := token_instance.BalanceOf(&bind.CallOpts{}, account)
	formatBalance := new(big.Float)
	formatBalance.SetString(bal.String())
	balanceValue := new(big.Float).Quo(formatBalance, big.NewFloat(math.Pow10(18)))
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("aWAVAX Balance:", balanceValue)

}

func Get_aLINKBalance(blockchain *ethclient.Client, address string) {
	const token = "0x210a3f864812eAF7f89eE7337EAA1FeA1830C57e" 
	// Get Current Balance of address
	account := common.HexToAddress(address)
	tokenAddress := common.HexToAddress(token)
	token_instance, err := erc20.NewErc20(tokenAddress, blockchain)
	if err != nil {
		log.Fatal(err)
	}

	bal, err := token_instance.BalanceOf(&bind.CallOpts{}, account)
	formatBalance := new(big.Float)
	formatBalance.SetString(bal.String())
	balanceValue := new(big.Float).Quo(formatBalance, big.NewFloat(math.Pow10(18)))
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("aLINK Balance:", balanceValue)

}

func Get_aWETHBalance(blockchain *ethclient.Client, address string) {
	const token = "0x618922b15a1a92652818473741531eE255f68741" 
	// Get Current Balance of address
	account := common.HexToAddress(address)
	tokenAddress := common.HexToAddress(token)
	token_instance, err := erc20.NewErc20(tokenAddress, blockchain)
	if err != nil {
		log.Fatal(err)
	}

	bal, err := token_instance.BalanceOf(&bind.CallOpts{}, account)
	formatBalance := new(big.Float)
	formatBalance.SetString(bal.String())
	balanceValue := new(big.Float).Quo(formatBalance, big.NewFloat(math.Pow10(18)))
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("aWETH Balance:", balanceValue)

}

func Get_aWBTCBalance(blockchain *ethclient.Client, address string) {
	const token = "0x07B2C0b69c70e89C94A20A555Ab376E5a6181eE6" 
	// Get Current Balance of address
	account := common.HexToAddress(address)
	tokenAddress := common.HexToAddress(token)
	token_instance, err := erc20.NewErc20(tokenAddress, blockchain)
	if err != nil {
		log.Fatal(err)
	}

	bal, err := token_instance.BalanceOf(&bind.CallOpts{}, account)
	formatBalance := new(big.Float)
	formatBalance.SetString(bal.String())
	balanceValue := new(big.Float).Quo(formatBalance, big.NewFloat(math.Pow10(18)))
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("aWBTC Balance:", balanceValue)

}





