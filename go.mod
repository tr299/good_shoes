module git.onepay.vn/onepay/ddsp

go 1.16

require (
	github.com/gin-gonic/gin v1.9.1
	github.com/godror/godror v0.40.3
	github.com/google/uuid v1.3.0
	github.com/lib/pq v1.10.9
	github.com/sirupsen/logrus v1.9.3
	github.com/spf13/cobra v1.7.0
	github.com/spf13/viper v1.16.0
	go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin v0.44.0
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.44.0
	go.opentelemetry.io/otel v1.18.0
	go.opentelemetry.io/otel/exporters/jaeger v1.17.0
	go.opentelemetry.io/otel/sdk v1.18.0
	go.opentelemetry.io/otel/trace v1.18.0
	go.uber.org/zap v1.25.0
	golang.org/x/net v0.15.0
)

require golang.org/x/crypto v0.14.0 // indirect
