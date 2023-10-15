package example

import (
	"math/big"
	"social-pet/contracts"
	"social-pet/global"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func BalanceOf(account string) (balance *big.Int, err error) {
	var cli *contracts.SocialPet
	if cli, err = contracts.NewSocialPet(common.HexToAddress(global.Social_Pet), global.Ethclient); err != nil {
		return
	}

	balance, err = cli.BalanceOf(&bind.CallOpts{}, common.HexToAddress(account))

	return
}
