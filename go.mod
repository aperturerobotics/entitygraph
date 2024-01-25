module github.com/aperturerobotics/entitygraph

go 1.21

require github.com/aperturerobotics/controllerbus v0.32.3

// Note: the below is from the ControllerBus go.mod

replace (
	github.com/sirupsen/logrus => github.com/aperturerobotics/logrus v1.9.4-0.20240119050608-13332fb58195 // aperture
	google.golang.org/protobuf => github.com/aperturerobotics/protobuf-go v1.32.1-0.20240118233629-6aeee82c476f // aperture
)

require (
	github.com/aperturerobotics/util v1.13.1 // indirect; latest
	github.com/blang/semver v3.5.1+incompatible
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.9.3
	google.golang.org/protobuf v1.32.0
)

require (
	github.com/stretchr/testify v1.8.4 // indirect
	golang.org/x/sys v0.15.0 // indirect
)
