module github.com/aperturerobotics/entitygraph

go 1.23

toolchain go1.23.2

require github.com/aperturerobotics/controllerbus v0.47.8 // latest

// Note: the below is from the controllerbus go.mod

// Optional: this fork uses go-protobuf-lite and adds post-quantum crypto support.
replace github.com/libp2p/go-libp2p => github.com/aperturerobotics/go-libp2p v0.36.3-0.20241002070357-a2e1c3498dd3 // aperture

replace github.com/libp2p/go-msgio => github.com/aperturerobotics/go-libp2p-msgio v0.0.0-20240511033615-1b69178aa5c8 // aperture

require (
	github.com/aperturerobotics/common v0.18.8 // latest
	github.com/aperturerobotics/json-iterator-lite v1.0.0 // indirect; latest
	github.com/aperturerobotics/protobuf-go-lite v0.7.0 // latest
	github.com/aperturerobotics/util v1.25.10 // indirect; latest
)

require (
	github.com/blang/semver/v4 v4.0.0
	github.com/sirupsen/logrus v1.9.3
)

require (
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/exp v0.0.0-20241009180824-f66d83c29e7c // indirect
	golang.org/x/sys v0.26.0 // indirect
)
