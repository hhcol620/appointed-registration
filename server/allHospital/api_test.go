package alladdress

import (
	"fmt"
	"testing"
)

func Test_GetAddress(t *testing.T) {
	arr, err := GetAddress()
	if err != nil {
		return
	}
	fmt.Println(len(arr))
}
