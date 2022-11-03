module github.com/aperturerobotics/entitygraph

go 1.18

require github.com/aperturerobotics/controllerbus v0.18.2-0.20221019082518-c9323cca3baa

// Note: the below is from the ControllerBus go.mod

replace (
	github.com/sirupsen/logrus => github.com/aperturerobotics/logrus v1.8.2-0.20220322010420-77ab346a2cf8 // aperture
	google.golang.org/protobuf => github.com/aperturerobotics/protobuf-go v1.28.2-0.20221007002036-6510dd3bc392 // aperture
)

require (
	github.com/blang/semver v3.5.1+incompatible
	github.com/sirupsen/logrus v1.9.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab // indirect
)
