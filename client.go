package main

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/ethermint/app"
	"github.com/cosmos/ethermint/codec"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

type Client struct {
	cliCtx context.CLIContext
}

func NewClient(nodeRpc string) *Client {
	//cliCtx := context.NewCLIContext().WithNodeURI("tcp://192.168.3.200:26657")
	cli := Client{
		cliCtx: context.NewCLIContext().WithNodeURI(nodeRpc),
	}

	return &cli
}

// LatestBlockHeight returns the latest block height on the active chain
func (c *Client) LatestBlockHeight() (int64, error) {
	status, err := getNodeStatus(c.cliCtx)
	if err != nil {
		return 0, err
	}
	return status.SyncInfo.LatestBlockHeight, nil
}

func getNodeStatus(cliCtx context.CLIContext) (*ctypes.ResultStatus, error) {
	node, err := cliCtx.GetNode()
	if err != nil {
		return &ctypes.ResultStatus{}, err
	}

	return node.Status()
}

func (c *Client) GetValidators() (types.Validators, error) {

	//cliCtx := context.NewCLIContext().WithNodeURI("tcp://192.168.3.200:26657")

	cliCtx := c.cliCtx
	cdc := codec.MakeCodec(app.ModuleBasics)

	cliCtx = cliCtx.WithCodec(cdc)

	resKVs, _, err := cliCtx.QuerySubspace(types.ValidatorsKey, types.StoreKey)
	if err != nil {
		return nil, err
	}

	var validators types.Validators
	for _, kv := range resKVs {
		validator, err := types.UnmarshalValidator(types.ModuleCdc, kv.Value)
		if err != nil {
			return nil, err
		}

		validators = append(validators, validator)
	}

	return validators, nil
}
