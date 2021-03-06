package main

import (
	"testing"
)

func TestValidators(t *testing.T) {

	cli, err := NewClient("tcp://192.168.3.200:26657")
	if err != nil {
		t.Error(err)
	}

	vals, err := cli.GetValidators()

	if err != nil {
		t.Error(err)
	}

	for k, v := range vals {
		t.Logf("k: %v v : %v \n", k, v)
	}

}

func TestLatestBlockHeight(t *testing.T) {

	cli, err := NewClient("tcp://192.168.3.200:26657")
	if err != nil {
		t.Error(err)
	}

	height, err := cli.LatestBlockHeight()

	if err != nil {
		t.Error(err)
	}

	t.Logf("latest height: %v \n", height)

}
