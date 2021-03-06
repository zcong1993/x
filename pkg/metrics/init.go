package metrics

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus/collectors"

	"github.com/prometheus/client_golang/prometheus"
)

var once sync.Once

// InitMetrics 不要使用默认 register, 太容易产生冲突.
// InitMetrics create a clean prometheus Registry to avoid conflicts.
func InitMetrics() *prometheus.Registry {
	me := prometheus.NewRegistry()

	once.Do(func() {
		me.MustRegister(
			collectors.NewGoCollector(),
			collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		)
	})

	prometheus.DefaultRegisterer = me
	return me
}
