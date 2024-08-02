module github.com/aperturerobotics/entitygraph

go 1.22

require github.com/aperturerobotics/controllerbus v0.47.3 // latest

// Note: the below is from the controllerbus go.mod

// Optional: this fork uses go-protobuf-lite and adds post-quantum crypto support.
replace github.com/libp2p/go-libp2p => github.com/aperturerobotics/go-libp2p v0.33.1-0.20240511223728-e0b67c111765 // aperture

replace github.com/libp2p/go-msgio => github.com/aperturerobotics/go-libp2p-msgio v0.0.0-20240511033615-1b69178aa5c8 // aperture

require (
	github.com/aperturerobotics/common v0.18.3 // latest
	github.com/aperturerobotics/json-iterator-lite v1.0.0 // indirect; latest
	github.com/aperturerobotics/protobuf-go-lite v0.6.5 // latest
	github.com/aperturerobotics/util v1.25.7 // indirect; latest
)

require (
	github.com/blang/semver/v4 v4.0.0
	github.com/sirupsen/logrus v1.9.3
)

require (
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/exp v0.0.0-20240719175910-8a7402abbf56 // indirect
	golang.org/x/sys v0.22.0 // indirect
)
