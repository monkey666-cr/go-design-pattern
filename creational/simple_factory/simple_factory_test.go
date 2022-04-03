package simple_factory

import (
	"fmt"
	"testing"
)

func TestNewAPI(t *testing.T) {
	api := NewAPI("cn")
	fmt.Println(api.Say("chenrun"))
}
