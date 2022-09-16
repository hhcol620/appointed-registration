package alladdress

import (
	"fmt"
	"testing"
)

func Test_GetAddress(t *testing.T) {
	arr, err := GetAddress(3, 1, "110101")
	if err != nil {
		return
	}
	fmt.Println(arr)
}
