package main

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// var cdc *amino.Codec
//
// func init() {
// 	cdc := codec.New()
// 	sdk.RegisterCodec(cdc)
// }

func makeTestCodec() (cdc *codec.Codec) {
	cdc = codec.New()
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	types.RegisterCodec(cdc)
	return
}

func GetValidators(cliCtx context.CLIContext) (types.Validators, error) {
	cliCtx = cliCtx.WithCodec(makeTestCodec())

	resKVs, _, err := cliCtx.QuerySubspace(types.ValidatorsKey, types.StoreKey)
	if err != nil {
		return nil, err
	}

	var validators types.Validators
	for _, kv := range resKVs {
		fmt.Println("111111111111")
		validator, err := types.UnmarshalValidator(types.ModuleCdc, kv.Value)
		if err != nil {
			return nil, err
		}

		validators = append(validators, validator)
	}

	return validators, nil
}
