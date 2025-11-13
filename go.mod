module github.com/aperturerobotics/entitygraph

go 1.25

toolchain go1.25.4

require github.com/aperturerobotics/controllerbus v0.51.3 // latest

// Note: the below is from the controllerbus go.mod

// This fork drops ecdsa, drops secp256k1, adds eddilithium2 and eddilithium3
replace github.com/libp2p/go-libp2p => github.com/aperturerobotics/go-libp2p v0.37.1-0.20241111002741-5cfbb50b74e0 // aperture

replace github.com/libp2p/go-msgio => github.com/aperturerobotics/go-libp2p-msgio v0.0.0-20240511033615-1b69178aa5c8 // aperture

require (
	github.com/aperturerobotics/common v0.22.14 // latest
	github.com/aperturerobotics/json-iterator-lite v1.0.1-0.20251104042408-0c9eb8a3f726 // indirect; latest
	github.com/aperturerobotics/protobuf-go-lite v0.11.0 // latest
	github.com/aperturerobotics/util v1.31.4 // indirect; latest
)

require (
	github.com/blang/semver/v4 v4.0.0
	github.com/sirupsen/logrus v1.9.3
)

require (
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/exp v0.0.0-20250408133849-7e4ce0ab07d0 // indirect
	golang.org/x/sys v0.37.0 // indirect
)
