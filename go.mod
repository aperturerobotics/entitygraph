module github.com/aperturerobotics/entitygraph

go 1.18

replace (
	github.com/sirupsen/logrus => github.com/aperturerobotics/logrus v1.8.2-0.20220322010420-77ab346a2cf8 // aperture
	google.golang.org/protobuf => github.com/aperturerobotics/protobuf-go v1.27.2-0.20220603060022-78b627edc1c2 // aperture
)

require (
	github.com/aperturerobotics/controllerbus v0.10.2
	github.com/blang/semver v3.5.1+incompatible
	github.com/golang/protobuf v1.5.2
	github.com/sirupsen/logrus v1.8.0
	google.golang.org/protobuf v1.27.1
)

require golang.org/x/sys v0.0.0-20220319134239-a9b59b0215f8 // indirect
