package influxdb

import (
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/mochahub/coinprice-scraper/config"
	"go.uber.org/fx"
)

func NewInfluxDBClient(lc fx.Lifecycle, secrets *config.Secrets) (*influxdb2.Client, error) {
	influxDBURL := fmt.Sprintf("http://%s:%d", secrets.InfluxDbCredentials.Host, secrets.InfluxDbCredentials.Port)
	// Generate AuthToken Here: http://localhost:8086/onboarding/0
	client := influxdb2.NewClient(influxDBURL, secrets.InfluxDbCredentials.Token)
	//lc.Append(fx.Hook{
	//	OnStart: func(ctx context.Context) error {
	//
	//		response, err := client.Setup(
	//			ctx,
	//			secrets.InfluxDbCredentials.User,
	//			secrets.InfluxDbCredentials.Password,
	//			secrets.InfluxDbCredentials.Org,
	//			secrets.InfluxDbCredentials.Bucket,
	//			0)
	//		if err != nil {
	//			return err
	//		}
	//
	//		return nil
	//	},
	//	OnStop: func(ctx context.Context) error {
	//		client.Close()
	//		return nil
	//	},
	//})
	return &client, nil
}
