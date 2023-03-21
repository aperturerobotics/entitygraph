module github.com/aperturerobotics/entitygraph

go 1.19

require github.com/aperturerobotics/controllerbus v0.24.3-0.20230321080215-8e8b2646ba47

// Note: the below is from the ControllerBus go.mod

replace (
	github.com/sirupsen/logrus => github.com/aperturerobotics/logrus v1.9.1-0.20221224130652-ff61cbb763af // aperture
	google.golang.org/protobuf => github.com/aperturerobotics/protobuf-go v1.28.2-0.20230301012226-7fb3cdbd9197 // aperture
)

require (
	github.com/aperturerobotics/util v1.0.6-0.20230313071030-6260e05b62e5 // indirect; latest
	github.com/blang/semver v3.5.1+incompatible
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.9.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/stretchr/testify v1.8.1 // indirect
	golang.org/x/sys v0.6.0 // indirect
)
