module github.com/aperturerobotics/genproto/tools

go 1.13

// aperture: use protobuf 1.3.x based fork for compatibility
replace (
	github.com/golang/protobuf => github.com/aperturerobotics/go-protobuf-1.3.x v0.0.0-20200726220404-fa7f51c52df0 // aperture-1.3.x
	google.golang.org/genproto => google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55
	google.golang.org/grpc => google.golang.org/grpc v1.30.0
)

require (
	github.com/golang/protobuf v1.3.3-0.20190827175835-822fe56949f5
	github.com/square/goprotowrap v0.0.0-20190116012208-bb93590db2db
)
