package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"testing"
)

func TestCasbin01(t *testing.T) {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		fmt.Println("casbin.NewEnforcer err:" + err.Error())
	}
	b, err := e.Enforce("alice", "data1", "read")
	if err != nil {
		fmt.Println("e.Enforce err:" + err.Error())
	}
	if !b {
		t.Log("enforce failed")
	} else {
		t.Log("enforce success")
	}
}

func TestCasbin02(t *testing.T) {
	e, _ := casbin.NewEnforcer("./model.conf", "./policy.csv")
	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.
	if res, _ := e.Enforce(sub, obj, act); res {
		fmt.Println("enforce success")
		// permit alice to read data1
	} else {
		fmt.Println("enforce failed")
		// deny the request, show an error
	}
}
