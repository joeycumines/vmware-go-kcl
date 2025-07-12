module github.com/vmware/vmware-go-kcl

go 1.24.5

require (
	github.com/aws/aws-sdk-go v1.55.7
	github.com/awslabs/kinesis-aggregation/go v0.0.0-20241004223953-c2774b1ab29b
	github.com/golang/protobuf v1.5.4
	github.com/google/uuid v1.6.0
	github.com/prometheus/client_golang v1.22.0
	github.com/prometheus/common v0.65.0
	github.com/rs/zerolog v1.34.0
	github.com/sirupsen/logrus v1.9.3
	github.com/stretchr/testify v1.10.0
	go.uber.org/zap v1.27.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)

require (
	github.com/BurntSushi/toml v1.5.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/joeycumines/simple-command-output-filter v0.2.1 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.6.2 // indirect
	github.com/prometheus/procfs v0.17.0 // indirect
	github.com/yuin/goldmark v1.7.12 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/exp/typeparams v0.0.0-20250711185948-6ae5c78190dc // indirect
	golang.org/x/mod v0.26.0 // indirect
	golang.org/x/sync v0.16.0 // indirect
	golang.org/x/sys v0.34.0 // indirect
	golang.org/x/telemetry v0.0.0-20250710130107-8d8967aff50b // indirect
	golang.org/x/tools v0.35.0 // indirect
	golang.org/x/tools/go/expect v0.1.1-deprecated // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	honnef.co/go/tools v0.6.1 // indirect
)

tool (
	github.com/joeycumines/simple-command-output-filter
	golang.org/x/tools/cmd/deadcode
	golang.org/x/tools/cmd/godoc
	honnef.co/go/tools/cmd/staticcheck
)
