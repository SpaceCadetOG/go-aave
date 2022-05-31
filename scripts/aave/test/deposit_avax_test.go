package aave

import (
	"go-aave/scripts/aave"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestDepositAVAX(t *testing.T) {
	type args struct {
		blockchain *ethclient.Client
		deployer   *bind.TransactOpts
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aave.DepositAVAX(tt.args.blockchain, tt.args.deployer)
		})
	}
}
