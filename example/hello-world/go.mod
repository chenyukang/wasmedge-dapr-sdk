module github.com/dapr/quickstarts/hello-go-sdk

go 1.16

require (
	github.com/alecthomas/units v0.0.0-20210208195552-ff826a37aa15 // indirect
	github.com/dapr/go-sdk v0.0.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
)

replace github.com/dapr/go-sdk v0.0.0 => ../../../wasmedge-dapr-sdk
