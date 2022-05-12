// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aave_pool

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DataTypesEModeCategory is an auto generated low-level Go binding around an user-defined struct.
type DataTypesEModeCategory struct {
	Ltv                  uint16
	LiquidationThreshold uint16
	LiquidationBonus     uint16
	PriceSource          common.Address
	Label                string
}

// DataTypesReserveConfigurationMap is an auto generated low-level Go binding around an user-defined struct.
type DataTypesReserveConfigurationMap struct {
	Data *big.Int
}

// DataTypesReserveData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesReserveData struct {
	Configuration               DataTypesReserveConfigurationMap
	LiquidityIndex              *big.Int
	CurrentLiquidityRate        *big.Int
	VariableBorrowIndex         *big.Int
	CurrentVariableBorrowRate   *big.Int
	CurrentStableBorrowRate     *big.Int
	LastUpdateTimestamp         *big.Int
	Id                          uint16
	ATokenAddress               common.Address
	StableDebtTokenAddress      common.Address
	VariableDebtTokenAddress    common.Address
	InterestRateStrategyAddress common.Address
	AccruedToTreasury           *big.Int
	Unbacked                    *big.Int
	IsolationModeTotalDebt      *big.Int
}

// DataTypesUserConfigurationMap is an auto generated low-level Go binding around an user-defined struct.
type DataTypesUserConfigurationMap struct {
	Data *big.Int
}

// AavePoolMetaData contains all meta data concerning the AavePool contract.
var AavePoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"backer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"BackUnbacked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumDataTypes.InterestRateMode\",\"name\":\"interestRateMode\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumDataTypes.InterestRateMode\",\"name\":\"interestRateMode\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"FlashLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDebt\",\"type\":\"uint256\"}],\"name\":\"IsolationModeTotalDebtUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralAsset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"debtAsset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtToCover\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidatedCollateralAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"receiveAToken\",\"type\":\"bool\"}],\"name\":\"LiquidationCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"MintUnbacked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountMinted\",\"type\":\"uint256\"}],\"name\":\"MintedToTreasury\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"RebalanceStableBorrowRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"repayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"useATokens\",\"type\":\"bool\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"variableBorrowRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidityIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"variableBorrowIndex\",\"type\":\"uint256\"}],\"name\":\"ReserveDataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"ReserveUsedAsCollateralDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"ReserveUsedAsCollateralEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"Supply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumDataTypes.InterestRateMode\",\"name\":\"interestRateMode\",\"type\":\"uint8\"}],\"name\":\"SwapBorrowRateMode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"categoryId\",\"type\":\"uint8\"}],\"name\":\"UserEModeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESSES_PROVIDER\",\"outputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BRIDGE_PROTOCOL_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FLASHLOAN_PREMIUM_TOTAL\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FLASHLOAN_PREMIUM_TO_PROTOCOL\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_NUMBER_RESERVES\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_STABLE_RATE_BORROW_SIZE_PERCENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"backUnbacked\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRateMode\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"ltv\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationThreshold\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationBonus\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"priceSource\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"internalType\":\"structDataTypes.EModeCategory\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"configureEModeCategory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"dropReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceFromBefore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceToBefore\",\"type\":\"uint256\"}],\"name\":\"finalizeTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiverAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"interestRateModes\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiverAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"flashLoanSimple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ReserveConfigurationMap\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"name\":\"getEModeCategoryData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"ltv\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationThreshold\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationBonus\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"priceSource\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"internalType\":\"structDataTypes.EModeCategory\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"id\",\"type\":\"uint16\"}],\"name\":\"getReserveAddressById\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveData\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ReserveConfigurationMap\",\"name\":\"configuration\",\"type\":\"tuple\"},{\"internalType\":\"uint128\",\"name\":\"liquidityIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentLiquidityRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"variableBorrowIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentVariableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentStableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"id\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"variableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateStrategyAddress\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"accruedToTreasury\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"unbacked\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"isolationModeTotalDebt\",\"type\":\"uint128\"}],\"internalType\":\"structDataTypes.ReserveData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveNormalizedIncome\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveNormalizedVariableDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReservesList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserAccountData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalCollateralBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDebtBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableBorrowsBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentLiquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.UserConfigurationMap\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserEMode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stableDebtAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"variableDebtAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateStrategyAddress\",\"type\":\"address\"}],\"name\":\"initReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateralAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"debtAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"debtToCover\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"receiveAToken\",\"type\":\"bool\"}],\"name\":\"liquidationCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"name\":\"mintToTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"mintUnbacked\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"rebalanceStableBorrowRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRateMode\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"repay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRateMode\",\"type\":\"uint256\"}],\"name\":\"repayWithATokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRateMode\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"permitV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"permitR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"permitS\",\"type\":\"bytes32\"}],\"name\":\"repayWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"resetIsolationModeTotalDebt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ReserveConfigurationMap\",\"name\":\"configuration\",\"type\":\"tuple\"}],\"name\":\"setConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rateStrategyAddress\",\"type\":\"address\"}],\"name\":\"setReserveInterestRateStrategyAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"categoryId\",\"type\":\"uint8\"}],\"name\":\"setUserEMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"useAsCollateral\",\"type\":\"bool\"}],\"name\":\"setUserUseReserveAsCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"supply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"permitV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"permitR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"permitS\",\"type\":\"bytes32\"}],\"name\":\"supplyWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"interestRateMode\",\"type\":\"uint256\"}],\"name\":\"swapBorrowRateMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bridgeProtocolFee\",\"type\":\"uint256\"}],\"name\":\"updateBridgeProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"flashLoanPremiumTotal\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"flashLoanPremiumToProtocol\",\"type\":\"uint128\"}],\"name\":\"updateFlashloanPremiums\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AavePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use AavePoolMetaData.ABI instead.
var AavePoolABI = AavePoolMetaData.ABI

// AavePool is an auto generated Go binding around an Ethereum contract.
type AavePool struct {
	AavePoolCaller     // Read-only binding to the contract
	AavePoolTransactor // Write-only binding to the contract
	AavePoolFilterer   // Log filterer for contract events
}

// AavePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type AavePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AavePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AavePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AavePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AavePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AavePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AavePoolSession struct {
	Contract     *AavePool         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AavePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AavePoolCallerSession struct {
	Contract *AavePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AavePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AavePoolTransactorSession struct {
	Contract     *AavePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AavePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type AavePoolRaw struct {
	Contract *AavePool // Generic contract binding to access the raw methods on
}

// AavePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AavePoolCallerRaw struct {
	Contract *AavePoolCaller // Generic read-only contract binding to access the raw methods on
}

// AavePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AavePoolTransactorRaw struct {
	Contract *AavePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAavePool creates a new instance of AavePool, bound to a specific deployed contract.
func NewAavePool(address common.Address, backend bind.ContractBackend) (*AavePool, error) {
	contract, err := bindAavePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AavePool{AavePoolCaller: AavePoolCaller{contract: contract}, AavePoolTransactor: AavePoolTransactor{contract: contract}, AavePoolFilterer: AavePoolFilterer{contract: contract}}, nil
}

// NewAavePoolCaller creates a new read-only instance of AavePool, bound to a specific deployed contract.
func NewAavePoolCaller(address common.Address, caller bind.ContractCaller) (*AavePoolCaller, error) {
	contract, err := bindAavePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AavePoolCaller{contract: contract}, nil
}

// NewAavePoolTransactor creates a new write-only instance of AavePool, bound to a specific deployed contract.
func NewAavePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*AavePoolTransactor, error) {
	contract, err := bindAavePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AavePoolTransactor{contract: contract}, nil
}

// NewAavePoolFilterer creates a new log filterer instance of AavePool, bound to a specific deployed contract.
func NewAavePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*AavePoolFilterer, error) {
	contract, err := bindAavePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AavePoolFilterer{contract: contract}, nil
}

// bindAavePool binds a generic wrapper to an already deployed contract.
func bindAavePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AavePoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AavePool *AavePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AavePool.Contract.AavePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AavePool *AavePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AavePool.Contract.AavePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AavePool *AavePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AavePool.Contract.AavePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AavePool *AavePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AavePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AavePool *AavePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AavePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AavePool *AavePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AavePool.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_AavePool *AavePoolCaller) ADDRESSESPROVIDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "ADDRESSES_PROVIDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_AavePool *AavePoolSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _AavePool.Contract.ADDRESSESPROVIDER(&_AavePool.CallOpts)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_AavePool *AavePoolCallerSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _AavePool.Contract.ADDRESSESPROVIDER(&_AavePool.CallOpts)
}

// BRIDGEPROTOCOLFEE is a free data retrieval call binding the contract method 0x272d9072.
//
// Solidity: function BRIDGE_PROTOCOL_FEE() view returns(uint256)
func (_AavePool *AavePoolCaller) BRIDGEPROTOCOLFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "BRIDGE_PROTOCOL_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BRIDGEPROTOCOLFEE is a free data retrieval call binding the contract method 0x272d9072.
//
// Solidity: function BRIDGE_PROTOCOL_FEE() view returns(uint256)
func (_AavePool *AavePoolSession) BRIDGEPROTOCOLFEE() (*big.Int, error) {
	return _AavePool.Contract.BRIDGEPROTOCOLFEE(&_AavePool.CallOpts)
}

// BRIDGEPROTOCOLFEE is a free data retrieval call binding the contract method 0x272d9072.
//
// Solidity: function BRIDGE_PROTOCOL_FEE() view returns(uint256)
func (_AavePool *AavePoolCallerSession) BRIDGEPROTOCOLFEE() (*big.Int, error) {
	return _AavePool.Contract.BRIDGEPROTOCOLFEE(&_AavePool.CallOpts)
}

// FLASHLOANPREMIUMTOTAL is a free data retrieval call binding the contract method 0x074b2e43.
//
// Solidity: function FLASHLOAN_PREMIUM_TOTAL() view returns(uint128)
func (_AavePool *AavePoolCaller) FLASHLOANPREMIUMTOTAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "FLASHLOAN_PREMIUM_TOTAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLASHLOANPREMIUMTOTAL is a free data retrieval call binding the contract method 0x074b2e43.
//
// Solidity: function FLASHLOAN_PREMIUM_TOTAL() view returns(uint128)
func (_AavePool *AavePoolSession) FLASHLOANPREMIUMTOTAL() (*big.Int, error) {
	return _AavePool.Contract.FLASHLOANPREMIUMTOTAL(&_AavePool.CallOpts)
}

// FLASHLOANPREMIUMTOTAL is a free data retrieval call binding the contract method 0x074b2e43.
//
// Solidity: function FLASHLOAN_PREMIUM_TOTAL() view returns(uint128)
func (_AavePool *AavePoolCallerSession) FLASHLOANPREMIUMTOTAL() (*big.Int, error) {
	return _AavePool.Contract.FLASHLOANPREMIUMTOTAL(&_AavePool.CallOpts)
}

// FLASHLOANPREMIUMTOPROTOCOL is a free data retrieval call binding the contract method 0x6a99c036.
//
// Solidity: function FLASHLOAN_PREMIUM_TO_PROTOCOL() view returns(uint128)
func (_AavePool *AavePoolCaller) FLASHLOANPREMIUMTOPROTOCOL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "FLASHLOAN_PREMIUM_TO_PROTOCOL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLASHLOANPREMIUMTOPROTOCOL is a free data retrieval call binding the contract method 0x6a99c036.
//
// Solidity: function FLASHLOAN_PREMIUM_TO_PROTOCOL() view returns(uint128)
func (_AavePool *AavePoolSession) FLASHLOANPREMIUMTOPROTOCOL() (*big.Int, error) {
	return _AavePool.Contract.FLASHLOANPREMIUMTOPROTOCOL(&_AavePool.CallOpts)
}

// FLASHLOANPREMIUMTOPROTOCOL is a free data retrieval call binding the contract method 0x6a99c036.
//
// Solidity: function FLASHLOAN_PREMIUM_TO_PROTOCOL() view returns(uint128)
func (_AavePool *AavePoolCallerSession) FLASHLOANPREMIUMTOPROTOCOL() (*big.Int, error) {
	return _AavePool.Contract.FLASHLOANPREMIUMTOPROTOCOL(&_AavePool.CallOpts)
}

// MAXNUMBERRESERVES is a free data retrieval call binding the contract method 0xf8119d51.
//
// Solidity: function MAX_NUMBER_RESERVES() view returns(uint16)
func (_AavePool *AavePoolCaller) MAXNUMBERRESERVES(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "MAX_NUMBER_RESERVES")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXNUMBERRESERVES is a free data retrieval call binding the contract method 0xf8119d51.
//
// Solidity: function MAX_NUMBER_RESERVES() view returns(uint16)
func (_AavePool *AavePoolSession) MAXNUMBERRESERVES() (uint16, error) {
	return _AavePool.Contract.MAXNUMBERRESERVES(&_AavePool.CallOpts)
}

// MAXNUMBERRESERVES is a free data retrieval call binding the contract method 0xf8119d51.
//
// Solidity: function MAX_NUMBER_RESERVES() view returns(uint16)
func (_AavePool *AavePoolCallerSession) MAXNUMBERRESERVES() (uint16, error) {
	return _AavePool.Contract.MAXNUMBERRESERVES(&_AavePool.CallOpts)
}

// MAXSTABLERATEBORROWSIZEPERCENT is a free data retrieval call binding the contract method 0xe82fec2f.
//
// Solidity: function MAX_STABLE_RATE_BORROW_SIZE_PERCENT() view returns(uint256)
func (_AavePool *AavePoolCaller) MAXSTABLERATEBORROWSIZEPERCENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "MAX_STABLE_RATE_BORROW_SIZE_PERCENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSTABLERATEBORROWSIZEPERCENT is a free data retrieval call binding the contract method 0xe82fec2f.
//
// Solidity: function MAX_STABLE_RATE_BORROW_SIZE_PERCENT() view returns(uint256)
func (_AavePool *AavePoolSession) MAXSTABLERATEBORROWSIZEPERCENT() (*big.Int, error) {
	return _AavePool.Contract.MAXSTABLERATEBORROWSIZEPERCENT(&_AavePool.CallOpts)
}

// MAXSTABLERATEBORROWSIZEPERCENT is a free data retrieval call binding the contract method 0xe82fec2f.
//
// Solidity: function MAX_STABLE_RATE_BORROW_SIZE_PERCENT() view returns(uint256)
func (_AavePool *AavePoolCallerSession) MAXSTABLERATEBORROWSIZEPERCENT() (*big.Int, error) {
	return _AavePool.Contract.MAXSTABLERATEBORROWSIZEPERCENT(&_AavePool.CallOpts)
}

// GetConfiguration is a free data retrieval call binding the contract method 0xc44b11f7.
//
// Solidity: function getConfiguration(address asset) view returns((uint256))
func (_AavePool *AavePoolCaller) GetConfiguration(opts *bind.CallOpts, asset common.Address) (DataTypesReserveConfigurationMap, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getConfiguration", asset)

	if err != nil {
		return *new(DataTypesReserveConfigurationMap), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesReserveConfigurationMap)).(*DataTypesReserveConfigurationMap)

	return out0, err

}

// GetConfiguration is a free data retrieval call binding the contract method 0xc44b11f7.
//
// Solidity: function getConfiguration(address asset) view returns((uint256))
func (_AavePool *AavePoolSession) GetConfiguration(asset common.Address) (DataTypesReserveConfigurationMap, error) {
	return _AavePool.Contract.GetConfiguration(&_AavePool.CallOpts, asset)
}

// GetConfiguration is a free data retrieval call binding the contract method 0xc44b11f7.
//
// Solidity: function getConfiguration(address asset) view returns((uint256))
func (_AavePool *AavePoolCallerSession) GetConfiguration(asset common.Address) (DataTypesReserveConfigurationMap, error) {
	return _AavePool.Contract.GetConfiguration(&_AavePool.CallOpts, asset)
}

// GetEModeCategoryData is a free data retrieval call binding the contract method 0x6c6f6ae1.
//
// Solidity: function getEModeCategoryData(uint8 id) view returns((uint16,uint16,uint16,address,string))
func (_AavePool *AavePoolCaller) GetEModeCategoryData(opts *bind.CallOpts, id uint8) (DataTypesEModeCategory, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getEModeCategoryData", id)

	if err != nil {
		return *new(DataTypesEModeCategory), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesEModeCategory)).(*DataTypesEModeCategory)

	return out0, err

}

// GetEModeCategoryData is a free data retrieval call binding the contract method 0x6c6f6ae1.
//
// Solidity: function getEModeCategoryData(uint8 id) view returns((uint16,uint16,uint16,address,string))
func (_AavePool *AavePoolSession) GetEModeCategoryData(id uint8) (DataTypesEModeCategory, error) {
	return _AavePool.Contract.GetEModeCategoryData(&_AavePool.CallOpts, id)
}

// GetEModeCategoryData is a free data retrieval call binding the contract method 0x6c6f6ae1.
//
// Solidity: function getEModeCategoryData(uint8 id) view returns((uint16,uint16,uint16,address,string))
func (_AavePool *AavePoolCallerSession) GetEModeCategoryData(id uint8) (DataTypesEModeCategory, error) {
	return _AavePool.Contract.GetEModeCategoryData(&_AavePool.CallOpts, id)
}

// GetReserveAddressById is a free data retrieval call binding the contract method 0x52751797.
//
// Solidity: function getReserveAddressById(uint16 id) view returns(address)
func (_AavePool *AavePoolCaller) GetReserveAddressById(opts *bind.CallOpts, id uint16) (common.Address, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getReserveAddressById", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReserveAddressById is a free data retrieval call binding the contract method 0x52751797.
//
// Solidity: function getReserveAddressById(uint16 id) view returns(address)
func (_AavePool *AavePoolSession) GetReserveAddressById(id uint16) (common.Address, error) {
	return _AavePool.Contract.GetReserveAddressById(&_AavePool.CallOpts, id)
}

// GetReserveAddressById is a free data retrieval call binding the contract method 0x52751797.
//
// Solidity: function getReserveAddressById(uint16 id) view returns(address)
func (_AavePool *AavePoolCallerSession) GetReserveAddressById(id uint16) (common.Address, error) {
	return _AavePool.Contract.GetReserveAddressById(&_AavePool.CallOpts, id)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint128,uint40,uint16,address,address,address,address,uint128,uint128,uint128))
func (_AavePool *AavePoolCaller) GetReserveData(opts *bind.CallOpts, asset common.Address) (DataTypesReserveData, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getReserveData", asset)

	if err != nil {
		return *new(DataTypesReserveData), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesReserveData)).(*DataTypesReserveData)

	return out0, err

}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint128,uint40,uint16,address,address,address,address,uint128,uint128,uint128))
func (_AavePool *AavePoolSession) GetReserveData(asset common.Address) (DataTypesReserveData, error) {
	return _AavePool.Contract.GetReserveData(&_AavePool.CallOpts, asset)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint128,uint40,uint16,address,address,address,address,uint128,uint128,uint128))
func (_AavePool *AavePoolCallerSession) GetReserveData(asset common.Address) (DataTypesReserveData, error) {
	return _AavePool.Contract.GetReserveData(&_AavePool.CallOpts, asset)
}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address asset) view returns(uint256)
func (_AavePool *AavePoolCaller) GetReserveNormalizedIncome(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getReserveNormalizedIncome", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address asset) view returns(uint256)
func (_AavePool *AavePoolSession) GetReserveNormalizedIncome(asset common.Address) (*big.Int, error) {
	return _AavePool.Contract.GetReserveNormalizedIncome(&_AavePool.CallOpts, asset)
}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address asset) view returns(uint256)
func (_AavePool *AavePoolCallerSession) GetReserveNormalizedIncome(asset common.Address) (*big.Int, error) {
	return _AavePool.Contract.GetReserveNormalizedIncome(&_AavePool.CallOpts, asset)
}

// GetReserveNormalizedVariableDebt is a free data retrieval call binding the contract method 0x386497fd.
//
// Solidity: function getReserveNormalizedVariableDebt(address asset) view returns(uint256)
func (_AavePool *AavePoolCaller) GetReserveNormalizedVariableDebt(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getReserveNormalizedVariableDebt", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveNormalizedVariableDebt is a free data retrieval call binding the contract method 0x386497fd.
//
// Solidity: function getReserveNormalizedVariableDebt(address asset) view returns(uint256)
func (_AavePool *AavePoolSession) GetReserveNormalizedVariableDebt(asset common.Address) (*big.Int, error) {
	return _AavePool.Contract.GetReserveNormalizedVariableDebt(&_AavePool.CallOpts, asset)
}

// GetReserveNormalizedVariableDebt is a free data retrieval call binding the contract method 0x386497fd.
//
// Solidity: function getReserveNormalizedVariableDebt(address asset) view returns(uint256)
func (_AavePool *AavePoolCallerSession) GetReserveNormalizedVariableDebt(asset common.Address) (*big.Int, error) {
	return _AavePool.Contract.GetReserveNormalizedVariableDebt(&_AavePool.CallOpts, asset)
}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_AavePool *AavePoolCaller) GetReservesList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getReservesList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_AavePool *AavePoolSession) GetReservesList() ([]common.Address, error) {
	return _AavePool.Contract.GetReservesList(&_AavePool.CallOpts)
}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_AavePool *AavePoolCallerSession) GetReservesList() ([]common.Address, error) {
	return _AavePool.Contract.GetReservesList(&_AavePool.CallOpts)
}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address user) view returns(uint256 totalCollateralBase, uint256 totalDebtBase, uint256 availableBorrowsBase, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_AavePool *AavePoolCaller) GetUserAccountData(opts *bind.CallOpts, user common.Address) (struct {
	TotalCollateralBase         *big.Int
	TotalDebtBase               *big.Int
	AvailableBorrowsBase        *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getUserAccountData", user)

	outstruct := new(struct {
		TotalCollateralBase         *big.Int
		TotalDebtBase               *big.Int
		AvailableBorrowsBase        *big.Int
		CurrentLiquidationThreshold *big.Int
		Ltv                         *big.Int
		HealthFactor                *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalCollateralBase = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalDebtBase = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AvailableBorrowsBase = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CurrentLiquidationThreshold = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Ltv = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.HealthFactor = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address user) view returns(uint256 totalCollateralBase, uint256 totalDebtBase, uint256 availableBorrowsBase, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_AavePool *AavePoolSession) GetUserAccountData(user common.Address) (struct {
	TotalCollateralBase         *big.Int
	TotalDebtBase               *big.Int
	AvailableBorrowsBase        *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	return _AavePool.Contract.GetUserAccountData(&_AavePool.CallOpts, user)
}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address user) view returns(uint256 totalCollateralBase, uint256 totalDebtBase, uint256 availableBorrowsBase, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_AavePool *AavePoolCallerSession) GetUserAccountData(user common.Address) (struct {
	TotalCollateralBase         *big.Int
	TotalDebtBase               *big.Int
	AvailableBorrowsBase        *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	return _AavePool.Contract.GetUserAccountData(&_AavePool.CallOpts, user)
}

// GetUserConfiguration is a free data retrieval call binding the contract method 0x4417a583.
//
// Solidity: function getUserConfiguration(address user) view returns((uint256))
func (_AavePool *AavePoolCaller) GetUserConfiguration(opts *bind.CallOpts, user common.Address) (DataTypesUserConfigurationMap, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getUserConfiguration", user)

	if err != nil {
		return *new(DataTypesUserConfigurationMap), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesUserConfigurationMap)).(*DataTypesUserConfigurationMap)

	return out0, err

}

// GetUserConfiguration is a free data retrieval call binding the contract method 0x4417a583.
//
// Solidity: function getUserConfiguration(address user) view returns((uint256))
func (_AavePool *AavePoolSession) GetUserConfiguration(user common.Address) (DataTypesUserConfigurationMap, error) {
	return _AavePool.Contract.GetUserConfiguration(&_AavePool.CallOpts, user)
}

// GetUserConfiguration is a free data retrieval call binding the contract method 0x4417a583.
//
// Solidity: function getUserConfiguration(address user) view returns((uint256))
func (_AavePool *AavePoolCallerSession) GetUserConfiguration(user common.Address) (DataTypesUserConfigurationMap, error) {
	return _AavePool.Contract.GetUserConfiguration(&_AavePool.CallOpts, user)
}

// GetUserEMode is a free data retrieval call binding the contract method 0xeddf1b79.
//
// Solidity: function getUserEMode(address user) view returns(uint256)
func (_AavePool *AavePoolCaller) GetUserEMode(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AavePool.contract.Call(opts, &out, "getUserEMode", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserEMode is a free data retrieval call binding the contract method 0xeddf1b79.
//
// Solidity: function getUserEMode(address user) view returns(uint256)
func (_AavePool *AavePoolSession) GetUserEMode(user common.Address) (*big.Int, error) {
	return _AavePool.Contract.GetUserEMode(&_AavePool.CallOpts, user)
}

// GetUserEMode is a free data retrieval call binding the contract method 0xeddf1b79.
//
// Solidity: function getUserEMode(address user) view returns(uint256)
func (_AavePool *AavePoolCallerSession) GetUserEMode(user common.Address) (*big.Int, error) {
	return _AavePool.Contract.GetUserEMode(&_AavePool.CallOpts, user)
}

// BackUnbacked is a paid mutator transaction binding the contract method 0xd65dc7a1.
//
// Solidity: function backUnbacked(address asset, uint256 amount, uint256 fee) returns()
func (_AavePool *AavePoolTransactor) BackUnbacked(opts *bind.TransactOpts, asset common.Address, amount *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "backUnbacked", asset, amount, fee)
}

// BackUnbacked is a paid mutator transaction binding the contract method 0xd65dc7a1.
//
// Solidity: function backUnbacked(address asset, uint256 amount, uint256 fee) returns()
func (_AavePool *AavePoolSession) BackUnbacked(asset common.Address, amount *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.BackUnbacked(&_AavePool.TransactOpts, asset, amount, fee)
}

// BackUnbacked is a paid mutator transaction binding the contract method 0xd65dc7a1.
//
// Solidity: function backUnbacked(address asset, uint256 amount, uint256 fee) returns()
func (_AavePool *AavePoolTransactorSession) BackUnbacked(asset common.Address, amount *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.BackUnbacked(&_AavePool.TransactOpts, asset, amount, fee)
}

// Borrow is a paid mutator transaction binding the contract method 0xa415bcad.
//
// Solidity: function borrow(address asset, uint256 amount, uint256 interestRateMode, uint16 referralCode, address onBehalfOf) returns()
func (_AavePool *AavePoolTransactor) Borrow(opts *bind.TransactOpts, asset common.Address, amount *big.Int, interestRateMode *big.Int, referralCode uint16, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "borrow", asset, amount, interestRateMode, referralCode, onBehalfOf)
}

// Borrow is a paid mutator transaction binding the contract method 0xa415bcad.
//
// Solidity: function borrow(address asset, uint256 amount, uint256 interestRateMode, uint16 referralCode, address onBehalfOf) returns()
func (_AavePool *AavePoolSession) Borrow(asset common.Address, amount *big.Int, interestRateMode *big.Int, referralCode uint16, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.Borrow(&_AavePool.TransactOpts, asset, amount, interestRateMode, referralCode, onBehalfOf)
}

// Borrow is a paid mutator transaction binding the contract method 0xa415bcad.
//
// Solidity: function borrow(address asset, uint256 amount, uint256 interestRateMode, uint16 referralCode, address onBehalfOf) returns()
func (_AavePool *AavePoolTransactorSession) Borrow(asset common.Address, amount *big.Int, interestRateMode *big.Int, referralCode uint16, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.Borrow(&_AavePool.TransactOpts, asset, amount, interestRateMode, referralCode, onBehalfOf)
}

// ConfigureEModeCategory is a paid mutator transaction binding the contract method 0xd579ea7d.
//
// Solidity: function configureEModeCategory(uint8 id, (uint16,uint16,uint16,address,string) config) returns()
func (_AavePool *AavePoolTransactor) ConfigureEModeCategory(opts *bind.TransactOpts, id uint8, config DataTypesEModeCategory) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "configureEModeCategory", id, config)
}

// ConfigureEModeCategory is a paid mutator transaction binding the contract method 0xd579ea7d.
//
// Solidity: function configureEModeCategory(uint8 id, (uint16,uint16,uint16,address,string) config) returns()
func (_AavePool *AavePoolSession) ConfigureEModeCategory(id uint8, config DataTypesEModeCategory) (*types.Transaction, error) {
	return _AavePool.Contract.ConfigureEModeCategory(&_AavePool.TransactOpts, id, config)
}

// ConfigureEModeCategory is a paid mutator transaction binding the contract method 0xd579ea7d.
//
// Solidity: function configureEModeCategory(uint8 id, (uint16,uint16,uint16,address,string) config) returns()
func (_AavePool *AavePoolTransactorSession) ConfigureEModeCategory(id uint8, config DataTypesEModeCategory) (*types.Transaction, error) {
	return _AavePool.Contract.ConfigureEModeCategory(&_AavePool.TransactOpts, id, config)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_AavePool *AavePoolTransactor) Deposit(opts *bind.TransactOpts, asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "deposit", asset, amount, onBehalfOf, referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_AavePool *AavePoolSession) Deposit(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AavePool.Contract.Deposit(&_AavePool.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_AavePool *AavePoolTransactorSession) Deposit(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AavePool.Contract.Deposit(&_AavePool.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// DropReserve is a paid mutator transaction binding the contract method 0x63c9b860.
//
// Solidity: function dropReserve(address asset) returns()
func (_AavePool *AavePoolTransactor) DropReserve(opts *bind.TransactOpts, asset common.Address) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "dropReserve", asset)
}

// DropReserve is a paid mutator transaction binding the contract method 0x63c9b860.
//
// Solidity: function dropReserve(address asset) returns()
func (_AavePool *AavePoolSession) DropReserve(asset common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.DropReserve(&_AavePool.TransactOpts, asset)
}

// DropReserve is a paid mutator transaction binding the contract method 0x63c9b860.
//
// Solidity: function dropReserve(address asset) returns()
func (_AavePool *AavePoolTransactorSession) DropReserve(asset common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.DropReserve(&_AavePool.TransactOpts, asset)
}

// FinalizeTransfer is a paid mutator transaction binding the contract method 0xd5ed3933.
//
// Solidity: function finalizeTransfer(address asset, address from, address to, uint256 amount, uint256 balanceFromBefore, uint256 balanceToBefore) returns()
func (_AavePool *AavePoolTransactor) FinalizeTransfer(opts *bind.TransactOpts, asset common.Address, from common.Address, to common.Address, amount *big.Int, balanceFromBefore *big.Int, balanceToBefore *big.Int) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "finalizeTransfer", asset, from, to, amount, balanceFromBefore, balanceToBefore)
}

// FinalizeTransfer is a paid mutator transaction binding the contract method 0xd5ed3933.
//
// Solidity: function finalizeTransfer(address asset, address from, address to, uint256 amount, uint256 balanceFromBefore, uint256 balanceToBefore) returns()
func (_AavePool *AavePoolSession) FinalizeTransfer(asset common.Address, from common.Address, to common.Address, amount *big.Int, balanceFromBefore *big.Int, balanceToBefore *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.FinalizeTransfer(&_AavePool.TransactOpts, asset, from, to, amount, balanceFromBefore, balanceToBefore)
}

// FinalizeTransfer is a paid mutator transaction binding the contract method 0xd5ed3933.
//
// Solidity: function finalizeTransfer(address asset, address from, address to, uint256 amount, uint256 balanceFromBefore, uint256 balanceToBefore) returns()
func (_AavePool *AavePoolTransactorSession) FinalizeTransfer(asset common.Address, from common.Address, to common.Address, amount *big.Int, balanceFromBefore *big.Int, balanceToBefore *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.FinalizeTransfer(&_AavePool.TransactOpts, asset, from, to, amount, balanceFromBefore, balanceToBefore)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xab9c4b5d.
//
// Solidity: function flashLoan(address receiverAddress, address[] assets, uint256[] amounts, uint256[] interestRateModes, address onBehalfOf, bytes params, uint16 referralCode) returns()
func (_AavePool *AavePoolTransactor) FlashLoan(opts *bind.TransactOpts, receiverAddress common.Address, assets []common.Address, amounts []*big.Int, interestRateModes []*big.Int, onBehalfOf common.Address, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "flashLoan", receiverAddress, assets, amounts, interestRateModes, onBehalfOf, params, referralCode)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xab9c4b5d.
//
// Solidity: function flashLoan(address receiverAddress, address[] assets, uint256[] amounts, uint256[] interestRateModes, address onBehalfOf, bytes params, uint16 referralCode) returns()
func (_AavePool *AavePoolSession) FlashLoan(receiverAddress common.Address, assets []common.Address, amounts []*big.Int, interestRateModes []*big.Int, onBehalfOf common.Address, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _AavePool.Contract.FlashLoan(&_AavePool.TransactOpts, receiverAddress, assets, amounts, interestRateModes, onBehalfOf, params, referralCode)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xab9c4b5d.
//
// Solidity: function flashLoan(address receiverAddress, address[] assets, uint256[] amounts, uint256[] interestRateModes, address onBehalfOf, bytes params, uint16 referralCode) returns()
func (_AavePool *AavePoolTransactorSession) FlashLoan(receiverAddress common.Address, assets []common.Address, amounts []*big.Int, interestRateModes []*big.Int, onBehalfOf common.Address, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _AavePool.Contract.FlashLoan(&_AavePool.TransactOpts, receiverAddress, assets, amounts, interestRateModes, onBehalfOf, params, referralCode)
}

// FlashLoanSimple is a paid mutator transaction binding the contract method 0x42b0b77c.
//
// Solidity: function flashLoanSimple(address receiverAddress, address asset, uint256 amount, bytes params, uint16 referralCode) returns()
func (_AavePool *AavePoolTransactor) FlashLoanSimple(opts *bind.TransactOpts, receiverAddress common.Address, asset common.Address, amount *big.Int, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "flashLoanSimple", receiverAddress, asset, amount, params, referralCode)
}

// FlashLoanSimple is a paid mutator transaction binding the contract method 0x42b0b77c.
//
// Solidity: function flashLoanSimple(address receiverAddress, address asset, uint256 amount, bytes params, uint16 referralCode) returns()
func (_AavePool *AavePoolSession) FlashLoanSimple(receiverAddress common.Address, asset common.Address, amount *big.Int, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _AavePool.Contract.FlashLoanSimple(&_AavePool.TransactOpts, receiverAddress, asset, amount, params, referralCode)
}

// FlashLoanSimple is a paid mutator transaction binding the contract method 0x42b0b77c.
//
// Solidity: function flashLoanSimple(address receiverAddress, address asset, uint256 amount, bytes params, uint16 referralCode) returns()
func (_AavePool *AavePoolTransactorSession) FlashLoanSimple(receiverAddress common.Address, asset common.Address, amount *big.Int, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _AavePool.Contract.FlashLoanSimple(&_AavePool.TransactOpts, receiverAddress, asset, amount, params, referralCode)
}

// InitReserve is a paid mutator transaction binding the contract method 0x7a708e92.
//
// Solidity: function initReserve(address asset, address aTokenAddress, address stableDebtAddress, address variableDebtAddress, address interestRateStrategyAddress) returns()
func (_AavePool *AavePoolTransactor) InitReserve(opts *bind.TransactOpts, asset common.Address, aTokenAddress common.Address, stableDebtAddress common.Address, variableDebtAddress common.Address, interestRateStrategyAddress common.Address) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "initReserve", asset, aTokenAddress, stableDebtAddress, variableDebtAddress, interestRateStrategyAddress)
}

// InitReserve is a paid mutator transaction binding the contract method 0x7a708e92.
//
// Solidity: function initReserve(address asset, address aTokenAddress, address stableDebtAddress, address variableDebtAddress, address interestRateStrategyAddress) returns()
func (_AavePool *AavePoolSession) InitReserve(asset common.Address, aTokenAddress common.Address, stableDebtAddress common.Address, variableDebtAddress common.Address, interestRateStrategyAddress common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.InitReserve(&_AavePool.TransactOpts, asset, aTokenAddress, stableDebtAddress, variableDebtAddress, interestRateStrategyAddress)
}

// InitReserve is a paid mutator transaction binding the contract method 0x7a708e92.
//
// Solidity: function initReserve(address asset, address aTokenAddress, address stableDebtAddress, address variableDebtAddress, address interestRateStrategyAddress) returns()
func (_AavePool *AavePoolTransactorSession) InitReserve(asset common.Address, aTokenAddress common.Address, stableDebtAddress common.Address, variableDebtAddress common.Address, interestRateStrategyAddress common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.InitReserve(&_AavePool.TransactOpts, asset, aTokenAddress, stableDebtAddress, variableDebtAddress, interestRateStrategyAddress)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address collateralAsset, address debtAsset, address user, uint256 debtToCover, bool receiveAToken) returns()
func (_AavePool *AavePoolTransactor) LiquidationCall(opts *bind.TransactOpts, collateralAsset common.Address, debtAsset common.Address, user common.Address, debtToCover *big.Int, receiveAToken bool) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "liquidationCall", collateralAsset, debtAsset, user, debtToCover, receiveAToken)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address collateralAsset, address debtAsset, address user, uint256 debtToCover, bool receiveAToken) returns()
func (_AavePool *AavePoolSession) LiquidationCall(collateralAsset common.Address, debtAsset common.Address, user common.Address, debtToCover *big.Int, receiveAToken bool) (*types.Transaction, error) {
	return _AavePool.Contract.LiquidationCall(&_AavePool.TransactOpts, collateralAsset, debtAsset, user, debtToCover, receiveAToken)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address collateralAsset, address debtAsset, address user, uint256 debtToCover, bool receiveAToken) returns()
func (_AavePool *AavePoolTransactorSession) LiquidationCall(collateralAsset common.Address, debtAsset common.Address, user common.Address, debtToCover *big.Int, receiveAToken bool) (*types.Transaction, error) {
	return _AavePool.Contract.LiquidationCall(&_AavePool.TransactOpts, collateralAsset, debtAsset, user, debtToCover, receiveAToken)
}

// MintToTreasury is a paid mutator transaction binding the contract method 0x9cd19996.
//
// Solidity: function mintToTreasury(address[] assets) returns()
func (_AavePool *AavePoolTransactor) MintToTreasury(opts *bind.TransactOpts, assets []common.Address) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "mintToTreasury", assets)
}

// MintToTreasury is a paid mutator transaction binding the contract method 0x9cd19996.
//
// Solidity: function mintToTreasury(address[] assets) returns()
func (_AavePool *AavePoolSession) MintToTreasury(assets []common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.MintToTreasury(&_AavePool.TransactOpts, assets)
}

// MintToTreasury is a paid mutator transaction binding the contract method 0x9cd19996.
//
// Solidity: function mintToTreasury(address[] assets) returns()
func (_AavePool *AavePoolTransactorSession) MintToTreasury(assets []common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.MintToTreasury(&_AavePool.TransactOpts, assets)
}

// MintUnbacked is a paid mutator transaction binding the contract method 0x69a933a5.
//
// Solidity: function mintUnbacked(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_AavePool *AavePoolTransactor) MintUnbacked(opts *bind.TransactOpts, asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "mintUnbacked", asset, amount, onBehalfOf, referralCode)
}

// MintUnbacked is a paid mutator transaction binding the contract method 0x69a933a5.
//
// Solidity: function mintUnbacked(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_AavePool *AavePoolSession) MintUnbacked(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AavePool.Contract.MintUnbacked(&_AavePool.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// MintUnbacked is a paid mutator transaction binding the contract method 0x69a933a5.
//
// Solidity: function mintUnbacked(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_AavePool *AavePoolTransactorSession) MintUnbacked(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AavePool.Contract.MintUnbacked(&_AavePool.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// RebalanceStableBorrowRate is a paid mutator transaction binding the contract method 0xcd112382.
//
// Solidity: function rebalanceStableBorrowRate(address asset, address user) returns()
func (_AavePool *AavePoolTransactor) RebalanceStableBorrowRate(opts *bind.TransactOpts, asset common.Address, user common.Address) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "rebalanceStableBorrowRate", asset, user)
}

// RebalanceStableBorrowRate is a paid mutator transaction binding the contract method 0xcd112382.
//
// Solidity: function rebalanceStableBorrowRate(address asset, address user) returns()
func (_AavePool *AavePoolSession) RebalanceStableBorrowRate(asset common.Address, user common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.RebalanceStableBorrowRate(&_AavePool.TransactOpts, asset, user)
}

// RebalanceStableBorrowRate is a paid mutator transaction binding the contract method 0xcd112382.
//
// Solidity: function rebalanceStableBorrowRate(address asset, address user) returns()
func (_AavePool *AavePoolTransactorSession) RebalanceStableBorrowRate(asset common.Address, user common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.RebalanceStableBorrowRate(&_AavePool.TransactOpts, asset, user)
}

// Repay is a paid mutator transaction binding the contract method 0x573ade81.
//
// Solidity: function repay(address asset, uint256 amount, uint256 interestRateMode, address onBehalfOf) returns(uint256)
func (_AavePool *AavePoolTransactor) Repay(opts *bind.TransactOpts, asset common.Address, amount *big.Int, interestRateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "repay", asset, amount, interestRateMode, onBehalfOf)
}

// Repay is a paid mutator transaction binding the contract method 0x573ade81.
//
// Solidity: function repay(address asset, uint256 amount, uint256 interestRateMode, address onBehalfOf) returns(uint256)
func (_AavePool *AavePoolSession) Repay(asset common.Address, amount *big.Int, interestRateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.Repay(&_AavePool.TransactOpts, asset, amount, interestRateMode, onBehalfOf)
}

// Repay is a paid mutator transaction binding the contract method 0x573ade81.
//
// Solidity: function repay(address asset, uint256 amount, uint256 interestRateMode, address onBehalfOf) returns(uint256)
func (_AavePool *AavePoolTransactorSession) Repay(asset common.Address, amount *big.Int, interestRateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.Repay(&_AavePool.TransactOpts, asset, amount, interestRateMode, onBehalfOf)
}

// RepayWithATokens is a paid mutator transaction binding the contract method 0x2dad97d4.
//
// Solidity: function repayWithATokens(address asset, uint256 amount, uint256 interestRateMode) returns(uint256)
func (_AavePool *AavePoolTransactor) RepayWithATokens(opts *bind.TransactOpts, asset common.Address, amount *big.Int, interestRateMode *big.Int) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "repayWithATokens", asset, amount, interestRateMode)
}

// RepayWithATokens is a paid mutator transaction binding the contract method 0x2dad97d4.
//
// Solidity: function repayWithATokens(address asset, uint256 amount, uint256 interestRateMode) returns(uint256)
func (_AavePool *AavePoolSession) RepayWithATokens(asset common.Address, amount *big.Int, interestRateMode *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.RepayWithATokens(&_AavePool.TransactOpts, asset, amount, interestRateMode)
}

// RepayWithATokens is a paid mutator transaction binding the contract method 0x2dad97d4.
//
// Solidity: function repayWithATokens(address asset, uint256 amount, uint256 interestRateMode) returns(uint256)
func (_AavePool *AavePoolTransactorSession) RepayWithATokens(asset common.Address, amount *big.Int, interestRateMode *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.RepayWithATokens(&_AavePool.TransactOpts, asset, amount, interestRateMode)
}

// RepayWithPermit is a paid mutator transaction binding the contract method 0xee3e210b.
//
// Solidity: function repayWithPermit(address asset, uint256 amount, uint256 interestRateMode, address onBehalfOf, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns(uint256)
func (_AavePool *AavePoolTransactor) RepayWithPermit(opts *bind.TransactOpts, asset common.Address, amount *big.Int, interestRateMode *big.Int, onBehalfOf common.Address, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "repayWithPermit", asset, amount, interestRateMode, onBehalfOf, deadline, permitV, permitR, permitS)
}

// RepayWithPermit is a paid mutator transaction binding the contract method 0xee3e210b.
//
// Solidity: function repayWithPermit(address asset, uint256 amount, uint256 interestRateMode, address onBehalfOf, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns(uint256)
func (_AavePool *AavePoolSession) RepayWithPermit(asset common.Address, amount *big.Int, interestRateMode *big.Int, onBehalfOf common.Address, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _AavePool.Contract.RepayWithPermit(&_AavePool.TransactOpts, asset, amount, interestRateMode, onBehalfOf, deadline, permitV, permitR, permitS)
}

// RepayWithPermit is a paid mutator transaction binding the contract method 0xee3e210b.
//
// Solidity: function repayWithPermit(address asset, uint256 amount, uint256 interestRateMode, address onBehalfOf, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns(uint256)
func (_AavePool *AavePoolTransactorSession) RepayWithPermit(asset common.Address, amount *big.Int, interestRateMode *big.Int, onBehalfOf common.Address, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _AavePool.Contract.RepayWithPermit(&_AavePool.TransactOpts, asset, amount, interestRateMode, onBehalfOf, deadline, permitV, permitR, permitS)
}

// RescueTokens is a paid mutator transaction binding the contract method 0xcea9d26f.
//
// Solidity: function rescueTokens(address token, address to, uint256 amount) returns()
func (_AavePool *AavePoolTransactor) RescueTokens(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "rescueTokens", token, to, amount)
}

// RescueTokens is a paid mutator transaction binding the contract method 0xcea9d26f.
//
// Solidity: function rescueTokens(address token, address to, uint256 amount) returns()
func (_AavePool *AavePoolSession) RescueTokens(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.RescueTokens(&_AavePool.TransactOpts, token, to, amount)
}

// RescueTokens is a paid mutator transaction binding the contract method 0xcea9d26f.
//
// Solidity: function rescueTokens(address token, address to, uint256 amount) returns()
func (_AavePool *AavePoolTransactorSession) RescueTokens(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.RescueTokens(&_AavePool.TransactOpts, token, to, amount)
}

// ResetIsolationModeTotalDebt is a paid mutator transaction binding the contract method 0xe43e88a1.
//
// Solidity: function resetIsolationModeTotalDebt(address asset) returns()
func (_AavePool *AavePoolTransactor) ResetIsolationModeTotalDebt(opts *bind.TransactOpts, asset common.Address) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "resetIsolationModeTotalDebt", asset)
}

// ResetIsolationModeTotalDebt is a paid mutator transaction binding the contract method 0xe43e88a1.
//
// Solidity: function resetIsolationModeTotalDebt(address asset) returns()
func (_AavePool *AavePoolSession) ResetIsolationModeTotalDebt(asset common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.ResetIsolationModeTotalDebt(&_AavePool.TransactOpts, asset)
}

// ResetIsolationModeTotalDebt is a paid mutator transaction binding the contract method 0xe43e88a1.
//
// Solidity: function resetIsolationModeTotalDebt(address asset) returns()
func (_AavePool *AavePoolTransactorSession) ResetIsolationModeTotalDebt(asset common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.ResetIsolationModeTotalDebt(&_AavePool.TransactOpts, asset)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0xf51e435b.
//
// Solidity: function setConfiguration(address asset, (uint256) configuration) returns()
func (_AavePool *AavePoolTransactor) SetConfiguration(opts *bind.TransactOpts, asset common.Address, configuration DataTypesReserveConfigurationMap) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "setConfiguration", asset, configuration)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0xf51e435b.
//
// Solidity: function setConfiguration(address asset, (uint256) configuration) returns()
func (_AavePool *AavePoolSession) SetConfiguration(asset common.Address, configuration DataTypesReserveConfigurationMap) (*types.Transaction, error) {
	return _AavePool.Contract.SetConfiguration(&_AavePool.TransactOpts, asset, configuration)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0xf51e435b.
//
// Solidity: function setConfiguration(address asset, (uint256) configuration) returns()
func (_AavePool *AavePoolTransactorSession) SetConfiguration(asset common.Address, configuration DataTypesReserveConfigurationMap) (*types.Transaction, error) {
	return _AavePool.Contract.SetConfiguration(&_AavePool.TransactOpts, asset, configuration)
}

// SetReserveInterestRateStrategyAddress is a paid mutator transaction binding the contract method 0x1d2118f9.
//
// Solidity: function setReserveInterestRateStrategyAddress(address asset, address rateStrategyAddress) returns()
func (_AavePool *AavePoolTransactor) SetReserveInterestRateStrategyAddress(opts *bind.TransactOpts, asset common.Address, rateStrategyAddress common.Address) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "setReserveInterestRateStrategyAddress", asset, rateStrategyAddress)
}

// SetReserveInterestRateStrategyAddress is a paid mutator transaction binding the contract method 0x1d2118f9.
//
// Solidity: function setReserveInterestRateStrategyAddress(address asset, address rateStrategyAddress) returns()
func (_AavePool *AavePoolSession) SetReserveInterestRateStrategyAddress(asset common.Address, rateStrategyAddress common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.SetReserveInterestRateStrategyAddress(&_AavePool.TransactOpts, asset, rateStrategyAddress)
}

// SetReserveInterestRateStrategyAddress is a paid mutator transaction binding the contract method 0x1d2118f9.
//
// Solidity: function setReserveInterestRateStrategyAddress(address asset, address rateStrategyAddress) returns()
func (_AavePool *AavePoolTransactorSession) SetReserveInterestRateStrategyAddress(asset common.Address, rateStrategyAddress common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.SetReserveInterestRateStrategyAddress(&_AavePool.TransactOpts, asset, rateStrategyAddress)
}

// SetUserEMode is a paid mutator transaction binding the contract method 0x28530a47.
//
// Solidity: function setUserEMode(uint8 categoryId) returns()
func (_AavePool *AavePoolTransactor) SetUserEMode(opts *bind.TransactOpts, categoryId uint8) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "setUserEMode", categoryId)
}

// SetUserEMode is a paid mutator transaction binding the contract method 0x28530a47.
//
// Solidity: function setUserEMode(uint8 categoryId) returns()
func (_AavePool *AavePoolSession) SetUserEMode(categoryId uint8) (*types.Transaction, error) {
	return _AavePool.Contract.SetUserEMode(&_AavePool.TransactOpts, categoryId)
}

// SetUserEMode is a paid mutator transaction binding the contract method 0x28530a47.
//
// Solidity: function setUserEMode(uint8 categoryId) returns()
func (_AavePool *AavePoolTransactorSession) SetUserEMode(categoryId uint8) (*types.Transaction, error) {
	return _AavePool.Contract.SetUserEMode(&_AavePool.TransactOpts, categoryId)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address asset, bool useAsCollateral) returns()
func (_AavePool *AavePoolTransactor) SetUserUseReserveAsCollateral(opts *bind.TransactOpts, asset common.Address, useAsCollateral bool) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "setUserUseReserveAsCollateral", asset, useAsCollateral)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address asset, bool useAsCollateral) returns()
func (_AavePool *AavePoolSession) SetUserUseReserveAsCollateral(asset common.Address, useAsCollateral bool) (*types.Transaction, error) {
	return _AavePool.Contract.SetUserUseReserveAsCollateral(&_AavePool.TransactOpts, asset, useAsCollateral)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address asset, bool useAsCollateral) returns()
func (_AavePool *AavePoolTransactorSession) SetUserUseReserveAsCollateral(asset common.Address, useAsCollateral bool) (*types.Transaction, error) {
	return _AavePool.Contract.SetUserUseReserveAsCollateral(&_AavePool.TransactOpts, asset, useAsCollateral)
}

// Supply is a paid mutator transaction binding the contract method 0x617ba037.
//
// Solidity: function supply(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_AavePool *AavePoolTransactor) Supply(opts *bind.TransactOpts, asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "supply", asset, amount, onBehalfOf, referralCode)
}

// Supply is a paid mutator transaction binding the contract method 0x617ba037.
//
// Solidity: function supply(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_AavePool *AavePoolSession) Supply(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AavePool.Contract.Supply(&_AavePool.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// Supply is a paid mutator transaction binding the contract method 0x617ba037.
//
// Solidity: function supply(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_AavePool *AavePoolTransactorSession) Supply(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AavePool.Contract.Supply(&_AavePool.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// SupplyWithPermit is a paid mutator transaction binding the contract method 0x02c205f0.
//
// Solidity: function supplyWithPermit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns()
func (_AavePool *AavePoolTransactor) SupplyWithPermit(opts *bind.TransactOpts, asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "supplyWithPermit", asset, amount, onBehalfOf, referralCode, deadline, permitV, permitR, permitS)
}

// SupplyWithPermit is a paid mutator transaction binding the contract method 0x02c205f0.
//
// Solidity: function supplyWithPermit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns()
func (_AavePool *AavePoolSession) SupplyWithPermit(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _AavePool.Contract.SupplyWithPermit(&_AavePool.TransactOpts, asset, amount, onBehalfOf, referralCode, deadline, permitV, permitR, permitS)
}

// SupplyWithPermit is a paid mutator transaction binding the contract method 0x02c205f0.
//
// Solidity: function supplyWithPermit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns()
func (_AavePool *AavePoolTransactorSession) SupplyWithPermit(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _AavePool.Contract.SupplyWithPermit(&_AavePool.TransactOpts, asset, amount, onBehalfOf, referralCode, deadline, permitV, permitR, permitS)
}

// SwapBorrowRateMode is a paid mutator transaction binding the contract method 0x94ba89a2.
//
// Solidity: function swapBorrowRateMode(address asset, uint256 interestRateMode) returns()
func (_AavePool *AavePoolTransactor) SwapBorrowRateMode(opts *bind.TransactOpts, asset common.Address, interestRateMode *big.Int) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "swapBorrowRateMode", asset, interestRateMode)
}

// SwapBorrowRateMode is a paid mutator transaction binding the contract method 0x94ba89a2.
//
// Solidity: function swapBorrowRateMode(address asset, uint256 interestRateMode) returns()
func (_AavePool *AavePoolSession) SwapBorrowRateMode(asset common.Address, interestRateMode *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.SwapBorrowRateMode(&_AavePool.TransactOpts, asset, interestRateMode)
}

// SwapBorrowRateMode is a paid mutator transaction binding the contract method 0x94ba89a2.
//
// Solidity: function swapBorrowRateMode(address asset, uint256 interestRateMode) returns()
func (_AavePool *AavePoolTransactorSession) SwapBorrowRateMode(asset common.Address, interestRateMode *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.SwapBorrowRateMode(&_AavePool.TransactOpts, asset, interestRateMode)
}

// UpdateBridgeProtocolFee is a paid mutator transaction binding the contract method 0x3036b439.
//
// Solidity: function updateBridgeProtocolFee(uint256 bridgeProtocolFee) returns()
func (_AavePool *AavePoolTransactor) UpdateBridgeProtocolFee(opts *bind.TransactOpts, bridgeProtocolFee *big.Int) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "updateBridgeProtocolFee", bridgeProtocolFee)
}

// UpdateBridgeProtocolFee is a paid mutator transaction binding the contract method 0x3036b439.
//
// Solidity: function updateBridgeProtocolFee(uint256 bridgeProtocolFee) returns()
func (_AavePool *AavePoolSession) UpdateBridgeProtocolFee(bridgeProtocolFee *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.UpdateBridgeProtocolFee(&_AavePool.TransactOpts, bridgeProtocolFee)
}

// UpdateBridgeProtocolFee is a paid mutator transaction binding the contract method 0x3036b439.
//
// Solidity: function updateBridgeProtocolFee(uint256 bridgeProtocolFee) returns()
func (_AavePool *AavePoolTransactorSession) UpdateBridgeProtocolFee(bridgeProtocolFee *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.UpdateBridgeProtocolFee(&_AavePool.TransactOpts, bridgeProtocolFee)
}

// UpdateFlashloanPremiums is a paid mutator transaction binding the contract method 0xbcb6e522.
//
// Solidity: function updateFlashloanPremiums(uint128 flashLoanPremiumTotal, uint128 flashLoanPremiumToProtocol) returns()
func (_AavePool *AavePoolTransactor) UpdateFlashloanPremiums(opts *bind.TransactOpts, flashLoanPremiumTotal *big.Int, flashLoanPremiumToProtocol *big.Int) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "updateFlashloanPremiums", flashLoanPremiumTotal, flashLoanPremiumToProtocol)
}

// UpdateFlashloanPremiums is a paid mutator transaction binding the contract method 0xbcb6e522.
//
// Solidity: function updateFlashloanPremiums(uint128 flashLoanPremiumTotal, uint128 flashLoanPremiumToProtocol) returns()
func (_AavePool *AavePoolSession) UpdateFlashloanPremiums(flashLoanPremiumTotal *big.Int, flashLoanPremiumToProtocol *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.UpdateFlashloanPremiums(&_AavePool.TransactOpts, flashLoanPremiumTotal, flashLoanPremiumToProtocol)
}

// UpdateFlashloanPremiums is a paid mutator transaction binding the contract method 0xbcb6e522.
//
// Solidity: function updateFlashloanPremiums(uint128 flashLoanPremiumTotal, uint128 flashLoanPremiumToProtocol) returns()
func (_AavePool *AavePoolTransactorSession) UpdateFlashloanPremiums(flashLoanPremiumTotal *big.Int, flashLoanPremiumToProtocol *big.Int) (*types.Transaction, error) {
	return _AavePool.Contract.UpdateFlashloanPremiums(&_AavePool.TransactOpts, flashLoanPremiumTotal, flashLoanPremiumToProtocol)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (_AavePool *AavePoolTransactor) Withdraw(opts *bind.TransactOpts, asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AavePool.contract.Transact(opts, "withdraw", asset, amount, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (_AavePool *AavePoolSession) Withdraw(asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.Withdraw(&_AavePool.TransactOpts, asset, amount, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (_AavePool *AavePoolTransactorSession) Withdraw(asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AavePool.Contract.Withdraw(&_AavePool.TransactOpts, asset, amount, to)
}

// AavePoolBackUnbackedIterator is returned from FilterBackUnbacked and is used to iterate over the raw logs and unpacked data for BackUnbacked events raised by the AavePool contract.
type AavePoolBackUnbackedIterator struct {
	Event *AavePoolBackUnbacked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AavePoolBackUnbackedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolBackUnbacked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AavePoolBackUnbacked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AavePoolBackUnbackedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolBackUnbackedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolBackUnbacked represents a BackUnbacked event raised by the AavePool contract.
type AavePoolBackUnbacked struct {
	Reserve common.Address
	Backer  common.Address
	Amount  *big.Int
	Fee     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBackUnbacked is a free log retrieval operation binding the contract event 0x281596e92b2d974beb7d4f124df30a0b39067b096893e95011ce4bdad798b759.
//
// Solidity: event BackUnbacked(address indexed reserve, address indexed backer, uint256 amount, uint256 fee)
func (_AavePool *AavePoolFilterer) FilterBackUnbacked(opts *bind.FilterOpts, reserve []common.Address, backer []common.Address) (*AavePoolBackUnbackedIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var backerRule []interface{}
	for _, backerItem := range backer {
		backerRule = append(backerRule, backerItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "BackUnbacked", reserveRule, backerRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolBackUnbackedIterator{contract: _AavePool.contract, event: "BackUnbacked", logs: logs, sub: sub}, nil
}

// WatchBackUnbacked is a free log subscription operation binding the contract event 0x281596e92b2d974beb7d4f124df30a0b39067b096893e95011ce4bdad798b759.
//
// Solidity: event BackUnbacked(address indexed reserve, address indexed backer, uint256 amount, uint256 fee)
func (_AavePool *AavePoolFilterer) WatchBackUnbacked(opts *bind.WatchOpts, sink chan<- *AavePoolBackUnbacked, reserve []common.Address, backer []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var backerRule []interface{}
	for _, backerItem := range backer {
		backerRule = append(backerRule, backerItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "BackUnbacked", reserveRule, backerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolBackUnbacked)
				if err := _AavePool.contract.UnpackLog(event, "BackUnbacked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBackUnbacked is a log parse operation binding the contract event 0x281596e92b2d974beb7d4f124df30a0b39067b096893e95011ce4bdad798b759.
//
// Solidity: event BackUnbacked(address indexed reserve, address indexed backer, uint256 amount, uint256 fee)
func (_AavePool *AavePoolFilterer) ParseBackUnbacked(log types.Log) (*AavePoolBackUnbacked, error) {
	event := new(AavePoolBackUnbacked)
	if err := _AavePool.contract.UnpackLog(event, "BackUnbacked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the AavePool contract.
type AavePoolBorrowIterator struct {
	Event *AavePoolBorrow // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AavePoolBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolBorrow)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AavePoolBorrow)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AavePoolBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolBorrow represents a Borrow event raised by the AavePool contract.
type AavePoolBorrow struct {
	Reserve          common.Address
	User             common.Address
	OnBehalfOf       common.Address
	Amount           *big.Int
	InterestRateMode uint8
	BorrowRate       *big.Int
	ReferralCode     uint16
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0xb3d084820fb1a9decffb176436bd02558d15fac9b0ddfed8c465bc7359d7dce0.
//
// Solidity: event Borrow(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint8 interestRateMode, uint256 borrowRate, uint16 indexed referralCode)
func (_AavePool *AavePoolFilterer) FilterBorrow(opts *bind.FilterOpts, reserve []common.Address, onBehalfOf []common.Address, referralCode []uint16) (*AavePoolBorrowIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "Borrow", reserveRule, onBehalfOfRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolBorrowIterator{contract: _AavePool.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0xb3d084820fb1a9decffb176436bd02558d15fac9b0ddfed8c465bc7359d7dce0.
//
// Solidity: event Borrow(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint8 interestRateMode, uint256 borrowRate, uint16 indexed referralCode)
func (_AavePool *AavePoolFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *AavePoolBorrow, reserve []common.Address, onBehalfOf []common.Address, referralCode []uint16) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "Borrow", reserveRule, onBehalfOfRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolBorrow)
				if err := _AavePool.contract.UnpackLog(event, "Borrow", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBorrow is a log parse operation binding the contract event 0xb3d084820fb1a9decffb176436bd02558d15fac9b0ddfed8c465bc7359d7dce0.
//
// Solidity: event Borrow(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint8 interestRateMode, uint256 borrowRate, uint16 indexed referralCode)
func (_AavePool *AavePoolFilterer) ParseBorrow(log types.Log) (*AavePoolBorrow, error) {
	event := new(AavePoolBorrow)
	if err := _AavePool.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolFlashLoanIterator is returned from FilterFlashLoan and is used to iterate over the raw logs and unpacked data for FlashLoan events raised by the AavePool contract.
type AavePoolFlashLoanIterator struct {
	Event *AavePoolFlashLoan // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AavePoolFlashLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolFlashLoan)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AavePoolFlashLoan)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AavePoolFlashLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolFlashLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolFlashLoan represents a FlashLoan event raised by the AavePool contract.
type AavePoolFlashLoan struct {
	Target           common.Address
	Initiator        common.Address
	Asset            common.Address
	Amount           *big.Int
	InterestRateMode uint8
	Premium          *big.Int
	ReferralCode     uint16
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFlashLoan is a free log retrieval operation binding the contract event 0xefefaba5e921573100900a3ad9cf29f222d995fb3b6045797eaea7521bd8d6f0.
//
// Solidity: event FlashLoan(address indexed target, address initiator, address indexed asset, uint256 amount, uint8 interestRateMode, uint256 premium, uint16 indexed referralCode)
func (_AavePool *AavePoolFilterer) FilterFlashLoan(opts *bind.FilterOpts, target []common.Address, asset []common.Address, referralCode []uint16) (*AavePoolFlashLoanIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "FlashLoan", targetRule, assetRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolFlashLoanIterator{contract: _AavePool.contract, event: "FlashLoan", logs: logs, sub: sub}, nil
}

// WatchFlashLoan is a free log subscription operation binding the contract event 0xefefaba5e921573100900a3ad9cf29f222d995fb3b6045797eaea7521bd8d6f0.
//
// Solidity: event FlashLoan(address indexed target, address initiator, address indexed asset, uint256 amount, uint8 interestRateMode, uint256 premium, uint16 indexed referralCode)
func (_AavePool *AavePoolFilterer) WatchFlashLoan(opts *bind.WatchOpts, sink chan<- *AavePoolFlashLoan, target []common.Address, asset []common.Address, referralCode []uint16) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "FlashLoan", targetRule, assetRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolFlashLoan)
				if err := _AavePool.contract.UnpackLog(event, "FlashLoan", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFlashLoan is a log parse operation binding the contract event 0xefefaba5e921573100900a3ad9cf29f222d995fb3b6045797eaea7521bd8d6f0.
//
// Solidity: event FlashLoan(address indexed target, address initiator, address indexed asset, uint256 amount, uint8 interestRateMode, uint256 premium, uint16 indexed referralCode)
func (_AavePool *AavePoolFilterer) ParseFlashLoan(log types.Log) (*AavePoolFlashLoan, error) {
	event := new(AavePoolFlashLoan)
	if err := _AavePool.contract.UnpackLog(event, "FlashLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolIsolationModeTotalDebtUpdatedIterator is returned from FilterIsolationModeTotalDebtUpdated and is used to iterate over the raw logs and unpacked data for IsolationModeTotalDebtUpdated events raised by the AavePool contract.
type AavePoolIsolationModeTotalDebtUpdatedIterator struct {
	Event *AavePoolIsolationModeTotalDebtUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AavePoolIsolationModeTotalDebtUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolIsolationModeTotalDebtUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AavePoolIsolationModeTotalDebtUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AavePoolIsolationModeTotalDebtUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolIsolationModeTotalDebtUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolIsolationModeTotalDebtUpdated represents a IsolationModeTotalDebtUpdated event raised by the AavePool contract.
type AavePoolIsolationModeTotalDebtUpdated struct {
	Asset     common.Address
	TotalDebt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIsolationModeTotalDebtUpdated is a free log retrieval operation binding the contract event 0xaef84d3b40895fd58c561f3998000f0583abb992a52fbdc99ace8e8de4d676a5.
//
// Solidity: event IsolationModeTotalDebtUpdated(address indexed asset, uint256 totalDebt)
func (_AavePool *AavePoolFilterer) FilterIsolationModeTotalDebtUpdated(opts *bind.FilterOpts, asset []common.Address) (*AavePoolIsolationModeTotalDebtUpdatedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "IsolationModeTotalDebtUpdated", assetRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolIsolationModeTotalDebtUpdatedIterator{contract: _AavePool.contract, event: "IsolationModeTotalDebtUpdated", logs: logs, sub: sub}, nil
}

// WatchIsolationModeTotalDebtUpdated is a free log subscription operation binding the contract event 0xaef84d3b40895fd58c561f3998000f0583abb992a52fbdc99ace8e8de4d676a5.
//
// Solidity: event IsolationModeTotalDebtUpdated(address indexed asset, uint256 totalDebt)
func (_AavePool *AavePoolFilterer) WatchIsolationModeTotalDebtUpdated(opts *bind.WatchOpts, sink chan<- *AavePoolIsolationModeTotalDebtUpdated, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "IsolationModeTotalDebtUpdated", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolIsolationModeTotalDebtUpdated)
				if err := _AavePool.contract.UnpackLog(event, "IsolationModeTotalDebtUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIsolationModeTotalDebtUpdated is a log parse operation binding the contract event 0xaef84d3b40895fd58c561f3998000f0583abb992a52fbdc99ace8e8de4d676a5.
//
// Solidity: event IsolationModeTotalDebtUpdated(address indexed asset, uint256 totalDebt)
func (_AavePool *AavePoolFilterer) ParseIsolationModeTotalDebtUpdated(log types.Log) (*AavePoolIsolationModeTotalDebtUpdated, error) {
	event := new(AavePoolIsolationModeTotalDebtUpdated)
	if err := _AavePool.contract.UnpackLog(event, "IsolationModeTotalDebtUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolLiquidationCallIterator is returned from FilterLiquidationCall and is used to iterate over the raw logs and unpacked data for LiquidationCall events raised by the AavePool contract.
type AavePoolLiquidationCallIterator struct {
	Event *AavePoolLiquidationCall // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AavePoolLiquidationCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolLiquidationCall)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AavePoolLiquidationCall)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AavePoolLiquidationCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolLiquidationCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolLiquidationCall represents a LiquidationCall event raised by the AavePool contract.
type AavePoolLiquidationCall struct {
	CollateralAsset            common.Address
	DebtAsset                  common.Address
	User                       common.Address
	DebtToCover                *big.Int
	LiquidatedCollateralAmount *big.Int
	Liquidator                 common.Address
	ReceiveAToken              bool
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterLiquidationCall is a free log retrieval operation binding the contract event 0xe413a321e8681d831f4dbccbca790d2952b56f977908e45be37335533e005286.
//
// Solidity: event LiquidationCall(address indexed collateralAsset, address indexed debtAsset, address indexed user, uint256 debtToCover, uint256 liquidatedCollateralAmount, address liquidator, bool receiveAToken)
func (_AavePool *AavePoolFilterer) FilterLiquidationCall(opts *bind.FilterOpts, collateralAsset []common.Address, debtAsset []common.Address, user []common.Address) (*AavePoolLiquidationCallIterator, error) {

	var collateralAssetRule []interface{}
	for _, collateralAssetItem := range collateralAsset {
		collateralAssetRule = append(collateralAssetRule, collateralAssetItem)
	}
	var debtAssetRule []interface{}
	for _, debtAssetItem := range debtAsset {
		debtAssetRule = append(debtAssetRule, debtAssetItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "LiquidationCall", collateralAssetRule, debtAssetRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolLiquidationCallIterator{contract: _AavePool.contract, event: "LiquidationCall", logs: logs, sub: sub}, nil
}

// WatchLiquidationCall is a free log subscription operation binding the contract event 0xe413a321e8681d831f4dbccbca790d2952b56f977908e45be37335533e005286.
//
// Solidity: event LiquidationCall(address indexed collateralAsset, address indexed debtAsset, address indexed user, uint256 debtToCover, uint256 liquidatedCollateralAmount, address liquidator, bool receiveAToken)
func (_AavePool *AavePoolFilterer) WatchLiquidationCall(opts *bind.WatchOpts, sink chan<- *AavePoolLiquidationCall, collateralAsset []common.Address, debtAsset []common.Address, user []common.Address) (event.Subscription, error) {

	var collateralAssetRule []interface{}
	for _, collateralAssetItem := range collateralAsset {
		collateralAssetRule = append(collateralAssetRule, collateralAssetItem)
	}
	var debtAssetRule []interface{}
	for _, debtAssetItem := range debtAsset {
		debtAssetRule = append(debtAssetRule, debtAssetItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "LiquidationCall", collateralAssetRule, debtAssetRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolLiquidationCall)
				if err := _AavePool.contract.UnpackLog(event, "LiquidationCall", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLiquidationCall is a log parse operation binding the contract event 0xe413a321e8681d831f4dbccbca790d2952b56f977908e45be37335533e005286.
//
// Solidity: event LiquidationCall(address indexed collateralAsset, address indexed debtAsset, address indexed user, uint256 debtToCover, uint256 liquidatedCollateralAmount, address liquidator, bool receiveAToken)
func (_AavePool *AavePoolFilterer) ParseLiquidationCall(log types.Log) (*AavePoolLiquidationCall, error) {
	event := new(AavePoolLiquidationCall)
	if err := _AavePool.contract.UnpackLog(event, "LiquidationCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolMintUnbackedIterator is returned from FilterMintUnbacked and is used to iterate over the raw logs and unpacked data for MintUnbacked events raised by the AavePool contract.
type AavePoolMintUnbackedIterator struct {
	Event *AavePoolMintUnbacked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AavePoolMintUnbackedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolMintUnbacked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AavePoolMintUnbacked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AavePoolMintUnbackedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolMintUnbackedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolMintUnbacked represents a MintUnbacked event raised by the AavePool contract.
type AavePoolMintUnbacked struct {
	Reserve      common.Address
	User         common.Address
	OnBehalfOf   common.Address
	Amount       *big.Int
	ReferralCode uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMintUnbacked is a free log retrieval operation binding the contract event 0xf25af37b3d3ec226063dc9bdc103ece7eb110a50f340fe854bb7bc1b0676d7d0.
//
// Solidity: event MintUnbacked(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint16 indexed referralCode)
func (_AavePool *AavePoolFilterer) FilterMintUnbacked(opts *bind.FilterOpts, reserve []common.Address, onBehalfOf []common.Address, referralCode []uint16) (*AavePoolMintUnbackedIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "MintUnbacked", reserveRule, onBehalfOfRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolMintUnbackedIterator{contract: _AavePool.contract, event: "MintUnbacked", logs: logs, sub: sub}, nil
}

// WatchMintUnbacked is a free log subscription operation binding the contract event 0xf25af37b3d3ec226063dc9bdc103ece7eb110a50f340fe854bb7bc1b0676d7d0.
//
// Solidity: event MintUnbacked(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint16 indexed referralCode)
func (_AavePool *AavePoolFilterer) WatchMintUnbacked(opts *bind.WatchOpts, sink chan<- *AavePoolMintUnbacked, reserve []common.Address, onBehalfOf []common.Address, referralCode []uint16) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "MintUnbacked", reserveRule, onBehalfOfRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolMintUnbacked)
				if err := _AavePool.contract.UnpackLog(event, "MintUnbacked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMintUnbacked is a log parse operation binding the contract event 0xf25af37b3d3ec226063dc9bdc103ece7eb110a50f340fe854bb7bc1b0676d7d0.
//
// Solidity: event MintUnbacked(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint16 indexed referralCode)
func (_AavePool *AavePoolFilterer) ParseMintUnbacked(log types.Log) (*AavePoolMintUnbacked, error) {
	event := new(AavePoolMintUnbacked)
	if err := _AavePool.contract.UnpackLog(event, "MintUnbacked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolMintedToTreasuryIterator is returned from FilterMintedToTreasury and is used to iterate over the raw logs and unpacked data for MintedToTreasury events raised by the AavePool contract.
type AavePoolMintedToTreasuryIterator struct {
	Event *AavePoolMintedToTreasury // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AavePoolMintedToTreasuryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolMintedToTreasury)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AavePoolMintedToTreasury)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AavePoolMintedToTreasuryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolMintedToTreasuryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolMintedToTreasury represents a MintedToTreasury event raised by the AavePool contract.
type AavePoolMintedToTreasury struct {
	Reserve      common.Address
	AmountMinted *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMintedToTreasury is a free log retrieval operation binding the contract event 0xbfa21aa5d5f9a1f0120a95e7c0749f389863cbdbfff531aa7339077a5bc919de.
//
// Solidity: event MintedToTreasury(address indexed reserve, uint256 amountMinted)
func (_AavePool *AavePoolFilterer) FilterMintedToTreasury(opts *bind.FilterOpts, reserve []common.Address) (*AavePoolMintedToTreasuryIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "MintedToTreasury", reserveRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolMintedToTreasuryIterator{contract: _AavePool.contract, event: "MintedToTreasury", logs: logs, sub: sub}, nil
}

// WatchMintedToTreasury is a free log subscription operation binding the contract event 0xbfa21aa5d5f9a1f0120a95e7c0749f389863cbdbfff531aa7339077a5bc919de.
//
// Solidity: event MintedToTreasury(address indexed reserve, uint256 amountMinted)
func (_AavePool *AavePoolFilterer) WatchMintedToTreasury(opts *bind.WatchOpts, sink chan<- *AavePoolMintedToTreasury, reserve []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "MintedToTreasury", reserveRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolMintedToTreasury)
				if err := _AavePool.contract.UnpackLog(event, "MintedToTreasury", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMintedToTreasury is a log parse operation binding the contract event 0xbfa21aa5d5f9a1f0120a95e7c0749f389863cbdbfff531aa7339077a5bc919de.
//
// Solidity: event MintedToTreasury(address indexed reserve, uint256 amountMinted)
func (_AavePool *AavePoolFilterer) ParseMintedToTreasury(log types.Log) (*AavePoolMintedToTreasury, error) {
	event := new(AavePoolMintedToTreasury)
	if err := _AavePool.contract.UnpackLog(event, "MintedToTreasury", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolRebalanceStableBorrowRateIterator is returned from FilterRebalanceStableBorrowRate and is used to iterate over the raw logs and unpacked data for RebalanceStableBorrowRate events raised by the AavePool contract.
type AavePoolRebalanceStableBorrowRateIterator struct {
	Event *AavePoolRebalanceStableBorrowRate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AavePoolRebalanceStableBorrowRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolRebalanceStableBorrowRate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AavePoolRebalanceStableBorrowRate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AavePoolRebalanceStableBorrowRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolRebalanceStableBorrowRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolRebalanceStableBorrowRate represents a RebalanceStableBorrowRate event raised by the AavePool contract.
type AavePoolRebalanceStableBorrowRate struct {
	Reserve common.Address
	User    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRebalanceStableBorrowRate is a free log retrieval operation binding the contract event 0x9f439ae0c81e41a04d3fdfe07aed54e6a179fb0db15be7702eb66fa8ef6f5300.
//
// Solidity: event RebalanceStableBorrowRate(address indexed reserve, address indexed user)
func (_AavePool *AavePoolFilterer) FilterRebalanceStableBorrowRate(opts *bind.FilterOpts, reserve []common.Address, user []common.Address) (*AavePoolRebalanceStableBorrowRateIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "RebalanceStableBorrowRate", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolRebalanceStableBorrowRateIterator{contract: _AavePool.contract, event: "RebalanceStableBorrowRate", logs: logs, sub: sub}, nil
}

// WatchRebalanceStableBorrowRate is a free log subscription operation binding the contract event 0x9f439ae0c81e41a04d3fdfe07aed54e6a179fb0db15be7702eb66fa8ef6f5300.
//
// Solidity: event RebalanceStableBorrowRate(address indexed reserve, address indexed user)
func (_AavePool *AavePoolFilterer) WatchRebalanceStableBorrowRate(opts *bind.WatchOpts, sink chan<- *AavePoolRebalanceStableBorrowRate, reserve []common.Address, user []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "RebalanceStableBorrowRate", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolRebalanceStableBorrowRate)
				if err := _AavePool.contract.UnpackLog(event, "RebalanceStableBorrowRate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRebalanceStableBorrowRate is a log parse operation binding the contract event 0x9f439ae0c81e41a04d3fdfe07aed54e6a179fb0db15be7702eb66fa8ef6f5300.
//
// Solidity: event RebalanceStableBorrowRate(address indexed reserve, address indexed user)
func (_AavePool *AavePoolFilterer) ParseRebalanceStableBorrowRate(log types.Log) (*AavePoolRebalanceStableBorrowRate, error) {
	event := new(AavePoolRebalanceStableBorrowRate)
	if err := _AavePool.contract.UnpackLog(event, "RebalanceStableBorrowRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the AavePool contract.
type AavePoolRepayIterator struct {
	Event *AavePoolRepay // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AavePoolRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolRepay)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AavePoolRepay)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AavePoolRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolRepay represents a Repay event raised by the AavePool contract.
type AavePoolRepay struct {
	Reserve    common.Address
	User       common.Address
	Repayer    common.Address
	Amount     *big.Int
	UseATokens bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0xa534c8dbe71f871f9f3530e97a74601fea17b426cae02e1c5aee42c96c784051.
//
// Solidity: event Repay(address indexed reserve, address indexed user, address indexed repayer, uint256 amount, bool useATokens)
func (_AavePool *AavePoolFilterer) FilterRepay(opts *bind.FilterOpts, reserve []common.Address, user []common.Address, repayer []common.Address) (*AavePoolRepayIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var repayerRule []interface{}
	for _, repayerItem := range repayer {
		repayerRule = append(repayerRule, repayerItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "Repay", reserveRule, userRule, repayerRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolRepayIterator{contract: _AavePool.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0xa534c8dbe71f871f9f3530e97a74601fea17b426cae02e1c5aee42c96c784051.
//
// Solidity: event Repay(address indexed reserve, address indexed user, address indexed repayer, uint256 amount, bool useATokens)
func (_AavePool *AavePoolFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *AavePoolRepay, reserve []common.Address, user []common.Address, repayer []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var repayerRule []interface{}
	for _, repayerItem := range repayer {
		repayerRule = append(repayerRule, repayerItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "Repay", reserveRule, userRule, repayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolRepay)
				if err := _AavePool.contract.UnpackLog(event, "Repay", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRepay is a log parse operation binding the contract event 0xa534c8dbe71f871f9f3530e97a74601fea17b426cae02e1c5aee42c96c784051.
//
// Solidity: event Repay(address indexed reserve, address indexed user, address indexed repayer, uint256 amount, bool useATokens)
func (_AavePool *AavePoolFilterer) ParseRepay(log types.Log) (*AavePoolRepay, error) {
	event := new(AavePoolRepay)
	if err := _AavePool.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolReserveDataUpdatedIterator is returned from FilterReserveDataUpdated and is used to iterate over the raw logs and unpacked data for ReserveDataUpdated events raised by the AavePool contract.
type AavePoolReserveDataUpdatedIterator struct {
	Event *AavePoolReserveDataUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AavePoolReserveDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolReserveDataUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AavePoolReserveDataUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AavePoolReserveDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolReserveDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolReserveDataUpdated represents a ReserveDataUpdated event raised by the AavePool contract.
type AavePoolReserveDataUpdated struct {
	Reserve             common.Address
	LiquidityRate       *big.Int
	StableBorrowRate    *big.Int
	VariableBorrowRate  *big.Int
	LiquidityIndex      *big.Int
	VariableBorrowIndex *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReserveDataUpdated is a free log retrieval operation binding the contract event 0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a.
//
// Solidity: event ReserveDataUpdated(address indexed reserve, uint256 liquidityRate, uint256 stableBorrowRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_AavePool *AavePoolFilterer) FilterReserveDataUpdated(opts *bind.FilterOpts, reserve []common.Address) (*AavePoolReserveDataUpdatedIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "ReserveDataUpdated", reserveRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolReserveDataUpdatedIterator{contract: _AavePool.contract, event: "ReserveDataUpdated", logs: logs, sub: sub}, nil
}

// WatchReserveDataUpdated is a free log subscription operation binding the contract event 0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a.
//
// Solidity: event ReserveDataUpdated(address indexed reserve, uint256 liquidityRate, uint256 stableBorrowRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_AavePool *AavePoolFilterer) WatchReserveDataUpdated(opts *bind.WatchOpts, sink chan<- *AavePoolReserveDataUpdated, reserve []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "ReserveDataUpdated", reserveRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolReserveDataUpdated)
				if err := _AavePool.contract.UnpackLog(event, "ReserveDataUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReserveDataUpdated is a log parse operation binding the contract event 0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a.
//
// Solidity: event ReserveDataUpdated(address indexed reserve, uint256 liquidityRate, uint256 stableBorrowRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_AavePool *AavePoolFilterer) ParseReserveDataUpdated(log types.Log) (*AavePoolReserveDataUpdated, error) {
	event := new(AavePoolReserveDataUpdated)
	if err := _AavePool.contract.UnpackLog(event, "ReserveDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolReserveUsedAsCollateralDisabledIterator is returned from FilterReserveUsedAsCollateralDisabled and is used to iterate over the raw logs and unpacked data for ReserveUsedAsCollateralDisabled events raised by the AavePool contract.
type AavePoolReserveUsedAsCollateralDisabledIterator struct {
	Event *AavePoolReserveUsedAsCollateralDisabled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AavePoolReserveUsedAsCollateralDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolReserveUsedAsCollateralDisabled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AavePoolReserveUsedAsCollateralDisabled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AavePoolReserveUsedAsCollateralDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolReserveUsedAsCollateralDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolReserveUsedAsCollateralDisabled represents a ReserveUsedAsCollateralDisabled event raised by the AavePool contract.
type AavePoolReserveUsedAsCollateralDisabled struct {
	Reserve common.Address
	User    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReserveUsedAsCollateralDisabled is a free log retrieval operation binding the contract event 0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd.
//
// Solidity: event ReserveUsedAsCollateralDisabled(address indexed reserve, address indexed user)
func (_AavePool *AavePoolFilterer) FilterReserveUsedAsCollateralDisabled(opts *bind.FilterOpts, reserve []common.Address, user []common.Address) (*AavePoolReserveUsedAsCollateralDisabledIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "ReserveUsedAsCollateralDisabled", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolReserveUsedAsCollateralDisabledIterator{contract: _AavePool.contract, event: "ReserveUsedAsCollateralDisabled", logs: logs, sub: sub}, nil
}

// WatchReserveUsedAsCollateralDisabled is a free log subscription operation binding the contract event 0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd.
//
// Solidity: event ReserveUsedAsCollateralDisabled(address indexed reserve, address indexed user)
func (_AavePool *AavePoolFilterer) WatchReserveUsedAsCollateralDisabled(opts *bind.WatchOpts, sink chan<- *AavePoolReserveUsedAsCollateralDisabled, reserve []common.Address, user []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "ReserveUsedAsCollateralDisabled", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolReserveUsedAsCollateralDisabled)
				if err := _AavePool.contract.UnpackLog(event, "ReserveUsedAsCollateralDisabled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReserveUsedAsCollateralDisabled is a log parse operation binding the contract event 0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd.
//
// Solidity: event ReserveUsedAsCollateralDisabled(address indexed reserve, address indexed user)
func (_AavePool *AavePoolFilterer) ParseReserveUsedAsCollateralDisabled(log types.Log) (*AavePoolReserveUsedAsCollateralDisabled, error) {
	event := new(AavePoolReserveUsedAsCollateralDisabled)
	if err := _AavePool.contract.UnpackLog(event, "ReserveUsedAsCollateralDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolReserveUsedAsCollateralEnabledIterator is returned from FilterReserveUsedAsCollateralEnabled and is used to iterate over the raw logs and unpacked data for ReserveUsedAsCollateralEnabled events raised by the AavePool contract.
type AavePoolReserveUsedAsCollateralEnabledIterator struct {
	Event *AavePoolReserveUsedAsCollateralEnabled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AavePoolReserveUsedAsCollateralEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolReserveUsedAsCollateralEnabled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AavePoolReserveUsedAsCollateralEnabled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AavePoolReserveUsedAsCollateralEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolReserveUsedAsCollateralEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolReserveUsedAsCollateralEnabled represents a ReserveUsedAsCollateralEnabled event raised by the AavePool contract.
type AavePoolReserveUsedAsCollateralEnabled struct {
	Reserve common.Address
	User    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReserveUsedAsCollateralEnabled is a free log retrieval operation binding the contract event 0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2.
//
// Solidity: event ReserveUsedAsCollateralEnabled(address indexed reserve, address indexed user)
func (_AavePool *AavePoolFilterer) FilterReserveUsedAsCollateralEnabled(opts *bind.FilterOpts, reserve []common.Address, user []common.Address) (*AavePoolReserveUsedAsCollateralEnabledIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "ReserveUsedAsCollateralEnabled", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolReserveUsedAsCollateralEnabledIterator{contract: _AavePool.contract, event: "ReserveUsedAsCollateralEnabled", logs: logs, sub: sub}, nil
}

// WatchReserveUsedAsCollateralEnabled is a free log subscription operation binding the contract event 0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2.
//
// Solidity: event ReserveUsedAsCollateralEnabled(address indexed reserve, address indexed user)
func (_AavePool *AavePoolFilterer) WatchReserveUsedAsCollateralEnabled(opts *bind.WatchOpts, sink chan<- *AavePoolReserveUsedAsCollateralEnabled, reserve []common.Address, user []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "ReserveUsedAsCollateralEnabled", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolReserveUsedAsCollateralEnabled)
				if err := _AavePool.contract.UnpackLog(event, "ReserveUsedAsCollateralEnabled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReserveUsedAsCollateralEnabled is a log parse operation binding the contract event 0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2.
//
// Solidity: event ReserveUsedAsCollateralEnabled(address indexed reserve, address indexed user)
func (_AavePool *AavePoolFilterer) ParseReserveUsedAsCollateralEnabled(log types.Log) (*AavePoolReserveUsedAsCollateralEnabled, error) {
	event := new(AavePoolReserveUsedAsCollateralEnabled)
	if err := _AavePool.contract.UnpackLog(event, "ReserveUsedAsCollateralEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolSupplyIterator is returned from FilterSupply and is used to iterate over the raw logs and unpacked data for Supply events raised by the AavePool contract.
type AavePoolSupplyIterator struct {
	Event *AavePoolSupply // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AavePoolSupplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolSupply)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AavePoolSupply)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AavePoolSupplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolSupplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolSupply represents a Supply event raised by the AavePool contract.
type AavePoolSupply struct {
	Reserve      common.Address
	User         common.Address
	OnBehalfOf   common.Address
	Amount       *big.Int
	ReferralCode uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSupply is a free log retrieval operation binding the contract event 0x2b627736bca15cd5381dcf80b0bf11fd197d01a037c52b927a881a10fb73ba61.
//
// Solidity: event Supply(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint16 indexed referralCode)
func (_AavePool *AavePoolFilterer) FilterSupply(opts *bind.FilterOpts, reserve []common.Address, onBehalfOf []common.Address, referralCode []uint16) (*AavePoolSupplyIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "Supply", reserveRule, onBehalfOfRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolSupplyIterator{contract: _AavePool.contract, event: "Supply", logs: logs, sub: sub}, nil
}

// WatchSupply is a free log subscription operation binding the contract event 0x2b627736bca15cd5381dcf80b0bf11fd197d01a037c52b927a881a10fb73ba61.
//
// Solidity: event Supply(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint16 indexed referralCode)
func (_AavePool *AavePoolFilterer) WatchSupply(opts *bind.WatchOpts, sink chan<- *AavePoolSupply, reserve []common.Address, onBehalfOf []common.Address, referralCode []uint16) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "Supply", reserveRule, onBehalfOfRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolSupply)
				if err := _AavePool.contract.UnpackLog(event, "Supply", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSupply is a log parse operation binding the contract event 0x2b627736bca15cd5381dcf80b0bf11fd197d01a037c52b927a881a10fb73ba61.
//
// Solidity: event Supply(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint16 indexed referralCode)
func (_AavePool *AavePoolFilterer) ParseSupply(log types.Log) (*AavePoolSupply, error) {
	event := new(AavePoolSupply)
	if err := _AavePool.contract.UnpackLog(event, "Supply", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolSwapBorrowRateModeIterator is returned from FilterSwapBorrowRateMode and is used to iterate over the raw logs and unpacked data for SwapBorrowRateMode events raised by the AavePool contract.
type AavePoolSwapBorrowRateModeIterator struct {
	Event *AavePoolSwapBorrowRateMode // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AavePoolSwapBorrowRateModeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolSwapBorrowRateMode)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AavePoolSwapBorrowRateMode)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AavePoolSwapBorrowRateModeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolSwapBorrowRateModeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolSwapBorrowRateMode represents a SwapBorrowRateMode event raised by the AavePool contract.
type AavePoolSwapBorrowRateMode struct {
	Reserve          common.Address
	User             common.Address
	InterestRateMode uint8
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSwapBorrowRateMode is a free log retrieval operation binding the contract event 0x7962b394d85a534033ba2efcf43cd36de57b7ebeb3de0ca4428965d9b3ddc481.
//
// Solidity: event SwapBorrowRateMode(address indexed reserve, address indexed user, uint8 interestRateMode)
func (_AavePool *AavePoolFilterer) FilterSwapBorrowRateMode(opts *bind.FilterOpts, reserve []common.Address, user []common.Address) (*AavePoolSwapBorrowRateModeIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "SwapBorrowRateMode", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolSwapBorrowRateModeIterator{contract: _AavePool.contract, event: "SwapBorrowRateMode", logs: logs, sub: sub}, nil
}

// WatchSwapBorrowRateMode is a free log subscription operation binding the contract event 0x7962b394d85a534033ba2efcf43cd36de57b7ebeb3de0ca4428965d9b3ddc481.
//
// Solidity: event SwapBorrowRateMode(address indexed reserve, address indexed user, uint8 interestRateMode)
func (_AavePool *AavePoolFilterer) WatchSwapBorrowRateMode(opts *bind.WatchOpts, sink chan<- *AavePoolSwapBorrowRateMode, reserve []common.Address, user []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "SwapBorrowRateMode", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolSwapBorrowRateMode)
				if err := _AavePool.contract.UnpackLog(event, "SwapBorrowRateMode", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwapBorrowRateMode is a log parse operation binding the contract event 0x7962b394d85a534033ba2efcf43cd36de57b7ebeb3de0ca4428965d9b3ddc481.
//
// Solidity: event SwapBorrowRateMode(address indexed reserve, address indexed user, uint8 interestRateMode)
func (_AavePool *AavePoolFilterer) ParseSwapBorrowRateMode(log types.Log) (*AavePoolSwapBorrowRateMode, error) {
	event := new(AavePoolSwapBorrowRateMode)
	if err := _AavePool.contract.UnpackLog(event, "SwapBorrowRateMode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolUserEModeSetIterator is returned from FilterUserEModeSet and is used to iterate over the raw logs and unpacked data for UserEModeSet events raised by the AavePool contract.
type AavePoolUserEModeSetIterator struct {
	Event *AavePoolUserEModeSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AavePoolUserEModeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolUserEModeSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AavePoolUserEModeSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AavePoolUserEModeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolUserEModeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolUserEModeSet represents a UserEModeSet event raised by the AavePool contract.
type AavePoolUserEModeSet struct {
	User       common.Address
	CategoryId uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUserEModeSet is a free log retrieval operation binding the contract event 0xd728da875fc88944cbf17638bcbe4af0eedaef63becd1d1c57cc097eb4608d84.
//
// Solidity: event UserEModeSet(address indexed user, uint8 categoryId)
func (_AavePool *AavePoolFilterer) FilterUserEModeSet(opts *bind.FilterOpts, user []common.Address) (*AavePoolUserEModeSetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "UserEModeSet", userRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolUserEModeSetIterator{contract: _AavePool.contract, event: "UserEModeSet", logs: logs, sub: sub}, nil
}

// WatchUserEModeSet is a free log subscription operation binding the contract event 0xd728da875fc88944cbf17638bcbe4af0eedaef63becd1d1c57cc097eb4608d84.
//
// Solidity: event UserEModeSet(address indexed user, uint8 categoryId)
func (_AavePool *AavePoolFilterer) WatchUserEModeSet(opts *bind.WatchOpts, sink chan<- *AavePoolUserEModeSet, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "UserEModeSet", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolUserEModeSet)
				if err := _AavePool.contract.UnpackLog(event, "UserEModeSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUserEModeSet is a log parse operation binding the contract event 0xd728da875fc88944cbf17638bcbe4af0eedaef63becd1d1c57cc097eb4608d84.
//
// Solidity: event UserEModeSet(address indexed user, uint8 categoryId)
func (_AavePool *AavePoolFilterer) ParseUserEModeSet(log types.Log) (*AavePoolUserEModeSet, error) {
	event := new(AavePoolUserEModeSet)
	if err := _AavePool.contract.UnpackLog(event, "UserEModeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AavePoolWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the AavePool contract.
type AavePoolWithdrawIterator struct {
	Event *AavePoolWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AavePoolWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePoolWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AavePoolWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AavePoolWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePoolWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePoolWithdraw represents a Withdraw event raised by the AavePool contract.
type AavePoolWithdraw struct {
	Reserve common.Address
	User    common.Address
	To      common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7.
//
// Solidity: event Withdraw(address indexed reserve, address indexed user, address indexed to, uint256 amount)
func (_AavePool *AavePoolFilterer) FilterWithdraw(opts *bind.FilterOpts, reserve []common.Address, user []common.Address, to []common.Address) (*AavePoolWithdrawIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AavePool.contract.FilterLogs(opts, "Withdraw", reserveRule, userRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AavePoolWithdrawIterator{contract: _AavePool.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7.
//
// Solidity: event Withdraw(address indexed reserve, address indexed user, address indexed to, uint256 amount)
func (_AavePool *AavePoolFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *AavePoolWithdraw, reserve []common.Address, user []common.Address, to []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AavePool.contract.WatchLogs(opts, "Withdraw", reserveRule, userRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePoolWithdraw)
				if err := _AavePool.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7.
//
// Solidity: event Withdraw(address indexed reserve, address indexed user, address indexed to, uint256 amount)
func (_AavePool *AavePoolFilterer) ParseWithdraw(log types.Log) (*AavePoolWithdraw, error) {
	event := new(AavePoolWithdraw)
	if err := _AavePool.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
