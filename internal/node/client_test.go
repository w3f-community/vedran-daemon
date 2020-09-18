package node

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

type handleFnMock func(http.ResponseWriter, *http.Request)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
}

func teardown() {
	server.Close()
}

func TestGetMetrics(t *testing.T) {
	type args struct {
		baseURL string
	}

	type Test struct {
		name       string
		args       args
		want       *Metrics
		wantErr    bool
		handleFunc handleFnMock
	}

	peerCount := float64(19)
	bestBlockHeight := float64(432933)
	finalizedBlockHeight := float64(432640)
	readyTransactionCount := float64(0)
	tests := []Test{
		{
			name:    "Returns error if metrics endpoint does not exist",
			args:    args{"invalid"},
			want:    nil,
			wantErr: true,
			handleFunc: func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "Not Found", 404)
			}},
		{
			name:    "Returns error if metrics endpoint returns not found",
			args:    args{"valid"},
			want:    nil,
			wantErr: true,
			handleFunc: func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "Not Found", 404)
			}},
		{
			name:    "Returns error if parsing metrics fails",
			args:    args{"valid"},
			want:    nil,
			wantErr: true,
			handleFunc: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				_, _ = io.WriteString(w, `invalid`)
			}},
		{
			name: "Returns metrics if prometheus response valid",
			args: args{"valid"},
			want: &Metrics{
				PeerCount:             &peerCount,
				BestBlockHeight:       &bestBlockHeight,
				FinalizedBlockHeight:  &finalizedBlockHeight,
				ReadyTransactionCount: &readyTransactionCount,
			},
			wantErr: false,
			handleFunc: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				_, _ = io.WriteString(
					w,
					`
					# HELP polkadot_sync_peers Number of peers we sync with
					# TYPE polkadot_sync_peers gauge
					polkadot_sync_peers 19
					# HELP polkadot_block_height Block height info of the chain
					# TYPE polkadot_block_height gauge
					polkadot_block_height{status="best"} 432933
					polkadot_block_height{status="finalized"} 432640
					polkadot_block_height{status="sync_target"} 1547694
					# HELP polkadot_ready_transactions_number Number of transactions in the ready queue
					# TYPE polkadot_ready_transactions_number gauge
					polkadot_ready_transactions_number 0
					`)
			}},
	}

	for _, tt := range tests {
		setup()
		defer teardown()

		t.Run(tt.name, func(t *testing.T) {
			var baseURL *url.URL
			if tt.args.baseURL == "valid" {
				baseURL, _ = url.Parse(server.URL)
			} else {
				baseURL, _ = url.Parse("http://invalid:3000")
			}
			mux.HandleFunc("/metrics", tt.handleFunc)

			nodeClient := NewClient(baseURL, baseURL)
			got, err := nodeClient.GetMetrics()

			if (err != nil) != tt.wantErr {
				t.Errorf("GetMetrics() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMetrics() = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_client_GetRPCURL(t *testing.T) {
	rpcURL, _ := url.Parse("http://localhost:9933")
	metricsURL, _ := url.Parse("http://localhost:9615")

	type fields struct {
		MetricsBaseURL *url.URL
		RPCBaseURL     *url.URL
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Returns string representation of rpc url",
			fields: fields{metricsURL, rpcURL},
			want:   "http://localhost:9933"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &client{
				MetricsBaseURL: tt.fields.MetricsBaseURL,
				RPCBaseURL:     tt.fields.RPCBaseURL,
			}

			if got := client.GetRPCURL(); got != tt.want {
				t.Errorf("client.GetRPCURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_GetMetricsURL(t *testing.T) {
	rpcURL, _ := url.Parse("http://localhost:9933")
	metricsURL, _ := url.Parse("http://localhost:9615")

	type fields struct {
		MetricsBaseURL *url.URL
		RPCBaseURL     *url.URL
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Returns string representation of rpc url",
			fields: fields{metricsURL, rpcURL},
			want:   "http://localhost:9615"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &client{
				MetricsBaseURL: tt.fields.MetricsBaseURL,
				RPCBaseURL:     tt.fields.RPCBaseURL,
			}

			if got := client.GetMetricsURL(); got != tt.want {
				t.Errorf("client.GetMetricsURL() = %v, want %v", got, tt.want)
			}
		})
	}
}