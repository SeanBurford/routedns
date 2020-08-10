package rdns

import (
	"expvar"
	"fmt"
	"net"
)

// Listener is an interface for a DNS listener.
type Listener interface {
	Start() error
	fmt.Stringer
}

// ClientInfo carries information about the client making the request that
// can be used to route requests.
type ClientInfo struct {
	SourceIP net.IP
}

// Metrics that are available from listeners and clients.
type ListenerMetrics struct {
	// DNS query count.
	query *expvar.Int
	// DNS response type counts.
	response *expvar.Map
	// Number of queries dropped (denied).
	drop *expvar.Int
	// RouteDNS failure reason counts.
	err *expvar.Map
	// Maximum number of queries queued (optional).
	maxQueueLen *expvar.Int
}

func NewListenerMetrics(base string, id string) *ListenerMetrics {
	return &ListenerMetrics{
		query:       getVarInt(base, id, "query"),
		response:    getVarMap(base, id, "response"),
		drop:        getVarInt(base, id, "drop"),
		err:         getVarMap(base, id, "error"),
		maxQueueLen: getVarInt(base, id, "maxqueue"),
	}
}
