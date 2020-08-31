package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/Robbin-Liu/go-binance-sdk/client/rpc"
	cmtypes "github.com/Robbin-Liu/go-binance-sdk/common/types"
	resty "github.com/go-resty/resty/v2"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type NodeConfig struct {
	RPCNode           string               `yaml:"rpc_node"`
	APIServerEndpoint string               `yaml:"api_server_endpoint"`
	NetworkType       cmtypes.ChainNetwork `yaml:"network_type"`
}

func ParseConfig() *NodeConfig {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(errors.Wrap(err, "failed to read config"))
	}

	var node *NodeConfig
	if viper.GetString("active") == "" {
		log.Fatal("define active param in your config file.")
	}

	switch viper.GetString("active") {
	case "mainnet":
		node = &NodeConfig{
			RPCNode:           viper.GetString("mainnet.node.rpc_node"),
			APIServerEndpoint: viper.GetString("mainnet.node.api_server_endpoint"),
			NetworkType:       cmtypes.ProdNetwork,
		}

	case "testnet":
		node = &NodeConfig{
			RPCNode:           viper.GetString("testnet.node.rpc_node"),
			APIServerEndpoint: viper.GetString("testnet.node.api_server_endpoint"),
			NetworkType:       cmtypes.ProdNetwork, // ProdNetwork, TestNetwork
		}

	default:
		log.Fatal("active can be either mainnet or testnet.")
	}

	return node
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Client struct {
	rpcClient rpc.Client
	apiClient *resty.Client
}

func NewClient(cfg *NodeConfig) *Client {
	rpcClient := rpc.NewRPCClient(cfg.RPCNode, cfg.NetworkType)
	apiClient := resty.New().
		SetHostURL(cfg.APIServerEndpoint).
		SetTimeout(time.Duration(5 * time.Second))

	return &Client{rpcClient, apiClient}
}

// Block queries for a block by height. An error is returned if the query fails.
func (c Client) Block(height int64) (*tmctypes.ResultBlock, error) {
	return c.rpcClient.Block(&height)
}

// LatestBlockHeight returns the latest block height on the active chain
func (c Client) LatestBlockHeight() (int64, error) {
	status, err := c.rpcClient.Status()
	if err != nil {
		return -1, err
	}

	height := status.SyncInfo.LatestBlockHeight

	return height, nil
}

func (c Client) ValidatorSet(height int64) (*tmctypes.ResultValidators, error) {
	return c.rpcClient.Validators(&height)
}

type Validator struct {
	AccountAddress   string          `json:"account_address"`
	OperatorAddress  string          `json:"operator_address" sql:",notnull, unique"`
	ConsensusPubKey  json.RawMessage `json:"consensus_pubkey" sql:",notnull, unique"`
	ConsensusAddress string          `json:"consensus_address" sql:",notnull, unique"`
	Jailed           bool            `json:"jailed"`
	Status           int64           `json:"status"`
	Tokens           string          `json:"tokens"`
	DelegatorShares  string          `json:"delegator_shares"`
	//Description       Description     `json:"description"`
	UnbondingHeight int64  `json:"unbonding_height"`
	UnbondingTime   string `json:"unbonding_time"`
	//Commission        Commission      `json:"commission"`
	MinSelfDelegation string `json:"tokens"`
}

type HttpBody struct {
	Height    string       `json:"height"`
	Validator []*Validator `json:"result"`
}

func (c Client) Validators() ([]*Validator, error) {
	resp, err := c.apiClient.R().Get("/staking/validators")
	if err != nil {
		return nil, err
	}
	//fmt.Printf("\nResponse Body: %v", resp)

	var vals *HttpBody
	err = json.Unmarshal(resp.Body(), &vals)
	if err != nil {
		return nil, err
	}

	return vals.Validator, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func test(client *Client) {
	start := int64(8)
	for i := start; i < 100; i++ {

		block, err := client.Block(i)
		if err != nil {
			fmt.Errorf("failed to query block using rpc client: %s", err)
			return
		}

		valSet, err := client.ValidatorSet(block.Block.LastCommit.Height)
		if err != nil {
			fmt.Errorf("failed to query validator set using rpc client: %s", err)
			return
		}
		fmt.Println(valSet.BlockHeight, len(valSet.Validators))
		for inex, validator := range valSet.Validators {
			validatorJSON, _ := json.Marshal(validator)
			fmt.Println(inex, string(validatorJSON))
		}
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
var (
	lastAlarmUnix = int64(0)
	minsCount3    = 0
	minsCount5    = 0
)

func processAlarm(alarmTimeUnix int64) {
	if lastAlarmUnix == 0 {
		lastAlarmUnix = alarmTimeUnix
		minsCount3++
		fmt.Println("send sms in 3 minutes")
	} else if minsCount3 < 3 && alarmTimeUnix/180 == lastAlarmUnix/180 {
		minsCount3++
		fmt.Println("send sms in 3 minutes")
	} else if minsCount5 < 2 && alarmTimeUnix/300 == lastAlarmUnix/300 {
		minsCount5++
		fmt.Println("send sms in 5 minutes")
	} else if alarmTimeUnix/600 == lastAlarmUnix/600 {
		fmt.Println("send sms in 10 minutes")
	}

	if alarmTimeUnix/3600 != lastAlarmUnix/3600 {
		minsCount3 = 0
		minsCount5 = 0
	}
	lastAlarmUnix = alarmTimeUnix
}


func main() {
	cfg := ParseConfig()

	client := NewClient(cfg)

	//test(client)

	vals, err := client.Validators()
	if err != nil {
		fmt.Errorf("failed to query validators using rpc client: %s", err)
		return
	}
	validatorLen := len(vals)
	fmt.Println("validatorLen", validatorLen)

	lastHeight := int64(0)
	lastTime := time.Now()

	apiErrTry := 0
	for {
		if apiErrTry > 3 {
			fmt.Errorf("API failed 3 times, send alarm !!!")
			processAlarm(time.Now().Unix())
			apiErrTry = 0
		}

		height, err := client.LatestBlockHeight()
		if err != nil {
			fmt.Errorf("failed to query block height")
			apiErrTry++
			time.Sleep(1 * time.Second)
			continue
		}
		if lastHeight == height {
			alarmTime := time.Now()
			elapse := alarmTime.Sub(lastTime).Seconds()
			if elapse < 120 {
				time.Sleep(5 * time.Second)
				continue
			}
			fmt.Errorf("Height is not right, send alarm !!!")
			processAlarm(alarmTime.Unix())
			continue
		}
		lastHeight = height
		lastTime = time.Now()

		block, err := client.Block(height)
		if err != nil {
			fmt.Errorf("failed to query block using rpc client: %s", err)
			time.Sleep(1 * time.Second)
			apiErrTry++
			continue
		}

		valSet, err := client.ValidatorSet(block.Block.LastCommit.Height)
		if err != nil {
			fmt.Errorf("failed to query validator set using rpc client: %s", err)
			time.Sleep(1 * time.Second)
			apiErrTry++
			continue
		}
		fmt.Println(valSet.BlockHeight, len(valSet.Validators), validatorLen)

		if len(valSet.Validators) < validatorLen {
			fmt.Errorf("Validators is not enough, send alarm !!!")
			processAlarm(time.Now().Unix())
		}
		apiErrTry = 0
	}
}
