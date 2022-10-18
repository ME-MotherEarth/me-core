//go:generate protoc -I=proto -I=$GOPATH/src -I=$GOPATH/src/github.com/ME-MotherEarth/protobuf/protobuf  --gogoslick_out=. mect.proto
package mect

import "math/big"

// New returns a new batch from given buffers
func New() *MECToken {
	return &MECToken{
		Value: big.NewInt(0),
	}
}
