package promeasurer

import "github.com/kazegusuri/grpcinstrument"

// NewMeasurer constructs an Measurer implementation that exports gRPC metrics
// via Prometheus.
func NewMeasurer() grpcinstrument.Measurer {
	return newMeasurer()
}
