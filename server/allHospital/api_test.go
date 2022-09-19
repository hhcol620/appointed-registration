package alladdress

import (
	"fmt"
	"testing"
)

func Test_GetAddress(t *testing.T) {
	arr, err := GetAddress(0, "0", 1)
	if err != nil {
		return
	}
	fmt.Println(arr)
}
