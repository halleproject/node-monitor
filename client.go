package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/ethermint/app"
	"github.com/cosmos/ethermint/codec"
	"github.com/tendermint/tendermint/libs/log"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

type Client struct {
	cliCtx      context.CLIContext
	timerTicker *TimerTicker
	Logger      log.Logger
	Moniker     string
}

func NewClient(nodeRpc string) (*Client, error) {
	//cliCtx := context.NewCLIContext().WithNodeURI("tcp://192.168.3.200:26657")

	cli := Client{
		cliCtx: context.NewCLIContext().WithNodeURI(nodeRpc),
		Logger: log.NewTMLogger(os.Stdout),
	}

	status, err := getNodeStatus(cli.cliCtx)
	if err != nil {
		return nil, err
	}

	cli.Moniker = status.NodeInfo.Moniker
	cli.timerTicker = NewTimerticker(&cli)

	cli.timerTicker.OnStart()

	return &cli, nil
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
