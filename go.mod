module github.com/aperturerobotics/entitygraph

go 1.22

require github.com/aperturerobotics/controllerbus v0.41.2 // latest

// Note: the below is from the ControllerBus go.mod

require (
	github.com/aperturerobotics/protobuf-go-lite v0.4.5 // indirect; latest
	github.com/aperturerobotics/util v1.17.1 // indirect; latest
)

require (
	github.com/blang/semver v3.5.1+incompatible
	github.com/sirupsen/logrus v1.9.3
	google.golang.org/protobuf v1.33.0
)

require (
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/exp v0.0.0-20240416160154-fe59bbe5cc7f // indirect
	golang.org/x/sys v0.19.0 // indirect
)
