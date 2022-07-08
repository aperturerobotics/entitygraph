module github.com/aperturerobotics/entitygraph

go 1.18

replace (
	github.com/sirupsen/logrus => github.com/aperturerobotics/logrus v1.8.2-0.20220322010420-77ab346a2cf8 // aperture
	google.golang.org/protobuf => github.com/aperturerobotics/protobuf-go v1.27.2-0.20220609075637-a1d116b0035f // aperture
)

require (
	github.com/aperturerobotics/controllerbus v0.11.1
	github.com/blang/semver v3.5.1+incompatible
	github.com/sirupsen/logrus v1.8.2-0.20220112234510-85981c045988
	google.golang.org/protobuf v1.28.0
)

require golang.org/x/sys v0.0.0-20220704084225-05e143d24a9e // indirect
