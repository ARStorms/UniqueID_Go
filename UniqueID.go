// Package UniqueID_Go Description
//
// Package UniqueID_Go is the home of a data type we call "UniqueID".
//
// Unique ID → UniqueID
//
// As the name implies, a Unique ID instance is a unique (lowercase-only) alphanumeric ID.
//
// These IDs can be used to identify database records, customer accounts, and so on.
//
// These IDs are meant for simple use-cases only. Look elsewhere if you need something that is
// also crypto-safe.
package UniqueID_Go

import  "github.com/thanhpk/randstr"
import  "math/big"
import  "strings"
import  "time"

// Initializer UniqueID_Go__Intlz creates and initializes a new Unique ID.
//
// This initializer guarantees a unique ID as long as two or more threads (goroutines) do not
// call it at the exact same time.
//
// This initializer never generates less than 12-char IDs.
func    UniqueID_Go__Intlz (Length int) (string) {
	_ba01 := time.Now  ().In (time.FixedZone ("UTC +00:00", 0))
	_bb01 := big.NewInt (_ba01.UnixNano ()).Text (35)
	for len (_bb01) < Length {
		 _bb01  =  _bb01 + randstr.String (1)
	}
	_bb01  = strings.ToLower (_bb01)
	return   _bb01
}
