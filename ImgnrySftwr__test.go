package UniqueID_Go

import "testing"
import "fmt"

func Test_ImgnrySftwr (t *testing.T) {
	for _ba01 := 1; _ba01 <= 10; _ba01 ++ {
		_ca01 := UniqueID_Go__Intlz(16)
		fmt.Println (_ca01, len(_ca01))
	}
}
