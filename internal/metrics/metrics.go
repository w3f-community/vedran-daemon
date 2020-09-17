package metrics

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/prometheus/common/expfmt"
)

// Client is used to scrape metrics data from prometheus server
type Client interface {
	GetNodeMetrics() (*Metrics, error)
}

// Metrics required to be sent to load balancer
type Metrics struct {
	PeerCount             *float64 `json:"peer_count"`
	BestBlockHeight       *float64 `json:"best_block_height"`
	FinalizedBlockHeight  *float64 `json:"finalized_block_height"`
	ReadyTransactionCount *float64 `json:"read_transaction_count"`
}

// NewClient creates metrics client instance
func NewClient(baseURL *url.URL) Client {
	return &client{BaseURL: baseURL}
}

type client struct {
	BaseURL *url.URL
}

const (
	metricsEndpoint = "/metrics"
)

// GetNodeMetrics retrieves polkadot metrics from prometheus server
func (client *client) GetNodeMetrics() (*Metrics, error) {
	metricsURL, _ := client.BaseURL.Parse(metricsEndpoint)
	resp, err := http.Get(metricsURL.String())
	if err != nil {
		return nil, fmt.Errorf("Metrics endpoint returned error: %v", err)
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Metrics endpoint returned invalid status code: %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	var parser expfmt.TextParser
	metricFamilies, err := parser.TextToMetricFamilies(resp.Body)
	if err != nil {
		return nil, err
	}

	metrics := &Metrics{
		metricFamilies["polkadot_sync_peers"].GetMetric()[0].Gauge.Value,
		metricFamilies["polkadot_block_height"].GetMetric()[0].Gauge.Value,
		metricFamilies["polkadot_block_height"].GetMetric()[1].Gauge.Value,
		metricFamilies["polkadot_ready_transactions_number"].GetMetric()[0].Gauge.Value,
	}
	return metrics, nil
}
