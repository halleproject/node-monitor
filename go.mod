module github.com/halleproject/node-monitor

go 1.14


require (
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.535
	github.com/cosmos/cosmos-sdk v0.34.4-0.20200406170659-df5badaf4c2b
	github.com/cosmos/ethermint v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.0.0
	github.com/tendermint/tendermint v0.33.3
)


replace github.com/cosmos/cosmos-sdk => github.com/halleproject/cosmos-sdk v0.34.4-0.1.0

replace github.com/cosmos/ethermint => github.com/halleproject/hallechain v0.1.4
