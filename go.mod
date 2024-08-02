module github.com/aperturerobotics/entitygraph

go 1.22

require github.com/aperturerobotics/controllerbus v0.47.3 // latest

// Note: the below is from the ControllerBus go.mod

require (
	github.com/aperturerobotics/protobuf-go-lite v0.6.5 // latest
	github.com/aperturerobotics/util v1.25.7 // indirect; latest
)

require (
	github.com/aperturerobotics/common v0.18.3
	github.com/blang/semver/v4 v4.0.0
	github.com/sirupsen/logrus v1.9.3
)

require (
	github.com/aperturerobotics/json-iterator-lite v1.0.0 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/exp v0.0.0-20240719175910-8a7402abbf56 // indirect
	golang.org/x/sys v0.22.0 // indirect
)
