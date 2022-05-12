* *Aave *
    => ABI is created through the following command in the build directory.
    => Create BIN file for the same contract through solidity compiler through this command in the build directory.
    => will Create a api go file that will allow you to interact with contract

# Deposit
ex: explain the function step by step
    => interact w/ Pool Contract on V3
    => Call *supply(address asset, uint256 amount, address onBehalfOf, uint16 referralCode)* on Pool V3

# Borrow
ex: explain the function step by step
    => interact w/ Pool Contract on V3
    => Call *borrow(address asset, uint256 amount, uint256 interestRateMode, uint16 referralCode, address onBehalfOf)* on Pool V3

# Repay
ex: explain the function step by step
    => interact w/ Pool Contract on V3
    => Call *repay(address asset, uint256 amount, uint256 rateMode, address onBehalfOf)* on Pool V3
    => Call *repayWithATokens(address asset, uint256 amount, uint256 interestRateMode)* on Pool V3

# Flashloan
ex: explain the function step by step
    => interact w/ Pool Contract on V3
    => Call *flashLoanSimple(address receiverAddress, address asset, uint256 amount, bytes calldata params, uint16 referralCode)* on Pool V3
    => Call *flashLoan(address receiverAddress, address[] calldata assets, uint256[] calldata amounts, uint256[] interestRateModes, address onBehalfOf, bytes calldata params, uint16 referralCode)* on Pool V3




* *How to connect with Smart-contract*
1. Connect with EVM or (Ganache).
2. Connect any account to make the transaction.
3. Deploying a smart contract with an admin deployer.
4. Creating endpoints.
5. Connecting with smart contract functions.
6. Using multiple accounts for transactions.