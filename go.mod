module apitest

go 1.15

require (
	chainmaker.org/chainmaker-go/common v0.0.0
	chainmaker.org/chainmaker-sdk-go v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.16.0
)

replace chainmaker.org/chainmaker-go/common => ./chainmaker-sdk-go/common

replace chainmaker.org/chainmaker-sdk-go => ./chainmaker-sdk-go
