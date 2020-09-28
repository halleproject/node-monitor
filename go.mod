module github.com/halleproject/node-monitor

go 1.14

require (
	github.com/Robbin-Liu/go-binance-sdk v0.0.0-20200728021042-9ef0842abec7
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.535
	github.com/cosmos/cosmos-sdk v0.39.1
	github.com/cosmos/ethermint v0.0.0-00010101000000-000000000000
	github.com/go-resty/resty/v2 v2.3.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v0.0.7
	github.com/spf13/viper v1.6.3
	github.com/tendermint/go-amino v0.15.1
	github.com/tendermint/tendermint v0.33.3
)

replace github.com/cosmos/cosmos-sdk => github.com/halleproject/cosmos-sdk v0.34.4-0.1.0

replace github.com/cosmos/ethermint => github.com/halleproject/hallechain v0.1.4
