package service

import (
	"social-pet/global"

	"github.com/astaxie/beego/logs"
	"github.com/ethereum/go-ethereum/ethclient"
)

func InitWeb3() {
	var err error
	var client *ethclient.Client
	client, err = ethclient.Dial(global.HOST)
	if err != nil {
		panic(err.(any))
	}

	global.Ethclient = client

	logs.Info("connect web3 success.")
}
