package global

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-xorm/xorm"
)

const (
	//链RPC
	HOST       = "https://goerli.infura.io/v3/53916f7f4e71473fbd9b2ba4235fa961"
	Social_Pet = "0x98920551BC17A6617872c3205fAa34C488A6cEdD"
)

var (
	//数据库引擎、=
	DB *xorm.Engine
	//web3引擎
	Ethclient *ethclient.Client
)
