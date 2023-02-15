package tracer

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"time"
)

func ProvideTracer(service string) (opentracing.Tracer, error) {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  "esgrs_jaeger:6831",
		},
	}
	tracer, _, err := cfg.New(
		service,
	)
	if err != nil {
		return nil, err
	}

	return tracer, nil
}
