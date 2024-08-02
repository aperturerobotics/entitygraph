module github.com/aperturerobotics/entitygraph

go 1.22

require github.com/aperturerobotics/controllerbus v0.45.6 // latest

// Note: the below is from the ControllerBus go.mod

require (
	github.com/aperturerobotics/protobuf-go-lite v0.6.5 // latest
	github.com/aperturerobotics/util v1.23.5 // indirect; latest
)

require (
	github.com/aperturerobotics/common v0.16.8
	github.com/blang/semver v3.5.1+incompatible
	github.com/sirupsen/logrus v1.9.3
)

require (
	github.com/aperturerobotics/json-iterator-lite v1.0.0 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect
	golang.org/x/sys v0.20.0 // indirect
)
