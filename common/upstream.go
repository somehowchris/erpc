package common

import "context"

type UpstreamType string

const (
	UpstreamTypeEvm        UpstreamType = "evm"
	UpstreamTypeEvmAlchemy UpstreamType = "evm+alchemy"
	UpstreamTypeEvmEnvio   UpstreamType = "evm+envio"
	UpstreamTypeEvmPimlico UpstreamType = "evm+pimlico"
)

type Upstream interface {
	Config() *UpstreamConfig
	Vendor() Vendor
	SupportsNetwork(networkId string) (bool, error)
	EvmGetChainId(ctx context.Context) (string, error)
}
