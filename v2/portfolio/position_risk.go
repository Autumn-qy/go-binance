package portfolio

import (
	"context"

	"github.com/Autumn-qy/go-binance/v2/delivery"
	"github.com/Autumn-qy/go-binance/v2/futures"
)

type GetPositionRiskServiceCM struct {
	ds *delivery.GetPositionRiskService
}

func (s *GetPositionRiskServiceCM) Do(
	ctx context.Context,
	opts ...delivery.RequestOption,
) ([]*delivery.PositionRisk, error) {
	opts = append(opts, delivery.WithEndpoint("/papi/v1/cm/positionRisk"))
	return s.ds.Do(ctx, opts...)
}

type GetPositionRiskServiceUM struct {
	fs *futures.GetPositionRiskService
}

func (s *GetPositionRiskServiceUM) Do(
	ctx context.Context,
	opts ...futures.RequestOption,
) ([]*futures.PositionRisk, error) {
	opts = append(opts, futures.WithEndpoint("/papi/v1/um/positionRisk"))
	return s.fs.Do(ctx, opts...)
}
