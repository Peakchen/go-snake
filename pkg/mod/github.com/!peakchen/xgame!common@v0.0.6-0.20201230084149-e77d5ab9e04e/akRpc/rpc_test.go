package akRpc

import (
	"fmt"
	"log"
	"testing"
)

type TModuleTest struct {
}

const (
	EModuleName string = "ModuleTest"
)

func (this *TModuleTest) AfterCall() error {
	fmt.Println("module test rpc, after call.")
	return nil
}

func (this *TModuleTest) Action(a int32, b bool, c []int32) (error, bool) {
	return nil, false
}

func rpctest(t *testing.T) {
	var (
		a int32   = 1
		b bool    = false
		c []int32 = []int32{1, 2, 3}
	)
	in := []interface{}{a, b, c}
	out := []interface{}{}
	if err := Aorpc.Call("key", EModuleName, "Action", in, out); err != nil {
		log.Fatal("rpc call fail: ", EModuleName, in)
	}
}

func init() {
	Register("ModuleTest", &TModuleTest{})
}
