* *How to compile with Smart-contract*
# ABI
solc --optimize --abi ./contracts/"FileHere.sol" -o build
ex: solc --optimize --abi ./contracts/interfaces/*.sol -o build
    => They will put the abi in a build folder
    => Compile the contracts and create a ABI file of our contract though solidity compiler which we install above, 
    => ABI is created through the following command in the build directory.

    solc --optimize --abi ./contracts/interfaces/*.sol -o build

# BIN
solc --optimize --bin ./contracts/"FileHere.sol" -o build
    ex: solc --optimize --abi ./contracts/interfaces/*.sol -o build
    => Create BIN file for the same contract through solidity compiler through this command in the build directory.

    solc --optimize --bin .token/WETH.sol -o build

# abigen
abigen --abi=./build/BalanceCheck.abi  --bin=./build/BalanceCheck.bin --pkg=api --out=./api/BalanceCheck.go

    => will Create a api go file that will allow you to interact with contract
    abigen --abi=./build/*.abi  --bin=./build/*.bin --pkg=defi_api --out=./defi_api/aave.go
    abigen --abi=./build/*.abi  --bin=./build/*.bin --pkg=defi_api --out=./defi_api/aave.go

        abigen --abi=./build/AggregatorInterface.abi  --bin=./build/AggregatorInterface.bin --pkg=chainlink_defi --out=./defi/chainlink/api/chainlink.go
        abigen --abi=./build/IERC20.abi  --bin=./build/IERC20.bin --pkg=token --out=./token/erc20.go
        abigen --abi=./build/WETH.abi  --bin=./build/WETH.bin --pkg=token --out=.//token/weth.go
        abigen --abi=./build/IPool.abi  --bin=./build/IPool.bin --pkg=aave --out=./defi_api/aave/aaveV3_lending.go
                abigen --abi=./build/IPoolAddressesProvider.abi  --bin=./build/IPoolAddressesProvider.bin --pkg=defi_api --out=./defi_api/aave/aaveV3_lending.go
        abigen --abi=./build/DataTypes.abi  --bin=./build/DataTypes.bin --pkg=defi_api --out=./defi_api/aave_lending.go
        abigen --abi=./build/IPoolAddressesProvider.abi  --bin=./build/IPoolAddressesProvider.bin --pkg=aave_getter --out=./aave/getter/aave_V3_lending_provider.go


# Solc-Select
- switch compiler
* *How to connect with Smart-contract*
1. Connect with EVM or (Ganache).
2. Connect any account to make the transaction.
3. Deploying a smart contract with an admin deployer.
4. Creating endpoints.
5. Connecting with smart contract functions.
6. Using multiple accounts for transactions.



pool_address_provider
    abigen --abi=./build/IPoolAddressesProvider.abi  --bin=./build/IPoolAddressesProvider.bin --pkg=pool_address_provider --out=./pool_address_provider/pool_address_provider.go

pool
    abigen --abi=./build/IPool.abi  --bin=./build/IPool.bin --pkg=aave_pool --out=./aave_pool/pool.go


aave_lending_eth
    abigen --abi=./build/IWETHGateway.abi  --bin=./build/IWETHGateway.bin --pkg=aave_lending_eth --out=./aave_lending_eth/weth_gateway.go
token
chainlink_agg


