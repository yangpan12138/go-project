package obsClient

import (
	"go-interface/config"

	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
)

type ObsClient struct {
	*obs.ObsClient
	conf config.Config
}

func NewObsClient(conf config.Config) (*ObsClient, error) {
	ak := conf.ObsConfig.Ak
	sk := conf.ObsConfig.Sk
	endpoint := conf.ObsConfig.Endpoint
	// 创建ObsClient结构体
	var obsClient, err = obs.New(ak, sk, endpoint, obs.WithSignature(obs.SignatureObs))
	if err != nil {
		return nil, err
	}

	return &ObsClient{obsClient, conf}, nil
}

func (o *ObsClient) Close() error {
	o.ObsClient.Close()
	return nil
}
