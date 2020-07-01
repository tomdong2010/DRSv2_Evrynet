package vfactory

import (
	"github.com/velo-protocol/DRSv2_Evrynet/go/cmd/gvel/entity"
	"github.com/velo-protocol/DRSv2_Evrynet/go/libs/vclient"
	"github.com/velo-protocol/DRSv2_Evrynet/go/libs/vclient/ivclient"
)

func (v *veloFactory) NewClientFromConfig(privateKey string) (ivclient.VClient, error) {
	return v.NewClient(&entity.NewClientInput{
		RpcUrl:     v.AppConfig.GetRpcUrl(),
		PrivateKey: privateKey,
		ContractAddress: &vclient.ContractAddress{
			DrsAddress:   v.AppConfig.GetDrsAddress(),
			HeartAddress: v.AppConfig.GetHeartAddress(),
		},
	})
}
