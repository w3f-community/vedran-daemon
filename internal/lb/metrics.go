package lb

import (
	"log"
	"net/http"
	"time"

	"github.com/NodeFactoryIo/vedran-daemon/internal/metrics"
	"github.com/getsentry/sentry-go"
)

// MetricsService is used for sending node metrics to load balancer
type MetricsService interface {
	Send(metrics *metrics.Metrics) (*http.Response, error)
}

type metricsService struct {
	client *Client
}

type metricsRequest struct {
	timestamp time.Time
}

const (
	metricsEndpoint = "/api/v1/nodes/metrics"
)

func (ms *metricsService) Send(metrics *metrics.Metrics) (*http.Response, error) {
	log.Println("Sending metrics to load balancer")
	req, _ := ms.client.NewRequest(http.MethodPut, metricsEndpoint, metrics)
	resp, err := ms.client.Do(req, nil)

	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	return resp, err
}
