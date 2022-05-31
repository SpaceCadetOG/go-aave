package aave

import (
	"go-aave/scripts/aave"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestDepositDAI(t *testing.T) {
	type args struct {
		blockchain    *ethclient.Client
		deployer      *bind.TransactOpts
		token_address common.Address
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aave.DepositDAI(tt.args.blockchain, tt.args.deployer, tt.args.token_address)
		})
	}
}
