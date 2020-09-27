package main

import (
	"encoding/json"
	"testing"
)

func TestValidatorsByHeight(t *testing.T) {

	height := 79895

	cfg := NodeConfig{
		RPCNode:           "tcp://192.168.3.200:26657",
		APIServerEndpoint: "http://192.168.3.100:8545",
	}

	cli := NewClient(&cfg)

	valSet, err := cli.ValidatorSet(int64(height))
	if err != nil {
		t.Errorf("validatorset  error: %v \n", err.Error())
	}

	jsonValSetAsJson, err := json.MarshalIndent(valSet, "", "\t")
	if err != nil {
		t.Errorf("validator json marshal err:  %v \n", err.Error())
	}

	t.Logf("height:  %v  \n   %s \n", height, jsonValSetAsJson)

}

//Validators()

func TestValidators(t *testing.T) {

	cfg := NodeConfig{
		RPCNode:           "tcp://192.168.3.200:26657",
		APIServerEndpoint: "http://192.168.3.100:8545",
	}

	cli := NewClient(&cfg)

	valSet, err := cli.Validators()
	if err != nil {
		t.Errorf("validatorset  error: %v \n", err.Error())
	}

	jsonValSetAsJson, err := json.MarshalIndent(valSet, "", "\t")
	if err != nil {
		t.Errorf("validator json marshal err:  %v \n", err.Error())
	}

	t.Logf("height: %s \n", jsonValSetAsJson)

}
