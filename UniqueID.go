package UniqueID_Go

import  "github.com/thanhpk/randstr"
import  "math/big"
import  "strings"
import  "time"

func    UniqueID_Go__Intlz (Length int) (string) {
	_ba01 := time.Now  ().In (time.FixedZone ("UTC +00:00", 0))
	_bb01 := big.NewInt (_ba01.UnixNano ()).Text (35)
	for len (_bb01) < Length {
		 _bb01  =  _bb01 + randstr.String (1)
	}
	_bb01  = strings.ToLower (_bb01)
	return   _bb01
}
