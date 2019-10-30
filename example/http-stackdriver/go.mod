module github.com/open-telemetry/opentelemetry-go/example/http-stackdriver

go 1.13

replace (
	go.opentelemetry.io => ../..
	go.opentelemetry.io/exporter/trace/stackdriver => ../../exporter/trace/stackdriver
)

require (
	go.opentelemetry.io v0.0.0-20191025183852-68310ab97435
	go.opentelemetry.io/exporter/trace/stackdriver v0.0.0-20191030043218-cdf3f492d11d
	google.golang.org/api v0.13.0
	google.golang.org/grpc v1.24.0
)
