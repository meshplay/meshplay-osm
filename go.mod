module github.com/khulnasoft/meshplay-adapter-template

go 1.15

replace github.com/kumarabd/gokit => github.com/realnighthawk/bucky v0.0.3

require (
	github.com/golang/protobuf v1.4.3
	github.com/golangci/golangci-lint v1.33.0 // indirect
	github.com/google/go-cmp v0.5.3 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/kumarabd/gokit v1.0.0 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.7.1
	go.opentelemetry.io/otel v0.10.0
	go.opentelemetry.io/otel/exporters/trace/jaeger v0.10.0
	go.opentelemetry.io/otel/sdk v0.10.0
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4
	google.golang.org/genproto v0.0.0-20200731012542-8145dea6a485 // indirect
	google.golang.org/grpc v1.34.0
)
