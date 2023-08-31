package server

import (
	"context"

	protos "github.com/bongerka/microservice/currency/protos/currency"
	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	protos.UnimplementedCurrencyServer
	log hclog.Logger
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{log: l}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle request for GetRate", "base", rr.GetBase(), "dest", rr.GetDestination())
	return &protos.RateResponse{Rate: 0.5}, nil
}

func (c *Currency) SubscribeRates(stream protos.Currency_SubscribeRatesServer) error {
	c.log.Info("Handle subscription request for SubscribeRates")

	for {
		// some logic
	}

	return nil
}
