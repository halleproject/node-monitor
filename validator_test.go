package main

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/client/context"
)

func TestValidators(t *testing.T) {
	cliCtx := context.NewCLIContext().WithNodeURI("tcp://192.168.3.200:26657")

	vals, err := GetValidators(cliCtx)

	if err != nil {
		t.Error(err)
	}

	for k, v := range vals {
		t.Logf("k: %v v : %v \n", k, v)
	}

}
